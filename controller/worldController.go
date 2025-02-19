package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"app/utils/res"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var min_last_accessed_diff int = 2

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

func GetWorldOneFieldInfoByWorldName(c echo.Context) error {
	name := c.Param("name")
	field := c.Param("field")
	err := config.DB.Model(&model.World{}).Where("name = ?", name).Select(field).First(&name).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("One field: "+field+" data successfully retrieved", name))
}

func GetOneWorldWithCustomWhere(c echo.Context) error {
	field := c.Param("field")
	operator := c.Param("operator")
	value := c.Param("value")
	var existingWorld model.World
	// Construct the query based on the operator
	query := fmt.Sprintf("%s %s ?", field, operator)

	// Execute the query
	err := config.DB.Where(query, value).First(&existingWorld).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("One world data where "+field+" "+operator+" "+value+" successfully retrieved", existingWorld))
}

func GetOneWorldWithCustomWhereV2(c echo.Context) error {
	//.Order("float_pepper_seed_count desc")
	where := c.Param("where")
	fieldSort := c.Param("fieldSort")
	typeSort := c.Param("typeSort")
	var existingWorld model.World
	// Construct the query based on the operator
	query := fmt.Sprintf("%s", where)
	querySort := fmt.Sprintf("%s %s", fieldSort, typeSort)

	// Execute the query
	err := config.DB.Where(query).Order(querySort).First(&existingWorld).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("One world data where "+where+" successfully retrieved", existingWorld))
}

func GetOneWorldWithCustomWhereV3(c echo.Context) error {
	var customWhere model.CustomWhere
	if err := c.Bind(&customWhere); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}
	var existingWorld model.World
	// Construct the query based on the operator
	query := fmt.Sprintf("%s", customWhere.Where)
	querySort := fmt.Sprintf("%s %s", customWhere.FieldSort, customWhere.TypeSort)

	// Execute the query
	err := config.DB.Where(query).Order(querySort).First(&existingWorld).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("One world data where "+customWhere.Where+" successfully retrieved", existingWorld))
}

