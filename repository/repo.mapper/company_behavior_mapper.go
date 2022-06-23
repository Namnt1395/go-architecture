package repository

import (
	"github.com/jinzhu/copier"
	dto "go-architecture/dto/company_behavior"
	"go-architecture/entities"
	util "go-architecture/pkg"
)

type CompanyBehaviorMapper struct {
}

func (m CompanyBehaviorMapper) ConvertToEntity(behaviorInput dto.BehaviorInput) entities.CompanyBehavior {
	var companyBehavior entities.CompanyBehavior
	err1 := copier.Copy(&companyBehavior, &behaviorInput)
	util.Must(err1)
	return companyBehavior
}
