package handle

import (
	"github.com/google/wire"
	"go-architecture/service"
)

var ProvideSet = wire.NewSet(ProvideCompanyCapacityHandler)

func ProvideCompanyCapacityHandler(srv service.CompanyCapacityService) CompanyCapacityHandler {
	return companyCapacityHandler{srv}
}
