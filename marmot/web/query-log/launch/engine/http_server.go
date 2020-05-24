package engine

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"toy/marmot/web/query-log/launch/config"
)

func InitLauchHttpServer() {

	//get server
	srv := InitHttpServer()

	//run
	go func() {
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
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(contextTimeout)*time.Second)
	defer cancel()

	//shutdown
	if err := srv.Shutdown(ctx); err != nil {
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

//get http server
func InitHttpServer() *http.Server {

	//init engine
	eg := InitEngine()

	addr := config.Gcfg.GetString("server.addr")
	readTimeout := config.Gcfg.GetInt("server.read_timeout")
	writeTimeout := config.Gcfg.GetInt("server.write_timeout")
	maxHeaderBytes := config.Gcfg.GetInt("server.max_body_bytes")

	//init server
	srv := &http.Server{
		Addr:           addr,
		Handler:        eg,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	return srv
}
