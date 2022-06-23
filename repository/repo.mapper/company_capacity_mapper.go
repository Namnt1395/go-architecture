package repository

import (
	"github.com/jinzhu/copier"
	dto "go-architecture/dto/company_capacity"
	"go-architecture/entities"
	util "go-architecture/pkg"
)

type CompanyCapacityMapper struct {
}

func (m CompanyCapacityMapper) ConvertToEntity(capacityInput *dto.CapacityInput) *entities.CompanyCapacity {
	var companyCapacity entities.CompanyCapacity
	err1 := copier.Copy(&companyCapacity, &capacityInput)
	util.Must(err1)
	return &companyCapacity
}
