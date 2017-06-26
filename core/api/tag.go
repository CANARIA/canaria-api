package api

import (
	"net/http"

	"fmt"

	"github.com/CANARIA/canaria-api/core/model"
	"github.com/CANARIA/canaria-api/core/service"
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

/*
タグ一覧
*/
func Tags() echo.HandlerFunc {
	return func(c echo.Context) error {

		tx := c.Get("Tx").(*gorm.DB)
		tags, err := service.FindTags(tx)

		if err != nil {
			return fmt.Errorf("failed get tags by condition. err={%s}", err.Error())
		}

		return c.JSON(http.StatusOK, tags)
	}
}
