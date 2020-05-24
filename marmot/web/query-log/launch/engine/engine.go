package engine

import (
	"toy/marmot/web/query-log/launch/config"
	"toy/marmot/web/query-log/launch/router"

	"github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {

	engine := gin.New()

	//init middleware
	engine.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(config.Gcfg.GetString("server.run_mode"))

	//init router
	router.InitRouter(engine)

	return engine
}
