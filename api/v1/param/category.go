package param

type CategoryCreateParams struct {
	Name string `form:"name"`
}

type CategoryUpdateParams struct {
	Name string `form:"name"`
	ID   int64  `form:"id"`
}

type CategoryDetailParams struct {
	Name string `form:"name"`
	ID   int64  `form:"id"`
}

type CategoryListParams struct {
	PageParams
}

type CategoryDeleteParams struct {
	ID int64 `form:"id"`
}
