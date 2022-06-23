package dto

import "go-architecture/dto"

type BehaviorData struct {
	Id          uint64 `json:"id"`
	CapacityId  uint64 `json:"capacity_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type BehaviorOutput struct {
	Data BehaviorData `json:"data"`
}

type ListBehaviorOutput struct {
	Page       int            `json:"page"`
	Size       int            `json:"size"`
	Sort       string         `json:"sort"`
	Total      int64          `json:"total"`
	TotalPages int64          `json:"total_pages"`
	Content    []BehaviorData `json:"content"`
}

func (d *ListBehaviorOutput) Build(page dto.Page, total int64) {
	d.Total = total
	d.Page = page.Page
	d.Size = page.Size
	d.Sort = page.Sort
	d.TotalPages = total / int64(d.Size)
	if d.TotalPages%int64(d.Size) != 0 {
		d.TotalPages += 1
	}
}
