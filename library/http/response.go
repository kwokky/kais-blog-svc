package http

import (
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/library/render"
	"net/http"
)

// Response
func (c *Context) Response(data interface{}, err error) {
	bcode := ecode.Cause(err)
	c.JSON(http.StatusOK, render.JSON{
		Code:    bcode.Code(),
		Message: bcode.Message(),
		Data:    data,
	})
}
