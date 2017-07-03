package api

import (
	"net/http"


	"github.com/CANARIA/canaria-api/core/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"time"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
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
		ctx := appengine.NewContext(c.Request())
		log.Debugf(ctx, "tag list")

		q := datastore.NewQuery("tags")
		var tmp []model.Tag
		keys, err := q.GetAll(ctx, &tmp)
		if err != nil {
			log.Errorf(ctx, "could not fetch tags")
		}

		//tx := c.Get("Tx").(*gorm.DB)
		//tags, err := service.FindTags(tx)
		//
		//if err != nil {
		//	return fmt.Errorf("failed get tags by condition. err={%s}", err.Error())
		//}

		return c.JSON(http.StatusOK, keys)
	}
}

func Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := model.Tag{
			TagId: 1,
			TagName: "テスト",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ctx := appengine.NewContext(c.Request())
		key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "tags", nil), &t)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, key)
		}
		return c.JSON(http.StatusOK, key)
	}
}
