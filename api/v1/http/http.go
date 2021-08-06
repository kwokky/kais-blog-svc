package http

import (
	"github.com/kwokky/kais-blog-svc/api/v1/service"
)

var svr *service.Service

func init() {
	svr = service.New()
}