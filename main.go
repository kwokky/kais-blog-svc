package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kwokky/kais-blog-svc/config"
	"github.com/kwokky/kais-blog-svc/model"
	"github.com/kwokky/kais-blog-svc/route"
)

func main() {
	r := gin.Default()

	// 初始化数据库
	model.DB()
	// 路由
	route.Routes(r)

	_ = r.Run(fmt.Sprintf(":%d", config.Get().Port))
}
