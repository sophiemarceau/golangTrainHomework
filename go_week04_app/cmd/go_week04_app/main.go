package main

import (
	"go_week04_app/internal/conf"
	"go_week04_app/internal/fbase"
	"go_week04_app/internal/server"
	"log"
	"net/http"
)

func main() {
	config := &conf.Config{}
	err := config.LoadFile("./configs/application.yml")
	if err != nil {
		log.Printf("%+v\n",err)
		return
	}
	app := initApp(config, log.Default())
	if err := app.Start(); err != nil {
		log.Printf("%+v\n", err)
		return
	}
}

func newApp(logger *log.Logger, server0 *server.HttpServer0, server1 *server.HttpServer1) *fbase.App {
	return fbase.New(fbase.Name("go_week04"), fbase.Logger(logger), fbase.Servers((*http.Server)(server0), (*http.Server)(server1)))
}


