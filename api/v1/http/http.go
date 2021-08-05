package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kwokky/kais-blog-svc/api/v1/service"
	"github.com/kwokky/kais-blog-svc/library/http"
)

var svr *service.Service

func init() {
	svr = service.New()
}

func Router(e *gin.Engine) {
	group := e.Group("v1")
	{
		postRouter(group)
		categoryRouter(group)
	}
}

// postRouter 文章路由
func postRouter(e *gin.RouterGroup) {
	post := e.Group("post")
	{
		post.POST("", http.HandleFunc(PostCreate))
		post.PUT(":id", http.HandleFunc(PostUpdate))
		post.GET("", http.HandleFunc(PostList))
		post.GET(":id", http.HandleFunc(PostDetail))
		post.DELETE(":id", http.HandleFunc(PostDelete))
	}
}

// categoryRouter 分类路由
func categoryRouter(e *gin.RouterGroup) {
	post := e.Group("category")
	{
		post.POST("", http.HandleFunc(CreateCreate))
		post.PUT(":id", http.HandleFunc(CategoryUpdate))
		post.GET("", http.HandleFunc(CategoryList))
		post.GET(":id", http.HandleFunc(CategoryDetail))
		post.DELETE(":id", http.HandleFunc(CategoryDelete))
	}
}
