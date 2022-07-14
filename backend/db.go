package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Status int64
	Todo   string
}

func db() {
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
		db.Create(&Todo{Status: 0, Todo: "todo01_incomplete"})
		db.Create(&Todo{Status: 0, Todo: "todo02_incomplete"})
		db.Create(&Todo{Status: 1, Todo: "todo03_completed"})
		db.Create(&Todo{Status: 0, Todo: "todo04_incomplete"})
		db.Create(&Todo{Status: 0, Todo: "todo05_incomplete"})
		db.Create(&Todo{Status: 1, Todo: "todo06_completed"})
	}

	var todo Todo
	db.First(&todo)
	fmt.Println(todo)
}
