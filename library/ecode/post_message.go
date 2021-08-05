package ecode

var PostMessage = map[Code]string{
	PostTitleEmpty:   "标题不能为空",
	PostContentEmpty: "内容不能为空",
	PostAuthorEmpty:  "作者不能为空",
	PostCreateError:  "文章创建失败",
	PostUpdateError:  "文章更新失败",
	PostDeleteError:  "文章删除失败",
	PostNotFound:     "文章不存在",
	PostParamError:   "文章参数错误",
}
