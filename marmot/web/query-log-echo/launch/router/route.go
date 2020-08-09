package router

import (

	"github.com/labstack/echo/v4"
	"toy/marmot/web/query-log-echo/controller"
)


//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
// swagger:meta
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