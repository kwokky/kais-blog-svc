package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kwokky/kais-blog-svc/api/v1/http"
	h "github.com/kwokky/kais-blog-svc/library/http"
)

// tagRouter 标签路由
func tagRouter(e *gin.RouterGroup) {
	group := e.Group("tag")
	{
		group.POST("", h.HandleFunc(http.TagCreate))
		group.PUT(":id", h.HandleFunc(http.TagUpdate))
		group.GET("", h.HandleFunc(http.TagList))
		group.GET(":id", h.HandleFunc(http.TagDetail))
		group.DELETE(":id", h.HandleFunc(http.TagDelete))
	}
}
