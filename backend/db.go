package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string
	Todo string
}

func main() {
	dsn := "host=localhost user=postgres password=passwd dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Todo{})
	fmt.Println("migrated")

	var count int64
	db.Model(&Todo{}).Count(&count)
	if count == 0 {
		db.Create(&Todo{Name: "user01", Todo: "todo01"})
		db.Create(&Todo{Name: "user02", Todo: "todo02"})
		db.Create(&Todo{Name: "user03", Todo: "todo03"})
	}

	var todo Todo
	db.First(&todo)
	fmt.Println(todo)
}
