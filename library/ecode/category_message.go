package ecode

var CategoryMessage = map[Code]string{
	CategoryNameEmpty:   "名称不能为空",
	CategoryCreateError: "分类创建失败",
	CategoryUpdateError: "分类更新失败",
	CategoryDeleteError: "分类删除失败",
	CategoryNotFound:    "分类不存在",
	CategoryParamError:  "分类参数错误",
	CategoryExist:       "分类名称已存在",
}
