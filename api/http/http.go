package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kwokky/kais-blog-svc/api/post/service"
	"github.com/kwokky/kais-blog-svc/library/http"
)

var svr *service.Service

func init() {
	svr = service.New()
}

func Router(e *gin.Engine) {
	group := e.Group("v1/post")
	{
		group.POST("", http.HandleFunc(Create))
		group.PUT(":id", http.HandleFunc(Update))
		group.GET("", http.HandleFunc(List))
		group.GET(":id", http.HandleFunc(Detail))
		group.DELETE(":id", http.HandleFunc(Delete))
	}
}
