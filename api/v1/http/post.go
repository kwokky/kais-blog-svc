package http

import (
	"github.com/kwokky/kais-blog-svc/api/v1/param"
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/library/http"
	"strconv"
)

// PostCreate 创建文章
func PostCreate(c *http.Context) {
	var params param.PostCreateParams
	if err := c.Bind(&params); err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	err := svr.PostCreate(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}

// PostUpdate 修改文章
func PostUpdate(c *http.Context) {
	var params param.PostUpdateParams
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

	err = svr.PostUpdate(params.ID, params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}

// PostList 文章列表
func PostList(c *http.Context) {
	var params param.PostListParams
	if err := c.Bind(&params); err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	handlePageSize(&params.Page, &params.Size)

	list, err := svr.PostList(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(list, nil)

}

// PostDetail 文章详情
func PostDetail(c *http.Context) {
	var params param.PostDetailParams
	var err error

	params.ID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	detail, err := svr.PostDetail(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(detail, nil)
}

// PostDelete 删除文章
func PostDelete(c *http.Context) {
	var params param.PostDeleteParams
	var err error

	params.ID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	err = svr.PostDelete(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}
