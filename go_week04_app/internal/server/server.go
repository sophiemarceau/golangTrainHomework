package server

import "github.com/google/wire"

var WireSet = wire.NewSet(InitHttpServer0, InitHttpServer1)