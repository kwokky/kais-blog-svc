package param

type TagCreateParams struct {
	Name string `form:"name"`
}

type TagUpdateParams struct {
	Name string `form:"name"`
	ID   int64  `form:"id"`
}

type TagDetailParams struct {
	Name string `form:"name"`
	ID   int64  `form:"id"`
}

type TagListParams struct {
	PageParams
}

type TagDeleteParams struct {
	ID int64 `form:"id"`
}
