package engine

import (
	"toy/marmot/web/query-log-echo/launch/config"
	"toy/marmot/web/query-log-echo/launch/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	DebugDesc = "debug"
)


func InitEngine() *echo.Echo {



	engine := echo.New()

	//init middleware
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())

	//debug
	engine.Debug = false
	if DebugDesc == config.Gcfg.GetString("server.run_mode"){
		engine.Debug = true
	}


	//init router
	router.InitRouter(engine)

	return engine
}
