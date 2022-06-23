package handle

import (
	"github.com/gin-gonic/gin"
	dto "go-architecture/dto/company_capacity"
	util "go-architecture/pkg"
	"go-architecture/service"
	"net/http"
	"strconv"
)

type CompanyCapacityHandler interface {
	Setup(group *gin.RouterGroup)
}

type companyCapacityHandler struct {
	service service.CompanyCapacityService
}

func (r companyCapacityHandler) Setup(group *gin.RouterGroup) {
	group.GET("", r.List)
	group.GET(":id", r.Retrieve)
	group.POST("", r.Create)
	group.PUT(":id", r.Update)
	group.DELETE(":id", r.Delete)
}

// Create
// @Summary Create new capacity
// @Description Create new capacity
// @Tags Capacity
// @Security ApiKeyAuth
// @Accept json
// @Param body body dto.CapacityInput true "JSON body"
// @Success 200 {object} dto.CapacityOutput
// @Failure 500 {object} interface{} "{"error_code": "<Mã lỗi>", "error_msg": "<Nội dung lỗi>"}"
// @Router /api/v1/capacities [post]
func (r companyCapacityHandler) Create(c *gin.Context) {
	var capacityInput dto.CapacityInput
	err := c.BindJSON(&capacityInput)
	util.Must(err)
	res, err := r.service.Create(c, &capacityInput)
	util.CheckError(err)
	c.JSON(http.StatusOK, res)
}

// Update
// @Summary Update capacity
// @Description Update capacity
// @Tags Capacity
// @Security ApiKeyAuth
// @Accept json
// @Param id path int true "id"
// @Param body body dto.CapacityInput true "JSON body"
// @Success 200 {object} dto.CapacityOutput
// @Failure 500 {object} interface{} "{"error_code": "<Mã lỗi>", "error_msg": "<Nội dung lỗi>"}"
// @Router /api/v1/capacities/{id} [put]
func (r companyCapacityHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	util.Must(err)
	var capacityInput dto.CapacityInput
	util.Must(c.BindJSON(&capacityInput))
	res, err1 := r.service.Update(c, uint64(id), &capacityInput)
	util.Must(err1)
	c.JSON(http.StatusOK, res)
}

// Delete
// @Summery Delete Capacity
// @Description  Delete Capacity
// @Tags Capacity
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} interface{}
// @Failure 500 {object} interface{} "{"error_code": "<Mã lỗi>", "error_msg": "<Nội dung lỗi>"}"
// @Router /api/v1/capacities/{id} [delete]
func (r companyCapacityHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	util.Must(err)
	err1 := r.service.Delete(c, uint64(id))
	util.Must(err1)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// Retrieve
// @Summary Get info capacity
// @DescriptionGet info capacity
// @Tags Capacity
// @Security ApiKeyAuth
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} dto.CapacityOutput
// @Failure 500 {object} interface{} "{"error_code": "<Mã lỗi>", "error_msg": "<Nội dung lỗi>"}"
// @Router /api/v1/capacities/{id} [get]
func (r companyCapacityHandler) Retrieve(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	util.Must(err)
	res, err1 := r.service.Retrieve(c, uint64(id))
	util.Must(err1)
	c.JSON(http.StatusOK, res)
}

// List
// @Summary Get list capacity
// @DescriptionGet list capacity
// @Tags Capacity
// @Security ApiKeyAuth
// @Param page query int 	false "page" default(1)
// @Param size query int 	false "size" default(10)
// @Param sort query string false "sort" default(created_at desc)
// @Param search query string false "search"
// @Param is_active query string false "is_active"
// @Param capacity_group_ids[] query string false "capacity_group_ids[]"
// @Success 200 {object} dto.ListCapacityOutput
// @Failure 500 {object} interface{} "{"error_code": "<Mã lỗi>", "error_msg": "<Nội dung lỗi>"}"
// @Router /api/v1/capacities [get]
func (r companyCapacityHandler) List(c *gin.Context) {
	output, err := r.service.List(c)
	util.CheckError(err)
	c.JSON(200, output)
}
