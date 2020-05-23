package engine

import (
	"github.com/gin-gonic/gin"
	"toy/marmot/web/query-log/lauch/config"

	"toy/marmot/web/query-log/lauch/router"
)

func InitEngine()*gin.Engine{

	engine := gin.New()

	//init middleware
	engine.Use(gin.Logger(),gin.Recovery())
	gin.SetMode(config.Gcfg.GetString("server.run_mode"))

	//init router
	router.InitRouter(engine)

	return engine
}