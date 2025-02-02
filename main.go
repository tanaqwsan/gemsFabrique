package main

import (
	"app/config"
	"app/routes"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/natefinch/lumberjack.v2" // Import library lumberjack untuk rotasi log
)

var logChannel chan string // Channel untuk log asinkron
var wg sync.WaitGroup      // WaitGroup untuk menunggu goroutine selesai

// Fungsi untuk menulis log secara asinkron
func asyncLogger(logFile *lumberjack.Logger) {
	defer wg.Done()
	for logMessage := range logChannel {
		_, _ = logFile.Write([]byte(logMessage + "\n")) // Tulis log ke file
	}
}

func main() {
	e := echo.New()

	// Konfigurasi rotasi log dengan lumberjack
	logFile := &lumberjack.Logger{
		Filename:   "app.log", // Nama file log
		MaxSize:    10,        // Ukuran maksimum file log dalam MB
		MaxBackups: 3,         // Jumlah file backup
		MaxAge:     28,        // Jumlah hari untuk menyimpan log
		Compress:   true,      // Kompres file log lama
	}
	defer logFile.Close()

	// Buat channel untuk log asinkron
	logChannel = make(chan string, 1000) // Buffer channel untuk 1000 log
	wg.Add(1)
	go asyncLogger(logFile) // Jalankan goroutine untuk menulis log asinkron

	// Set output log ke file menggunakan lumberjack
	e.Logger.SetOutput(logFile)

	// Log pesan ke channel (asinkron)
	logMessage := func(msg string) {
		logChannel <- msg
	}

	logMessage("Connecting to the database...")
	config.ConnectDB()
	logMessage("Successfully connected to the database!")

	e = routes.Init()

	// Tambahkan middleware logging
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	// Tutup channel dan tunggu goroutine selesai sebelum aplikasi berhenti
	defer func() {
		close(logChannel)
		wg.Wait()
	}()

	e.Logger.Fatal(e.Start("127.0.0.1:3090"))
}
