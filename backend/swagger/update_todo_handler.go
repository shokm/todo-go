package swagger

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func UpdContentDB(pid int) (int, string) {
	dsn := "host=localhost user=postgres password=passwd dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	type Todo struct {
		gorm.Model
		Status int
		Todo   string
	}
	var todo Todo
	todo.ID = uint(pid)
	db.First(&todo)
	return todo.Status, todo.Todo
}
