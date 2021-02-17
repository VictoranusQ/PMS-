//go:generate wire
//+build wireinject

package app

import (
	"PMSApp/app/basic"
	"PMSApp/app/daos"
	"PMSApp/app/models"
	"PMSApp/app/utils"
	"net/http"

	"github.com/google/wire"
	"PMSApp/app/config"
	"PMSApp/app/controllers"
	"PMSApp/app/routers"
)

func BuildInjector() *http.Server {
	wire.Build(
		config.NewConfig,
		basic.LoggerSet,
		basic.NewDB,
		utils.GinxSet,
		daos.DaoSet,
		models.ModelSet,
		controllers.ControllerSet,
		routers.RouterSet,
		NewApp,
	)

	return nil
}
