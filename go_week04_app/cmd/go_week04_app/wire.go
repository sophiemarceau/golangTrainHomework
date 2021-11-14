//+build wireinject
package main

import (
	"github.com/google/wire"
	"go_week04_app/internal/biz"
	"go_week04_app/internal/conf"
	"go_week04_app/internal/data"
	"go_week04_app/internal/fbase"
	"go_week04_app/internal/server"
	"log"
)

func initApp(config *conf.Config, logger *log.Logger) *fbase.App {
	wire.Build(server.WireSet, data.WireSet, biz.WireSet, server.WireSet)
	return &fbase.App{}
}
