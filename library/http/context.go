package http

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

// HandleFunc 扩展 gin.Context
func HandleFunc(fn func(*Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(&Context{ctx})
	}
}
