package repository

import (
	"github.com/google/wire"
	"go-architecture/database"
	repository "go-architecture/repository/repo.mapper"
)

var ProviderSet = wire.NewSet(ProviderCompanyCapacityRepository)

func ProviderCompanyCapacityRepository(db database.Repository, mapper repository.CompanyCapacityMapper) CompanyCapacityRepository {
	return companyCapacityRepository{db, mapper}
}
