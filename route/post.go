package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kwokky/kais-blog-svc/api/v1/http"
	h "github.com/kwokky/kais-blog-svc/library/http"
)

// postRouter 文章路由
func postRouter(e *gin.RouterGroup) {
	group := e.Group("post")
	{
		group.POST("", h.HandleFunc(http.PostCreate))
		group.PUT(":id", h.HandleFunc(http.PostUpdate))
		group.GET("", h.HandleFunc(http.PostList))
		group.GET(":id", h.HandleFunc(http.PostDetail))
		group.DELETE(":id", h.HandleFunc(http.PostDelete))
	}
}