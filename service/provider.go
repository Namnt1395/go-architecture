package service

import (
	"github.com/google/wire"
	"go-architecture/repository"
	service_mapper "go-architecture/service/service.mapper"
)

var ProvideSet = wire.NewSet(ProvideCompanyCapacityService)

func ProvideCompanyCapacityService(rp repository.CompanyCapacityRepository, mapper service_mapper.CompanyCapacityMapper) CompanyCapacityService {
	return companyCapacityService{repo: rp, mapper: mapper}
}