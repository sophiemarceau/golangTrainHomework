package biz

import "github.com/google/wire"

var WireSet = wire.NewSet(initUserBiz)