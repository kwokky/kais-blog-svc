package http

import (
	"github.com/kwokky/kais-blog-svc/api/post/param"
	"github.com/kwokky/kais-blog-svc/config"
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/library/http"
	"strconv"
)

// Create 创建文章
func Create(c *http.Context) {
	var params param.CreatePostParams
	if err := c.Bind(&params); err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	err := svr.CreatePost(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}

// Update 修改文章
func Update(c *http.Context) {
	var params param.UpdatePostParams
	if err := c.Bind(&params); err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	var err error
	params.ID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	err = svr.UpdatePost(params.ID, params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}

// List 文章列表
func List(c *http.Context) {
	var params param.ListPostParams
	if err := c.Bind(&params); err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	if params.Size == 0 {
		params.Size = config.Get().PageSize
	}

	if params.Page == 0 {
		params.Page = 1
	}

	list, err := svr.ListPost(params)
	if err != nil {
		c.Response(nil, err)
	}

	c.Response(list, nil)

}

// Detail 文章详情
func Detail(c *http.Context) {

}

// Delete 删除文章
func Delete(c *http.Context) {

}
