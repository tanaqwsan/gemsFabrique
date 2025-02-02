package main

import (
	"app/config"
	"app/routes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logChannel chan string
var wg sync.WaitGroup

// Goroutine untuk menulis log secara asinkron
func asyncLogger(logFile io.Writer) {
	defer wg.Done()
	for logMessage := range logChannel {
		fmt.Fprintln(logFile, logMessage) // Pastikan log ditulis dengan newline
	}
}

func main() {
	e := echo.New()

	// Konfigurasi rotasi log dengan lumberjack
	logFile := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}
	defer logFile.Close()

	// Gabungkan log ke file dan CLI
	multiWriter := io.MultiWriter(logFile, os.Stdout)

	// Inisialisasi channel log asinkron
	logChannel = make(chan string, 1000)
	wg.Add(1)
	go asyncLogger(multiWriter)

	// Fungsi untuk mengirim log secara asinkron
	logMessage := func(msg string) {
		select {
		case logChannel <- msg:
		default:
			// Drop log jika channel penuh agar tidak blocking
		}
	}

	logMessage("Connecting to the database...")
	config.ConnectDB()
	logMessage("Successfully connected to the database!")

	e = routes.Init()

	// Middleware logging yang mengirim log ke channel asinkron
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logMessage(fmt.Sprintf(
				"time=%s, method=%s, uri=%s, status=%d, latency=%s, remote_ip=%s",
				v.StartTime.Format(time.RFC3339), v.Method, v.URI, v.Status, v.Latency.String(), c.RealIP(),
			))
			return nil
		},
	}))

	// Tutup channel dan tunggu goroutine selesai sebelum aplikasi berhenti
	defer func() {
		close(logChannel)
		wg.Wait()
	}()

	e.Logger.Fatal(e.Start("127.0.0.1:3090"))
}
