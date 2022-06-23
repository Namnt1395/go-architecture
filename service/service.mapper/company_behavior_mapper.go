package service

import (
	"github.com/jinzhu/copier"
	dto "go-architecture/dto/company_behavior"
	"go-architecture/entities"
	util "go-architecture/pkg"
)

type CompanyBehaviorMapper struct {
}

func (m CompanyBehaviorMapper) ConvertToDto(behaviorEntity entities.CompanyBehavior) dto.BehaviorOutput {
	var companyBehavior dto.BehaviorOutput
	err1 := copier.Copy(&companyBehavior, &behaviorEntity)
	util.Must(err1)
	return companyBehavior
}

func (m CompanyBehaviorMapper) ConvertToListDto(behaviorEntity entities.CompanyBehavior) []dto.BehaviorOutput {
	var listCompanyBehavior []dto.BehaviorOutput
	err1 := copier.Copy(&listCompanyBehavior, &behaviorEntity)
	util.Must(err1)
	return listCompanyBehavior
}
