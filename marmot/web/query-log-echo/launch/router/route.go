package router

import (

	"github.com/labstack/echo/v4"
	"toy/marmot/web/query-log-echo/controller"
)


func InitRouter(engine*echo.Echo){

	v2 := engine.Group("/v2")
	{
		v2.GET("/ping", controller.Ping)
	}

	//v3
	v3 := engine.Group("/v3")
	{
		v3.GET("/ping", controller.Ping)
		v3.GET("/log/filtered/lists.json",controller.OpLogList)
	}

}