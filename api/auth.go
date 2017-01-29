package api

import (
	"net/http"

	"github.com/CANARIA/canaria-api/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var db gorm.DB

func AuthRegister() echo.HandlerFunc {
	return func(c echo.Context) error {

		authJson := new(model.AuthRegister)
		if err := c.Bind(authJson); err != nil {
			return err
		}

		model.AccountCreate(authJson, &db)

		return c.JSON(http.StatusOK, authJson)
	}
}
