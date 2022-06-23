package service

import (
	"github.com/jinzhu/copier"
	dto "go-architecture/dto/company_capacity"
	"go-architecture/entities"
	util "go-architecture/pkg"
)

type CompanyCapacityMapper struct {
}

func (m CompanyCapacityMapper) ConvertToDto(capacityInput entities.CompanyCapacity) *dto.CapacityOutput {
	var companyCapacity dto.CapacityOutput
	err1 := copier.Copy(&companyCapacity.Data, &capacityInput)
	util.Must(err1)
	return &companyCapacity
}

func (m CompanyCapacityMapper) ConvertToListDto(capacityEntity []entities.CompanyCapacity) *dto.ListCapacityOutput {
	var listCompanyCapacity dto.ListCapacityOutput
	err1 := copier.Copy(&listCompanyCapacity.Content, &capacityEntity)
	util.Must(err1)
	return &listCompanyCapacity
}
