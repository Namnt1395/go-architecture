package dto

import "go-architecture/dto"

type CapacityData struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreatedBy   uint64 `json:"created_by"`
	UpdatedBy   uint64 `json:"updated_by"`
}

type CapacityOutput struct {
	Data CapacityData `json:"data"`
}

type ListCapacityOutput struct {
	Page       int            `json:"page"`
	Size       int            `json:"size"`
	Sort       string         `json:"sort"`
	Total      int64          `json:"total"`
	TotalPages int64          `json:"total_pages"`
	Content    []CapacityData `json:"content"`
}

func (d *ListCapacityOutput) Build(page dto.Page, total int64) {
	d.Total = total
	d.Page = page.Page
	d.Size = page.Size
	d.Sort = page.Sort
	d.TotalPages = total / int64(d.Size)
	if d.TotalPages%int64(d.Size) != 0 {
		d.TotalPages += 1
	}
}
