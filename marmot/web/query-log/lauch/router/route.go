package router

import (
	"github.com/gin-gonic/gin"
	"toy/marmot/web/query-log/controller"
)


func InitRouter(engine*gin.Engine){

	v2 := engine.Group("/v2")
	{
		//v2.GET("/log/filted/lists.json",)
	}

	//v3
	v3 := engine.Group("/v3")
	{
		v3.GET("/ping", controller.Ping)
		//v3.GET("/log/filted/lists.json",)
	}

}