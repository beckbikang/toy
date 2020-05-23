package engine

import (
	"time"
	"net/http"
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

//get http server
func InitHttpServer()*http.Server{

	//init engine
	eg := InitEngine()

	addr := config.Gcfg.GetString("server.addr")
	readTimeout := config.Gcfg.GetInt("server.read_timeout")
	writeTimeout := config.Gcfg.GetInt("server.write_timeout")
	maxHeaderBytes:= config.Gcfg.GetInt("server.max_body_bytes")


	//init server
	srv := &http.Server{
		Addr: addr,
		Handler:        eg,
		ReadTimeout:   time.Duration(readTimeout)  * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	return srv
}