package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"app/utils/res"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

func IndexWorld(c echo.Context) error {
	var worlds []model.World

	err := config.DB.Find(&worlds).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}

	if len(worlds) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexWorld(worlds)

	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", response))
}

func IndexWorldOnlyName(c echo.Context) error {
	var worlds []model.World

	err := config.DB.Find(&worlds).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}

	if len(worlds) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexWorldNameIdOnly(worlds)

	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", response))

}

func ShowWorld(c echo.Context) error {
	name := c.Param("name")
	var existingWorld model.World
	err := config.DB.Where("name = ?", name).First(&existingWorld).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", existingWorld))
}

func StoreWorld(c echo.Context) error {
	var world model.World
	if err := c.Bind(&world); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	// Check if world already exists
	var checkWorld model.World
	if err := config.DB.Where("name = ?", world.Name).First(&checkWorld).Error; err == nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("World already exists"))
	}
	if err := config.DB.Create(&world).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store world data"))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", world))
}

func UpdateWorld(c echo.Context) error {
	name := c.Param("name")

	var existingWorld model.World
	var updatedWorld model.World
	result := config.DB.Where("name = ?", name).First(&existingWorld)
	if result.Error != nil {

	} else {
		updatedWorld = existingWorld
	}

	uName := c.QueryParam("uName")
	uNameId := c.QueryParam("uNameId")
	uOwner := c.QueryParam("uOwner")
	uType := c.QueryParam("uType")

	uIsSmallLock := c.QueryParam("uIsSmallLock")
	uIsNuked := c.QueryParam("uIsNuked")
	uSmallLockAge := c.QueryParam("uSmallLockAge")
	uFloatPepperBlockCount := c.QueryParam("uFloatPepperBlockCount")
	uFloatPepperSeedCount := c.QueryParam("uFloatPepperSeedCount")
	uTilePepperSeedCount := c.QueryParam("uTilePepperSeedCount")
	uTilePepperBlockCount := c.QueryParam("uTilePepperBlockCount")
	uFossilCount := c.QueryParam("uFossilCount")
	uSLOwner := c.QueryParam("uSLOwner")
	uBotHandlerId := c.QueryParam("uBotHandlerId")
	uGems := c.QueryParam("uGems")

	//http://localhost:8080/world/update/1?uName=world1&uNameId=world1&uOwner=owner1&uType=type1&uIsSmallLock=1&uIsNuked=1&uSmallLockAge=1&uFloatPepperBlockCount=1&uFloatPepperSeedCount=1&uTilePepperSeedCount=1&uTilePepperBlockCount=1&uFossilCount=1&uBotHandlerId=1

	if uName != "" {
		updatedWorld.Name = uName
	}
	if uNameId != "" {
		updatedWorld.NameId = uNameId
	}
	if uOwner != "" {
		updatedWorld.Owner = uOwner
	}
	if uType != "" {
		updatedWorld.Type = uType
	}

	if uIsSmallLock != "" {
		updatedWorld.IsSmallLock, _ = strconv.Atoi(uIsSmallLock)
	}
	if uIsNuked != "" {
		updatedWorld.IsNuked, _ = strconv.Atoi(uIsNuked)
	}
	if uSmallLockAge != "" {
		updatedWorld.SmallLockAge, _ = strconv.Atoi(uSmallLockAge)
	}
	if uFloatPepperBlockCount != "" {
		updatedWorld.FloatPepperBlockCount, _ = strconv.Atoi(uFloatPepperBlockCount)
	}
	if uFloatPepperSeedCount != "" {
		updatedWorld.FloatPepperSeedCount, _ = strconv.Atoi(uFloatPepperSeedCount)
	}
	if uTilePepperSeedCount != "" {
		updatedWorld.TilePepperSeedCount, _ = strconv.Atoi(uTilePepperSeedCount)
	}
	if uTilePepperBlockCount != "" {
		updatedWorld.TilePepperBlockCount, _ = strconv.Atoi(uTilePepperBlockCount)
	}
	if uFossilCount != "" {
		updatedWorld.FossilCount, _ = strconv.Atoi(uFossilCount)
	}
	if uSLOwner != "" {
		updatedWorld.SLOwner = uSLOwner
	}
	if uBotHandlerId != "" {
		updatedWorld.BotHandlerId, _ = strconv.Atoi(uBotHandlerId)
	}
	if uGems != "" {
		updatedWorld.Gems, _ = strconv.Atoi(uGems)
	}
	updatedWorld.LastAccessed = int(time.Now().Unix())
	config.DB.Save(&updatedWorld)

	if result.Error != nil {
		return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", updatedWorld))
	} else {
		updatedWorld = existingWorld
		return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully updated", nil))
	}
}

