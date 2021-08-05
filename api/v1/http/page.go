package http

import "github.com/kwokky/kais-blog-svc/config"

// handlePageSize 处理分页和条数
func handlePageSize(page *int64, size *int64) {
	if *page == 0 {
		*page = 1
	}
	if *size == 0 {
		*size = config.Get().PageSize
	}
}
