package lauch

import (
	"context"
	"log"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"toy/marmot/web/query-log/lauch/config"
	"toy/marmot/web/query-log/lauch/engine"
	"flag"
)

var (
	confRoot = flag.String("path", "./conf", "root path of configuration")
	env      = flag.String("env", "dev", "prod or dev or test")
)

func  InitLauch(){
	flag.Parse()

	config.LoadGlobalConfig(*confRoot, *env)

	//get server
	srv := engine.InitHttpServer()

	//run
	go func(){
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//signal
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	//context
	contextTimeout := config.Gcfg.GetInt("server.cancel_timeout")
	ctx,cancel := context.WithTimeout(context.Background(),
		time.Duration(contextTimeout)*time.Second)
	defer cancel()

	//shutdown
	if err := srv.Shutdown(ctx);err != nil{
		log.Fatalf("shutdown error :+%v", err)
	}

	//timeout
	select {
	case <-ctx.Done():
		log.Printf("timeout of %d seconds", contextTimeout)
	}

	//stop
	log.Println("server stop")
}






