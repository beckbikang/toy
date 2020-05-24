package lauch

import (
	"flag"
	"fmt"
	"toy/marmot/web/query-log/lauch/config"
	"toy/marmot/web/query-log/lauch/db"
	"toy/marmot/web/query-log/lauch/engine"
)

var (
	confRoot = flag.String("path", "./conf", "root path of configuration")
	env      = flag.String("env", "dev", "prod or dev or test")
)

func InitLauch() {
	flag.Parse()
	config.LoadGlobalConfig(*confRoot, *env)
	db.InitDb()
	engine.InitLauchHttpServer()

	//defer
	defer func() {
		db.GetDb().Close()
	}()

	fmt.Println("server end at lauch")
}
