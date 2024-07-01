package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"app/utils/res"
	"github.com/labstack/echo/v4"
	"net/http"
)

func IndexWord(c echo.Context) error {
	var words []model.Word

	err := config.DB.Find(&words).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve word"))
	}

	if len(words) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexWord(words)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Word data successfully retrieved", response))
}

func ShowWord(c echo.Context) error {
	growid := c.Param("growid")
	var word model.Word

	if err := config.DB.Where("growid = ?", growid).First(&word).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve word"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Word data successfully retrieved", word))
}

func UpdateWord(c echo.Context) error {
	growid := c.Param("growid")
	word := c.Param("word")

	var existingWord model.Word
	var updatedWord model.Word

	result := config.DB.Where("growid = ?", growid).First(&existingWord)
	if result.Error != nil {

	} else {
		updatedWord = existingWord
	}
	updatedWord.Growid = growid
	updatedWord.Word = word
	config.DB.Save(&updatedWord)
	if result.Error != nil {
		return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", updatedWord))
	} else {
		return c.JSON(http.StatusOK, utils.SuccessResponse("Word data successfully updated", nil))
	}
}
func EraseWord(c echo.Context) error {
	growid := c.Param("growid")

	var existingWord model.Word
	var updatedWord model.Word

	result := config.DB.Where("growid = ?", growid).First(&existingWord)
	if result.Error != nil {

	} else {
		updatedWord = existingWord
	}
	updatedWord.Growid = growid
	updatedWord.Word = ""
	config.DB.Save(&updatedWord)
	if result.Error != nil {
		return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", nil))
	} else {
		return c.JSON(http.StatusOK, utils.SuccessResponse("Word data successfully updated", nil))
	}
}
