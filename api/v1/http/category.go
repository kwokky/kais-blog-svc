package http

import (
	"github.com/kwokky/kais-blog-svc/api/v1/param"
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/library/http"
	"strconv"
)

// CreateCreate 创建分类
func CreateCreate(c *http.Context) {
	var params param.CategoryCreateParams
	if err := c.Bind(&params); err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	err := svr.CategoryCreate(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}

// CategoryUpdate 修改分类
func CategoryUpdate(c *http.Context) {
	var params param.CategoryUpdateParams
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

	err = svr.CategoryUpdate(params.ID, params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}

// CategoryList 分类列表
func CategoryList(c *http.Context) {
	var params param.CategoryListParams
	if err := c.Bind(&params); err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	handlePageSize(&params.Page, &params.Size)

	list, err := svr.CategoryList(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(list, nil)

}

// CategoryDetail 分类详情
func CategoryDetail(c *http.Context) {
	var params param.CategoryDetailParams
	var err error

	params.ID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	detail, err := svr.CategoryDetail(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(detail, nil)
}

// CategoryDelete 删除分类
func CategoryDelete(c *http.Context) {
	var params param.CategoryDeleteParams
	var err error

	params.ID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Response(nil, ecode.ServerError)
		return
	}

	err = svr.CategoryDelete(params)
	if err != nil {
		c.Response(nil, err)
		return
	}

	c.Response(nil, nil)
}
