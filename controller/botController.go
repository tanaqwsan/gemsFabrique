package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"app/utils/res"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func IndexBot(c echo.Context) error {
	var bots []model.Bot

	err := config.DB.Find(&bots).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bot"))
	}

	if len(bots) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexBot(bots)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Bot data successfully retrieved", response))
}

func ShowBot(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var bot model.Bot

	if err := config.DB.First(&bot, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bot"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Bot data successfully retrieved", bot))
}

func StoreBot(c echo.Context) error {
	var bot model.Bot
	if err := c.Bind(&bot); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	if err := config.DB.Create(&bot).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store bot data"))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", bot))
}

func UpdateBot(c echo.Context) error {
	growid := c.Param("growid")

	var existingBot model.Bot
	var updatedBot model.Bot

	result := config.DB.Where("growid = ?", growid).First(&existingBot)
	if result.Error != nil {

	} else {
		updatedBot = existingBot
		updatedBot.Model = &gorm.Model{ID: existingBot.ID}
	}

	uGrowid := c.QueryParam("uGrowid")
	uAge := c.QueryParam("uAge")
	uGems := c.QueryParam("uGems")
	uLevel := c.QueryParam("uLevel")
	uIsSuspended := c.QueryParam("uIsSuspended")

	if uGrowid != "" {
		updatedBot.Growid = uGrowid
	}
	if uAge != "" {
		updatedBot.Age, _ = strconv.Atoi(uAge)
	}
	if uGems != "" {
		updatedBot.Gems, _ = strconv.Atoi(uGems)
	}
	if uLevel != "" {
		updatedBot.Level, _ = strconv.Atoi(uLevel)
	}
	if uIsSuspended != "" {
		updatedBot.IsSuspended, _ = strconv.Atoi(uIsSuspended)
	}
	config.DB.Save(&updatedBot)
	if result.Error != nil {
		return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", updatedBot))
	} else {
		return c.JSON(http.StatusOK, utils.SuccessResponse("Bot data successfully updated", nil))
	}
}

func DeleteBot(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	var existingBot model.Bot
	result := config.DB.First(&existingBot, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bot"))
	}
	config.DB.Delete(&existingBot)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Bot data successfully deleted", nil))
}
