package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BehaviorInput struct {
	CapacityId  uint64 `json:"capacity_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreatedBy   uint64 `json:"-"`
	UpdatedBy   uint64 `json:"-"`
}

func (i BehaviorInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Name, validation.Required),
		validation.Field(&i.Code, validation.Required))
}
