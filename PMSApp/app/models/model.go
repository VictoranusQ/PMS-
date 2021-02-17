package models

import "github.com/google/wire"

var ModelSet = wire.NewSet(
	UserSet,
)
