package launch

import (
	"flag"
	"fmt"
	"toy/marmot/web/query-log/launch/config"
	"toy/marmot/web/query-log/launch/db"
	"toy/marmot/web/query-log/launch/engine"
	kl "toy/marmot/web/query-log/launch/log"
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
	fg := kl.InitLog()

	//defer
	defer func() {
		db.GetDb().Close()
		fg.Sync()
	}()

	fmt.Println("server end at launch")
}
