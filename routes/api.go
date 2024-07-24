package routes

import (
	"app/controller"
	"app/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services test")
	})

	//Manage World
	e.POST("/worlds", controller.StoreWorld)
	e.GET("/worlds", controller.IndexWorld)
	e.GET("/worldsName", controller.IndexWorldOnlyName)
	e.GET("/worlds/:name", controller.ShowWorld)
	e.GET("/worldsOneField/:name/:field", controller.GetWorldOneFieldInfoByWorldName)
	e.GET("/worldsCustomWhere/:field/:operator/:value", controller.GetOneWorldWithCustomWhere)
	e.GET("/worldsCustomWhereV2/:where/:fieldSort/:typeSort", controller.GetOneWorldWithCustomWhereV2)
	e.GET("/worldsCustomWhereV3", controller.GetOneWorldWithCustomWhereV3)
	e.GET("/worldsCustomWhereVCount", controller.GetCountWorldCustomWhere)
	e.GET("/worldsOwner/:growid", controller.GetOwnerWorld)
	e.GET("/worldsBiggestSeeds", controller.GetWorldTypeStorageSeedThatHasBiggestFloatingPepperSeed)
	e.GET("/worldsBiggestSeedsXK/:xk", controller.GetWorldTypeStorageSeedThatHasBiggestFloatingPepperSeedUnderXK)
	e.GET("/worldsSmallestSeeds", controller.GetWorldTypeStorageSeedThatHasSmallestFloatingPepperSeed)
	e.GET("/worldsSmallestSeedsTypeAll", controller.GetWorldTypeStorageSeedThatHasSmallestFloatingPepperSeedTypeAll)

	//For collecting block to break in world break
	e.PUT("/worldsGetAndSetBiggestFloatingBlock/:growid", controller.GetAndSetWorldThatHasBiggestFloatingBlock)

	//For finding world that has the smallest tile_pepper_seed_count then assign it to current bot (:growid)
	e.PUT("/worldsGetAndSetSmallestTilePepperSeed/:growid", controller.GetAndSetWorldThatHasSmallestTilePepperSeed)

	e.PUT("/worlds/:name", controller.UpdateWorld)
	e.PUT("/worlds2", controller.UpdateWorldVer2)
	e.PUT("/worldsAccess/:name", controller.UpdateWorldLastAccess)
	e.PUT("/worldsHandlerReset/:name", controller.UpdateResetWorldBotHandlerId)
	e.PUT("/worldsProblem/:name/:problem", controller.UpdateWorldProblem)
	e.DELETE("/worlds/:name", controller.DeleteWorld)

	//Manage Bot
	e.POST("/bots", controller.StoreBot)
	e.GET("/bots", controller.IndexBot)
	e.GET("/bots/owner/:group_owner", controller.IndexBotByGroupOwner)
	e.GET("/bots/worlds/:id", controller.GetWorldByBotHandlerId)
	e.GET("/bots/worlds/growid/:growid", controller.GetWorldByBotName)
	e.GET("/bots/:growid", controller.ShowBot)
	e.PUT("/bots/:growid", controller.UpdateBot)
	e.DELETE("/bots/:id", controller.DeleteBot)
	e.PUT("/bots/assignWorlds", controller.AssignBotToWorld)
	e.PUT("/bots/assignWorldsStorageSeed", controller.AssignBotToWorldStorageSeed)
	e.PUT("/bots/assignWorldsStorageSeed100bot", controller.AssignBotToWorldStorageSeedOneHundredBotOnly)
	e.PUT("/bots/unassignWorlds", controller.UnassignBotToWorld)
	e.PUT("/bots/unassignWorldsFast", controller.UnassignBotToWorldFast)

	//Manage Word
	e.GET("/words", controller.IndexWord)
	e.GET("/words/:growid", controller.ShowWord)
	e.PUT("/words/:growid/:word", controller.UpdateWord)
	//e.PUT("/words/:growid/targetxzx?uTarget=", controller.UpdateWord)
	e.PUT("/words/erase/:growid", controller.EraseWord)

	return e

}
