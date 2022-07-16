package swagger

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/shokm/todo-go/backend/swagger/gen/models"
	"github.com/shokm/todo-go/backend/swagger/gen/restapi/factory/user_api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpdContentDB(pid int, pBody *models.UpdateUserReq) error {
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
	todo.Status = int(*pBody.Status)
	todo.Todo = pBody.Todo

	// すでに値が存在した場合更新
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&todo)

	// セッションを切る
	sqlDB, err := db.DB()
	sqlDB.Close()

	return nil
}

func UpdTodo(p user_api.UpdateTodoByTodoIDParams) middleware.Responder {
	pid := p.TodoID
	pBody := p.Body
	UpdContentDB(int(pid), pBody)

	return user_api.NewUpdateTodoByTodoIDOK()
}
