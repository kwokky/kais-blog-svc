package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kwokky/kais-blog-svc/api/v1/http"
	h "github.com/kwokky/kais-blog-svc/library/http"
)

// categoryRouter 分类路由
func categoryRouter(e *gin.RouterGroup) {
	group := e.Group("category")
	{
		group.POST("", h.HandleFunc(http.CreateCreate))
		group.PUT(":id", h.HandleFunc(http.CategoryUpdate))
		group.GET("", h.HandleFunc(http.CategoryList))
		group.GET(":id", h.HandleFunc(http.CategoryDetail))
		group.DELETE(":id", h.HandleFunc(http.CategoryDelete))
	}
}
