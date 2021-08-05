package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kwokky/kais-blog-svc/api"
	post "github.com/kwokky/kais-blog-svc/api/v1/http"
	"github.com/kwokky/kais-blog-svc/library/http"
)

func Routes(app *gin.Engine) {
	app.GET("ping", http.HandleFunc(api.Ping))
	post.Router(app)
}
