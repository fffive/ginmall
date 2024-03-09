package serializer

import "ginmall/model"

type CategoryVo struct {
	Id           uint   `json:"id" form:"id"`
	CategoryName string `json:"category_name" form:"category_id"`
	CreateAt     int64  `json:"create_at" form:"create_at"`
}

func BuildCategory(item *model.Category) CategoryVo {
	return CategoryVo{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreateAt: item.CreatedAt.Unix(),
	}
}

func BuildCategories(items []*model.Category) (categories []CategoryVo) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}

	return
}
