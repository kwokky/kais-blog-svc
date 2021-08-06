package http

import (
	"github.com/kwokky/kais-blog-svc/api/v1/param"
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/library/http"
	"strconv"
)

// TagCreate 创建标签
func TagCreate(c *http.Context) {
	var params param.TagCreateParams
	if err := c.Bind(&params); err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	err := svr.TagCreate(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}

// TagUpdate 修改标签
func TagUpdate(c *http.Context) {
	var params param.TagUpdateParams
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

	err = svr.TagUpdate(params.ID, params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}

// TagList 标签列表
func TagList(c *http.Context) {
	var params param.TagListParams
	if err := c.Bind(&params); err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	handlePageSize(&params.Page, &params.Size)

	list, err := svr.TagList(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(list, nil)

}

// TagDetail 标签详情
func TagDetail(c *http.Context) {
	var params param.TagDetailParams
	var err error

	params.ID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	detail, err := svr.TagDetail(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(detail, nil)
}

// TagDelete 删除标签
func TagDelete(c *http.Context) {
	var params param.TagDeleteParams
	var err error

	params.ID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	err = svr.TagDelete(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}
