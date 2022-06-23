package repository

import (
	"fmt"
	"go-architecture/database"
	dto "go-architecture/dto/company_capacity"
	"go-architecture/entities"
	repository "go-architecture/repository/repo.mapper"
	"gorm.io/gorm"
)

type CompanyCapacityRepository interface {
	Create(capacityInput *dto.CapacityInput) (*entities.CompanyCapacity, error)
	Update(capacityInput *dto.CapacityInput) (*entities.CompanyCapacity, error)
	Retrieve(id uint64) (companyCapacity *entities.CompanyCapacity, err error)
	Delete(id uint64) error
	List(page int, size int, sort string) (listCompanyCapacity *[]entities.CompanyCapacity, err error)
	Count() (total int64, err error)
}

type companyCapacityRepository struct {
	database.Repository
	capacityMapper repository.CompanyCapacityMapper
}

func (r companyCapacityRepository) Create(capacityInput *dto.CapacityInput) (*entities.CompanyCapacity, error) {
	var companyCapacity *entities.CompanyCapacity
	companyCapacity = r.capacityMapper.ConvertToEntity(capacityInput)
	fmt.Println("companyCapacity....", companyCapacity)
	err1 := r.DB.DB.Transaction(func(tx *gorm.DB) error {
		if err2 := tx.Create(companyCapacity).Error; err2 != nil {
			return err2
		}
		return nil
	})
	if err1 != nil {
		return nil, err1
	}
	return companyCapacity, err1
}

func (r companyCapacityRepository) Update(capacityInput *dto.CapacityInput) (*entities.CompanyCapacity, error) {
	var companyCapacity *entities.CompanyCapacity
	companyCapacity = r.capacityMapper.ConvertToEntity(capacityInput)
	err1 := r.DB.DB.Transaction(func(tx *gorm.DB) error {
		if err2 := tx.Save(companyCapacity).Error; err2 != nil {
			return err2
		}
		return nil
	})
	if err1 != nil {
		return nil, err1
	}
	return companyCapacity, err1
}

func (r companyCapacityRepository) Retrieve(id uint64) (companyCapacity *entities.CompanyCapacity, err error) {
	return companyCapacity, r.DB.DB.Where("id = ?", id).First(&companyCapacity).Error
}

func (r companyCapacityRepository) Delete(id uint64) error {
	return r.DB.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&entities.CompanyCapacity{}, id).Error
	})
}

func (r companyCapacityRepository) List(page int, size int, sort string) (listCompanyCapacity *[]entities.CompanyCapacity, err error) {
	offset := (page - 1) * size
	return listCompanyCapacity, r.DB.DB.Offset(offset).
		Limit(size).
		Order(sort).
		Model(&entities.CompanyCapacity{}).
		Find(&listCompanyCapacity).Error
}

func (r companyCapacityRepository) Count() (total int64, err error) {
	return total, r.DB.DB.Model(entities.CompanyCapacity{}).Count(&total).Error
}
