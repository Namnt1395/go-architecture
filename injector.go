//+build wireinject

package main

import (
	"github.com/google/wire"
	"go-architecture/api/router"
	"go-architecture/config"
	"go-architecture/database"
	"go-architecture/i18n"
	repository "go-architecture/repository/repo.mapper"
	service_mapper "go-architecture/service/service.mapper"
)

func InitApp() (App, error) {
	panic(wire.Build(
		// infrastructure
		config.LoadConfig,
		database.NewDatabase,
		i18n.NewI18n,
		router.NewRouterWithAuthMw,
		wire.Struct(new(repository.CompanyCapacityMapper), "*"),
		wire.Struct(new(service_mapper.CompanyCapacityMapper), "*"),
		wire.Struct(new(database.Repository), "*"),
		wire.Struct(new(Infrastructure), "*"),
		// app
		wire.Struct(new(App), "*"),
		ProviderSet,
	))
	return App{}, nil
}
