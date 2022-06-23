package main

import (
	"go-architecture/api/router"
	"go-architecture/config"
	"go-architecture/database"
	"go-architecture/i18n"
)

type Infrastructure struct {
	Config   config.Config
	I18n     *i18n.I18n
	Database *database.Database
	Router   *router.Router
}
