package handler

import (
	"net/http"

	"github.com/go_nuxt_3/server/api/api004/domain"
	"github.com/go_nuxt_3/server/environment/db"
	"github.com/labstack/echo"
)

// DeleteTodo deletes todo
func DeleteTodo(c echo.Context) error {
	param := domain.RequestParam{}

	if err := c.Bind(&param); err != nil {
		return err
	}

	if err := deleteTodo(param); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, "Deleted")
}

func deleteTodo(param domain.RequestParam) error {
	db := db.CreateDBConnection()

	data := domain.Todo{}
	err := db.Table("todos").
		Where("id = ?", param.ID).
		Find(&data).
		Delete(data).Error

	if err != nil {
		return err
	}

	return nil
}