func GetCountWorldCustomWhere(c echo.Context) error {
	var customWhere model.CustomWhere
	var count int64
	/*
		err := config.DB.Model(&model.World{}).Where("tile_pepper_seed_count > ?", 1000).Count(&count).Error
		if err != nil {
		    // handle error
		    log.Println("Error counting worlds:", err)
		} else {
		    log.Println("Count of worlds with tile_pepper_seed_count > 1000:", count)
		}
	*/
	if err := c.Bind(&customWhere); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}
	// Construct the query based on the operator
	query := fmt.Sprintf("%s", customWhere.Where)
	err := config.DB.Model(&model.World{}).Where(query).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("One world data where "+customWhere.Where+" successfully retrieved", count))
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
	uProblem := c.QueryParam("uProblem")

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
	if uProblem != "" {
		updatedWorld.Problem = uProblem
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
func UpdateWorldVer2(c echo.Context) error {
	var world model.World
	if err := c.Bind(&world); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var existingWorld model.World
	var updatedWorld model.World
	result := config.DB.Where("name = ?", world.Name).First(&existingWorld)
	updatedWorld = world
	updatedWorld.LastAccessed = int(time.Now().Unix())
	if result.Error != nil {
		//data tidak ada di database
		errCreate := config.DB.Create(&updatedWorld).Error
		if errCreate != nil {
			return c.JSON(http.StatusBadRequest, errCreate)
		} else {
			return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Creating Data", true))
		}
	} else {
		//data ada di database
		//config.DB.Model(&existingArticle).Updates(updatedArticle)
		errUpdate := config.DB.Model(&existingWorld).Updates(updatedWorld).Error
		if errUpdate != nil {
			return c.JSON(http.StatusBadRequest, errUpdate)
		} else {
			return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Updating Data", true))
		}
	}
}

func UpdateWorldLastAccess(c echo.Context) error {
	name := c.Param("name")

	var existingWorld model.World
	var updatedWorld model.World
	result := config.DB.Where("name = ?", name).First(&existingWorld)
	if result.Error != nil {
		updatedWorld.Name = name
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

func UpdateResetWorldBotHandlerId(c echo.Context) error {
	name := c.Param("name")

	var existingWorld model.World
	var updatedWorld model.World
	result := config.DB.Where("name = ?", name).First(&existingWorld)
	if result.Error != nil {
		updatedWorld.Name = name
	} else {
		updatedWorld = existingWorld
	}
	updatedWorld.BotHandlerId = 0
	config.DB.Save(&updatedWorld)

	if result.Error != nil {
		return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", updatedWorld))
	} else {
		updatedWorld = existingWorld
		return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully updated", nil))
	}
}

func UpdateWorldProblem(c echo.Context) error {
	name := c.Param("name")
	problem := c.Param("problem")

	var existingWorld model.World
	var updatedWorld model.World
	result := config.DB.Where("name = ?", name).First(&existingWorld)
	if result.Error != nil {
		updatedWorld.Name = name
	} else {
		updatedWorld = existingWorld
	}
	updatedWorld.Problem = problem
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
			err := config.DB.Where("bot_handler_id = ? AND type = ?", 0, "farm").First(&checkWorld).Error
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
	err := config.DB.Where("type = ? AND float_pepper_seed_count > ? AND float_pepper_seed_count < ? AND sl_owner = ? AND ? - last_accessed > ?", "storage_seed", 0, 100000, "notfound", currentTime, 2).Order("float_pepper_seed_count asc").First(&existingWorld).Error
	if err != nil {
		errorSecond := config.DB.Where("type = ? AND float_pepper_seed_count > ? AND float_pepper_seed_count < ? AND ? - last_accessed > ?", "storage_seed", 0, 100000, currentTime, 2).Order("float_pepper_seed_count asc").First(&existingWorld).Error
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
func GetWorldTypeStorageSeedThatHasSmallestFloatingPepperSeedTypeAll(c echo.Context) error {
	var existingWorld model.World
	var updatedWorld model.World
	//Where("type = ? AND sl_owner = ?", "storage_seed", "notfound") add condition where "last_accessed" - time.Now().Unix() > 120
	currentTime := time.Now().Unix()
	errorSecond := config.DB.Where("float_pepper_seed_count > ? AND float_pepper_seed_count < ? AND ? - last_accessed > ? AND type = ?", 0, 47500, currentTime, 3, "storage_seed").Order("float_pepper_seed_count asc").First(&existingWorld).Error
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
}

func GetAndSetWorldThatHasBiggestFloatingBlock(c echo.Context) error {
	growid := c.Param("growid")
	var bot model.Bot
	var existingWorld model.World
	var updatedWorld model.World
	errGetBot := config.DB.Where("growid = ?", growid).First(&bot).Error
	if errGetBot != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bot"))
	}
	//find one world where bot_handler_id = id
	errGetBotWorld := config.DB.Where("bot_handler_id = ?", bot.ID).First(&existingWorld).Error
	if errGetBotWorld != nil {
		//if error get world by bot handler id, then we will find world with the biggest float_pepper_block_count
		//query to find first world with the biggest float_pepper_block_count with condition where float_pepper_block_count > 0 and bot_handler_id == 0
		currentTime := time.Now().Unix()
		errGetWorldHasOneOrMoreFloatBlock := config.DB.Where("float_pepper_block_count > ? AND bot_handler_id = ? AND ? - last_accessed > ? AND is_nuked = ? AND type = ?", 0, 0, currentTime, min_last_accessed_diff, 0, "farm").Order("float_pepper_block_count desc").First(&existingWorld).Error
		if errGetWorldHasOneOrMoreFloatBlock != nil {
			//if error get world that has floating block min 1, then we will find world with the biggest tile_pepper_seed_count
			//query to find first world with the biggest tile_pepper_seed_count with condition bot_handler_id == 0
			errGetWorldMore := config.DB.Where("bot_handler_id = ? AND ? - last_accessed > ? AND is_nuked = ? AND type = ?", 0, currentTime, min_last_accessed_diff, 0, "farm").Order("tile_pepper_seed_count desc").First(&existingWorld).Error
			if errGetWorldMore != nil {
				return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
			} else {
				updatedWorld = existingWorld
				updatedWorld.BotHandlerId = int(bot.ID)
				errUpdate := config.DB.Save(&updatedWorld).Error
				if errUpdate != nil {
					return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update world"))
				}
				return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", updatedWorld))
			}
		} else {
			updatedWorld = existingWorld
			updatedWorld.BotHandlerId = int(bot.ID)
			errUpdate := config.DB.Save(&updatedWorld).Error
			if errUpdate != nil {
				return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update world"))
			}
			return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", updatedWorld))
		}
	} else {
		return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", existingWorld))
	}
}

func GetAndSetWorldThatHasSmallestTilePepperSeed(c echo.Context) error {
	//This is for owner api, coz there is diff in bot_handler_id set , it is set to 70000 + bot.ID
	growid := c.Param("growid")
	var bot model.Bot
	var existingWorld model.World
	var updatedWorld model.World
	currentTime := time.Now().Unix()
	errGetBot := config.DB.Where("growid = ?", growid).First(&bot).Error
	if errGetBot != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bot"))
	}
	//find one world where bot_handler_id = id
	errGetBotWorld := config.DB.Where("bot_handler_id = ?", 70000+bot.ID).First(&existingWorld).Error
	if errGetBotWorld != nil {
		errGetWorldMore := config.DB.Where("bot_handler_id = ? AND ? - last_accessed > ? AND is_nuked = ? AND tile_pepper_seed_count < ? AND type = ?", 0, currentTime, min_last_accessed_diff, 0, 2451, "farm").Order("tile_pepper_seed_count asc").First(&existingWorld).Error
		if errGetWorldMore != nil {
			return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
		} else {
			updatedWorld = existingWorld
			updatedWorld.BotHandlerId = int(70000 + bot.ID)
			errUpdate := config.DB.Save(&updatedWorld).Error
			if errUpdate != nil {
				return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update world"))
			}
			return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", updatedWorld))
		}
	} else {
		return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", existingWorld))
	}
}

func GetOwnerWorld(c echo.Context) error {
	growid := c.Param("growid")
	var bot model.Bot
	var botGroupOwner model.Bot
	var existingWorld model.World
	errGetBot := config.DB.Where("growid = ?", growid).First(&bot).Error
	if errGetBot != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bot"))
	}
	errGetBotGroupOwner := config.DB.Where("growid = ?", bot.GroupOwner).First(&botGroupOwner).Error
	if errGetBotGroupOwner != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve bot group owner"))
	}
	//not +70k because we find for the break world, not the want to plant seed world
	//find break world
	errGetBotWorld := config.DB.Where("bot_handler_id = ?", botGroupOwner.ID).First(&existingWorld).Error
	if errGetBotWorld != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve world"))
	} else {
		return c.JSON(http.StatusOK, utils.SuccessResponse("World data successfully retrieved", existingWorld))
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
