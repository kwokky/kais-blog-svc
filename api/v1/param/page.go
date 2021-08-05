package param

type PageParams struct {
	Page int64 `form:"page"`
	Size int64 `form:"size"`
}
