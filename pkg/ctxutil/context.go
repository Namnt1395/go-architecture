package ctxutil

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-architecture/dto"
	"go-architecture/pkg/constant"
	"strconv"
)

func GetAcceptLanguageFromCtx(ctx context.Context) (lang string) {
	switch c := ctx.(type) {
	case *gin.Context:
		lang = c.GetHeader(constant.HeaderAcceptLanguage)
	default:
		lang = constant.DefaultLang
	}
	return lang
}

func GetPageFromCtx(ctx context.Context) (page dto.Page) {
	switch c := ctx.(type) {
	case *gin.Context:
		page.Page, _ = strconv.Atoi(c.Query("page"))
		page.Size, _ = strconv.Atoi(c.Query("size"))
		page.Sort = c.Query("sort")
	default:
		page.Page, _ = c.Value("page").(int)
		page.Size, _ = c.Value("size").(int)
		page.Sort, _ = c.Value("size").(string)
	}
	if page.Page <= 0 {
		page.Page = constant.DefaultPage
	}
	if page.Size <= 0 {
		page.Size = constant.DefaultPageSize
	}
	if page.Sort == "" {
		page.Sort = constant.DefaultPageSort
	}
	return page
}
