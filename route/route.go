package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kwokky/kais-blog-svc/api"
	"github.com/kwokky/kais-blog-svc/library/http"
)

func Routes(app *gin.Engine) {
	app.GET("ping", http.HandleFunc(api.Ping))
	group := app.Group("v1")
	{
		postRouter(group)
		categoryRouter(group)
		tagRouter(group)
	}
}
