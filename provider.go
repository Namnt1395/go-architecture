package main

import (
	"github.com/google/wire"
	"go-architecture/api/handle"
	"go-architecture/repository"
	"go-architecture/service"
)

var ProviderSet = wire.NewSet(
	repository.ProviderSet,
	service.ProvideSet,
	handle.ProvideSet,
)
