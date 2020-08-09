package main

import (
	"toy/marmot/web/query-log-echo/cmd"
 	"log"
 	)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