func UpdateWorldLastAccess(c echo.Context) error {
	name := c.Param("name")

	var existingWorld model.World
	var updatedWorld model.World
	result := config.DB.Where("name = ?", name).First(&existingWorld)
	if result.Error != nil {

	} else {
		updatedWorld = existingWorld
	}
	updatedWorld.LastAccessed = int(time.Now().Unix())
	config.DB.Save(&updatedWorld)

	if result.Error != nil {
		return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", updatedWorld))
	} else {
		updatedWorld = existingWorld
		return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully updated", nil))
	}
}

func AssignBotToWorld(c echo.Context) error {
	var bots []model.Bot
	var worlds []model.World
	err := config.DB.Find(&bots).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bots"))
	}
	errW := config.DB.Find(&worlds).Error
	if errW != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve worlds"))
	}
	divider := len(worlds) / len(bots)
	if len(bots) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}
	var checkWorld model.World
	var updatedWorld model.World
	for _, bot := range bots {
		for i := 1; i <= divider; i++ {
			checkWorld = model.World{}
			updatedWorld = model.World{}
			err := config.DB.Where("bot_handler_id = ?", 0).First(&checkWorld).Error
			if err != nil {
				return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully assigned", nil))
			}
			updatedWorld = checkWorld
			updatedWorld.BotHandlerId = int(bot.ID)
			config.DB.Save(&updatedWorld)
		}
	}
	for _, bot := range bots {
		for i := 1; i <= divider; i++ {
			checkWorld = model.World{}
			updatedWorld = model.World{}
			err := config.DB.Where("bot_handler_id = ?", 0).First(&checkWorld).Error
			if err != nil {
				return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully assigned", nil))
			}
			updatedWorld = checkWorld
			updatedWorld.BotHandlerId = int(bot.ID)
			config.DB.Save(&updatedWorld)
		}
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully assigned", nil))
}

func AssignBotToWorldStorageSeed(c echo.Context) error {
	var bots []model.Bot
	var worlds []model.World
	//err := config.DB.Find(&bots).Error
	//err = config.DB.Where("owner = ?", "storage_seed").Find(&bots).Error
	err := config.DB.Where("is_suspended = ?", 0).Find(&bots).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bots"))
	}
	errW := config.DB.Find(&worlds).Error
	if errW != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve worlds"))
	}
	divider := len(worlds) / len(bots)
	if len(bots) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}
	var checkWorld model.World
	var updatedWorld model.World
	for _, bot := range bots {
		for i := 1; i <= divider; i++ {
			checkWorld = model.World{}
			updatedWorld = model.World{}
			err := config.DB.Where("bot_handler_id = ? and owner = ?", 0, "storage_seed").First(&checkWorld).Error
			if err != nil {
				return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully assigned", nil))
			}
			updatedWorld = checkWorld
			updatedWorld.BotHandlerId = int(bot.ID)
			config.DB.Save(&updatedWorld)
		}
	}
	for _, bot := range bots {
		for i := 1; i <= divider; i++ {
			checkWorld = model.World{}
			updatedWorld = model.World{}
			err := config.DB.Where("bot_handler_id = ? and owner = ?", 0, "storage_seed").First(&checkWorld).Error
			if err != nil {
				return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully assigned", nil))
			}
			updatedWorld = checkWorld
			updatedWorld.BotHandlerId = int(bot.ID)
			config.DB.Save(&updatedWorld)
		}
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully assigned", nil))
}

func AssignBotToWorldStorageSeedOneHundredBotOnly(c echo.Context) error {
	var bots []model.Bot
	var worlds []model.World
	//err := config.DB.Find(&bots).Error
	//err = config.DB.Where("owner = ?", "storage_seed").Find(&bots).Error
	//limit the bots result to 100
	err := config.DB.Where("is_suspended = ?", 0).Limit(100).Find(&bots).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bots"))
	}
	//config.DB.Where("bot_handler_id = ? and owner = ?", 0, "storage_seed")
	errW := config.DB.Where("bot_handler_id = ? and owner = ?", 0, "storage_seed").Find(&worlds).Error
	if errW != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve worlds"))
	}
	divider := len(worlds) / len(bots)
	if len(bots) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}
	var checkWorld model.World
	var updatedWorld model.World
	for _, bot := range bots {
		for i := 1; i <= divider; i++ {
			checkWorld = model.World{}
			updatedWorld = model.World{}
			err := config.DB.Where("bot_handler_id = ? and owner = ?", 0, "storage_seed").First(&checkWorld).Error
			if err != nil {
				return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully assigned", bots))
			}
			updatedWorld = checkWorld
			updatedWorld.BotHandlerId = int(bot.ID)
			config.DB.Save(&updatedWorld)
		}
	}
	for _, bot := range bots {
		for i := 1; i <= divider; i++ {
			checkWorld = model.World{}
			updatedWorld = model.World{}
			err := config.DB.Where("bot_handler_id = ? and owner = ?", 0, "storage_seed").First(&checkWorld).Error
			if err != nil {
				return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully assigned", bots))
			}
			updatedWorld = checkWorld
			updatedWorld.BotHandlerId = int(bot.ID)
			config.DB.Save(&updatedWorld)
		}
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully assigned", bots))
}

