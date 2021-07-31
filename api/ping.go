package api

import "github.com/kwokky/kais-blog-svc/library/http"

// Ping 健康检查
func Ping(c *http.Context) {
	c.Response(nil, nil)
}
