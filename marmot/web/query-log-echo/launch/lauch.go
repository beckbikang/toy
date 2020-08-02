package launch

import (
	"flag"
	"toy/marmot/web/query-log-echo/launch/config"
	"toy/marmot/web/query-log-echo/launch/db"
	"toy/marmot/web/query-log-echo/launch/engine"
	kl "toy/marmot/web/query-log-echo/launch/log"
	"toy/marmot/web/query-log-echo/launch/cache"
)

var (
	confRoot = flag.String("path", "./conf", "root path of configuration")
	env      = flag.String("env", "dev", "prod or dev or test")
)

func InitLaunch() {
	kl.InitLog()
	flag.Parse()
	config.LoadGlobalConfig(*confRoot, *env)
	db.InitDb()
	engine.InitLaunchHttpServer()
	cache.InitRedisPool()


	kl.LOGGER.Info("init launch ok")

	//defer
	defer func() {
		db.GetDb().Close()
		kl.LOGGER.Sync()
	}()

	kl.LOGGER.Info("server end at launch")
}
