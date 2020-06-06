package launch

import (
	"flag"
	"toy/marmot/web/query-log-echo/launch/config"
	"toy/marmot/web/query-log-echo/launch/db"
	"toy/marmot/web/query-log-echo/launch/engine"
	kl "toy/marmot/web/query-log-echo/launch/log"
)

var (
	confRoot = flag.String("path", "./conf", "root path of configuration")
	env      = flag.String("env", "dev", "prod or dev or test")
)

func InitLaunch() {
	flag.Parse()
	config.LoadGlobalConfig(*confRoot, *env)
	db.InitDb()
	engine.InitLauchHttpServer()
	kl.InitLog()

	//defer
	defer func() {
		db.GetDb().Close()
		kl.LOGGER.Sync()
	}()

	kl.LOGGER.Info("server end at launch")
}