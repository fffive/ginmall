package dao

import "ginmall/model"

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Address{},
			&model.Cart{},
			&model.Carousel{},
			&model.Category{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.User{},
			&model.Product{},
			&model.ProductImg{},
			&model.Admin{},
		)
	if err != nil {
		panic(err)
	}
}
