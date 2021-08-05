package param

type CreatePostParams struct {
	Title      string `form:"title"`
	Content    string `form:"content"`
	Author     string `form:"author"`
	CategoryId int64  `form:"category_id"`
}

type UpdatePostParams struct {
	CreatePostParams
	ID int64 `form:"id"`
}

type PageParams struct {
	Page int64 `form:"page"`
	Size int64 `form:"size"`
}

type ListPostParams struct {
	PageParams
	Tag      string `form:"tag"`
	Category string `form:"category"`
}
