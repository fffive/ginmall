package dao

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate()
	if err != nil {
		panic(err)
	}
	
	return 
}
