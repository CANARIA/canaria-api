package api

import (
	"net/http"

	"github.com/CANARIA/canaria-api/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

func AuthRegister() echo.HandlerFunc {
	return func(c echo.Context) error {

		authJson := new(model.AuthRegister)
		if err := c.Bind(authJson); err != nil {
			return err
		}

		tx := c.Get("Tx").(*dbr.Tx)

		account := model.AccountImpl(authJson)
		if err := account.AccountCreate(tx); err != nil {
			println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, authJson)
	}
}
