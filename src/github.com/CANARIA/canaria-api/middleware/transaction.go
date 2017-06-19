package middleware

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

const (
	TxKey = "Tx"
)

func TransactionHandler(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			tx := db.Begin()
			logrus.Debug("Transaction Start")
			fmt.Println("Transaction Start")

			c.Set(TxKey, tx)

			if err := next(c); err != nil {
				tx.Rollback()
				fmt.Println("Transction Rollback")
				logrus.Debug("Transction Rollback: ", err)
				return err
			}

			fmt.Println("Transction Commit")
			logrus.Debug("Transction Commit")
			tx.Commit()

			return nil
		})
	}
}
