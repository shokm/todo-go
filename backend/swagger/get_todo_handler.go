package swagger

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/shokm/todo-go/backend/swagger/gen/models"
	"github.com/shokm/todo-go/backend/swagger/gen/restapi/factory/user_api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetContentDB(pid int) (int, string) {
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

	// セッションを切る
	sqlDB, err := db.DB()
	sqlDB.Close()

	return todo.Status, todo.Todo
}

func GetTodo(p user_api.GetTodoByTodoIDParams) middleware.Responder {
	pid := p.TodoID
	n, t := GetContentDB(int(pid))
	payload := &models.User{
		// id
		ID: int64(pid),
		// name
		Status: int64(n),
		// todo
		Todo: t,
	}
	return user_api.NewGetTodoByTodoIDOK().WithPayload(payload)
}
