//go:generate wire
//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/CommercialManagementSystem/back/internal/controller"
	"github.com/CommercialManagementSystem/back/internal/dao"
	"github.com/CommercialManagementSystem/back/internal/model"
	"github.com/CommercialManagementSystem/back/internal/router"
	"github.com/google/wire"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGorm,
		controller.ControllerSet,
		model.ModelSet,
		dao.DaoSet,
		router.RouterSet,
		InitGinEngine,
		InjectorSet,
	)

	return nil, nil, nil
}
