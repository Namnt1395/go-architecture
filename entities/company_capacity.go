package entities

import "gorm.io/gorm"

type CompanyCapacity struct {
	gorm.Model
	Code        string `gorm:"code;not null"`
	Name        string `gorm:"name;not null"`
	Description string `gorm:"description"`
	IsActive    bool   `gorm:"is_active;default:true"`
	CreatedBy   uint64 `gorm:"created_by"`
	UpdatedBy   uint64 `gorm:"updated_by"`
}

func (CompanyCapacity) TableName() string {
	return "company_capacity"
}