func UnassignBotToWorld(c echo.Context) error {
	var worlds []model.World
	errW := config.DB.Find(&worlds).Error
	if errW != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve worlds"))
	}
	var updatedWorld model.World
	for _, world := range worlds {
		updatedWorld = model.World{}
		updatedWorld = world
		updatedWorld.BotHandlerId = 0
		config.DB.Save(&updatedWorld)
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully unassigned", nil))
}

func UnassignBotToWorldFast(c echo.Context) error {
	var worlds []model.World
	//raw query
	errW := config.DB.Raw("UPDATE worlds SET bot_handler_id = 0").Scan(&worlds).Error
	if errW != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve worlds"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully unassigned", nil))
}

func GetWorldByBotHandlerId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var worlds []model.World
	//find all worlds where bot_handler_id = id
	err = config.DB.Where("bot_handler_id = ?", id).Find(&worlds).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}

	response := res.ConvertIndexWorldNameIdOnly(worlds)
	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", response))
}

func GetWorldByBotName(c echo.Context) error {
	growid := c.Param("growid")
	var worlds []model.World
	var bot model.Bot
	err := config.DB.Where("growid = ?", growid).First(&bot).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bot"))
	}
	//find all worlds where bot_handler_id = id
	err = config.DB.Where("bot_handler_id = ?", bot.ID).Find(&worlds).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}

	response := res.ConvertIndexWorld(worlds)
	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", response))
}

func GetWorldTypeStorageSeedThatHasBiggestFloatingPepperSeed(c echo.Context) error {
	var world model.World
	err := config.DB.Where("type = ? AND sl_owner != ?", "storage_seed", "notfound").Order("float_pepper_seed_count desc").First(&world).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", world))
}

func GetWorldTypeStorageSeedThatHasBiggestFloatingPepperSeedUnderXK(c echo.Context) error {
	var world model.World
	XK := c.Param("xk")
	err := config.DB.Where("type = ? AND sl_owner != ? AND float_pepper_seed_count < ?", "storage_seed", "notfound", XK).Order("float_pepper_seed_count desc").First(&world).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", world))
}

func GetWorldTypeStorageSeedThatHasSmallestFloatingPepperSeed(c echo.Context) error {
	var existingWorld model.World
	var updatedWorld model.World
	//Where("type = ? AND sl_owner = ?", "storage_seed", "notfound") add condition where "last_accessed" - time.Now().Unix() > 120
	currentTime := time.Now().Unix()
	err := config.DB.Where("type = ? AND float_pepper_seed_count > ? AND sl_owner = ? AND ? - last_accessed > ?", "storage_seed", 0, "notfound", currentTime, 10).Order("float_pepper_seed_count asc").First(&existingWorld).Error
	if err != nil {
		errorSecond := config.DB.Where("type = ? AND float_pepper_seed_count > ? AND ? - last_accessed > ?", "storage_seed", 0, currentTime, 10).Order("float_pepper_seed_count asc").First(&existingWorld).Error
		if errorSecond != nil {
			return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
		} else {
			updatedWorld = existingWorld
			updatedWorld.LastAccessed = int(time.Now().Unix())
			errUpdate := config.DB.Save(&updatedWorld).Error
			if errUpdate != nil {
				return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update world"))
			}
			return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", updatedWorld))
		}
	} else {
		updatedWorld = existingWorld
		updatedWorld.LastAccessed = int(time.Now().Unix())
		errUpdate := config.DB.Save(&updatedWorld).Error
		if errUpdate != nil {
			return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update world"))
		}
		return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", updatedWorld))
	}
}

func DeleteWorld(c echo.Context) error {
	name := c.Param("name")
	var existingWorld model.World
	err := config.DB.Where("name = ?", name).First(&existingWorld).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}

	config.DB.Delete(&existingWorld)

	return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully deleted", nil))
}
