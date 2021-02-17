package controllers

import "github.com/google/wire"

// ControllersSet 控制器 DI
var ControllerSet = wire.NewSet(
	LoginSet,
)
