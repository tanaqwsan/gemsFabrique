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
	e.GET("/worldsBiggestSeeds", controller.GetWorldTypeStorageSeedThatHasBiggestFloatingPepperSeed)
	e.GET("/worldsBiggestSeedsXK/:xk", controller.GetWorldTypeStorageSeedThatHasBiggestFloatingPepperSeedUnderXK)
	e.GET("/worldsSmallestSeeds", controller.GetWorldTypeStorageSeedThatHasSmallestFloatingPepperSeed)
	e.GET("/worldsSmallestSeedsTypeAll", controller.GetWorldTypeStorageSeedThatHasSmallestFloatingPepperSeedTypeAll)
	e.PUT("/worlds/:name", controller.UpdateWorld)
	e.PUT("/worldsAccess/:name", controller.UpdateWorldLastAccess)
	e.DELETE("/worlds/:name", controller.DeleteWorld)

	//Manage Bot
	e.POST("/bots", controller.StoreBot)
	e.GET("/bots", controller.IndexBot)
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
