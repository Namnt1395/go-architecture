package service

import (
	"context"
	dto "go-architecture/dto/company_capacity"
	util "go-architecture/pkg"
	"go-architecture/pkg/ctxutil"
	"go-architecture/pkg/errorcode"
	"go-architecture/repository"
	service_mapper "go-architecture/service/service.mapper"
)

type CompanyCapacityService interface {
	Create(ctx context.Context, capacityInput *dto.CapacityInput) (*dto.CapacityOutput, error)
	Retrieve(ctx context.Context, id uint64) (*dto.CapacityOutput, error)
	Update(ctx context.Context, id uint64, capacityInput *dto.CapacityInput) (*dto.CapacityOutput, error)
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context) (*dto.ListCapacityOutput, error)
}

type companyCapacityService struct {
	repo   repository.CompanyCapacityRepository
	mapper service_mapper.CompanyCapacityMapper
}

func (s companyCapacityService) Create(ctx context.Context, capacityInput *dto.CapacityInput) (*dto.CapacityOutput, error) {
	if err1 := capacityInput.Validate(); err1 != nil {
		return nil, err1
	}
	capacityEntity, err1 := s.repo.Create(capacityInput)
	util.Must(err1)
	return s.mapper.ConvertToDto(*capacityEntity), nil
}

func (s companyCapacityService) Retrieve(ctx context.Context, id uint64) (*dto.CapacityOutput, error) {
	capacity, err := s.repo.Retrieve(id)
	if err != nil || capacity.ID <= 0 {
		return nil, util.BadRequestError(ctx, errorcode.ERR_CAPACITY_NOT_FOUND)
	}
	return s.mapper.ConvertToDto(*capacity), nil
}

func (s companyCapacityService) Delete(ctx context.Context, id uint64) error {
	err1 := s.repo.Delete(id)
	if err1 != nil {
		return util.BadRequestError(ctx, errorcode.ERR_REMOVE)
	}
	return nil
}

func (s companyCapacityService) Update(ctx context.Context, id uint64, capacityInput *dto.CapacityInput) (*dto.CapacityOutput, error) {
	if err := capacityInput.Validate(); err != nil {
		return nil, err
	}
	capacity, err1 := s.repo.Retrieve(id)
	if err1 != nil || capacity.ID <= 0 {
		return nil, util.BadRequestError(ctx, errorcode.ERR_CAPACITY_NOT_FOUND)
	}
	capacity, err1 = s.repo.Update(capacityInput)
	util.Must(err1)
	return s.mapper.ConvertToDto(*capacity), nil
}

func (s companyCapacityService) List(ctx context.Context) (*dto.ListCapacityOutput, error) {
	page := ctxutil.GetPageFromCtx(ctx)
	var res dto.ListCapacityOutput
	total, err1 := s.repo.Count()
	util.Must(err1)
	listCapacity, err1 := s.repo.List(page.Page, page.Size, page.Sort)
	res = *s.mapper.ConvertToListDto(*listCapacity)
	res.Build(page, total)
	return &res, nil
}
