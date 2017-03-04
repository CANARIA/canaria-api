package api

import (
	"net/http"

	"github.com/CANARIA/canaria-api/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

/*
人気のタグ一覧
*/
func PopularTags() echo.HandlerFunc {
	return func(c echo.Context) error {

		tx := c.Get("Tx").(*gorm.DB)
		popularTagDao := model.PopularTagDaoFactory(tx)
		popTags, err := popularTagDao.FindAll()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed get all popular tags")
		}

		return c.JSON(http.StatusOK, popTags)
	}

}
