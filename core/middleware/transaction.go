package middleware

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"github.com/mjibson/goon"
	"github.com/CANARIA/canaria-api/core/config"
)

const (
	TxKey = "Tx"
)

func TransactionHandler(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			ctx := appengine.NewContext(c.Request())
			g := goon.FromContext(ctx)
			log.Infof(ctx, "Generate goon instance")

			c.Set(config.Goon, g)


			if err := next(c); err != nil {
				//tx.Rollback()
				log.Errorf(ctx, "End failed Datastore: %s", err)
				//fmt.Println("Transction Rollback")
				//logrus.Debug("Transction Rollback: ", err)
				return err
			}

			log.Infof(ctx, "End success Datastore")
			//fmt.Println("Transction Commit")
			//logrus.Debug("Transction Commit")
			//tx.Commit()

			return nil
		})
	}
}
