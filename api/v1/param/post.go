package param

type PostCreateParams struct {
	Title      string `form:"title"`
	Content    string `form:"content"`
	Author     string `form:"author"`
	CategoryId int64  `form:"category_id"`
}

type PostUpdateParams struct {
	PostCreateParams
	ID int64 `form:"id"`
}

type PostListParams struct {
	PageParams
	Tag      string `form:"tag"`
	Category string `form:"category"`
}

type PostDetailParams struct {
	ID int64 `form:"id"`
}

type PostDeleteParams struct {
	ID int64 `form:"id"`
}
