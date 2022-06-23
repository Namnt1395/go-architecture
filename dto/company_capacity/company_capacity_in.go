package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CapacityInput struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreatedBy   uint64 `json:"-"`
	UpdatedBy   uint64 `json:"-"`
}

func (i CapacityInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Code, validation.Required, validation.Length(3, 20)),
		validation.Field(&i.Name, validation.Required, validation.Length(1, 250)),
		validation.Field(&i.Description, validation.Length(0, 1000)))
}
