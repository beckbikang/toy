package main

import (
	"flag"
	"toy/marmot/web/query-log/lauch"
)

var (
	confRoot = flag.String("path", "./conf", "root path of configuration")
	env      = flag.String("env", "dev", "prod or dev or test")
)

func main() {
	flag.Parse()
	lauch.InitLauch(*confRoot, *env)
}
