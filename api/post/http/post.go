package http

import (
	"github.com/kwokky/kais-blog-svc/config"
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/library/http"
	"log"
	"strconv"
)

// Create 创建文章
func Create(c *http.Context) {
	var (
		title       = c.GetParam("title")
		content     = c.GetParam("content")
		author      = c.GetParam("author")
		categoryStr = c.GetParam("category_id")
	)

	categoryId, err := strconv.ParseInt(categoryStr, 10, 64)
	if err != nil {
		log.Printf("strconv.ParseInt(%s) error(%v)", categoryStr, err)
		c.Response(nil, ecode.RequestError)
		return
	}

	err = svr.CreatePost(title, content, author, categoryId)
	if err != nil {
		c.Response(nil, err)
	}

	c.Response(nil, nil)
}

// Update 修改文章
func Update(c *http.Context) {
	var (
		title       = c.GetParam("title")
		content     = c.GetParam("content")
		author      = c.GetParam("author")
		categoryStr = c.GetParam("category_id")
		idStr       = c.Param("id")
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("strconv.ParseInt(%s) error(%v)", idStr, err)
		c.Response(nil, ecode.RequestError)
		return
	}

	categoryId, err := strconv.ParseInt(categoryStr, 10, 64)
	if err != nil {
		log.Printf("strconv.ParseInt(%s) error(%v)", categoryStr, err)
		c.Response(nil, ecode.RequestError)
		return
	}

	err = svr.UpdatePost(id, title, content, author, categoryId)
	if err != nil {
		c.Response(nil, err)
	}

	c.Response(nil, nil)
}

// List 文章列表
func List(c *http.Context) {
	var (
		pageStr     = c.GetParam("page")
		sizeStr     = c.GetParam("size")
		tag      = c.GetParam("tag")
		category = c.GetParam("category_id")
	)

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}

	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = int64(config.Get().Default.PageSize)
	}

	list, err := svr.ListPost(page, size, tag, category)
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
