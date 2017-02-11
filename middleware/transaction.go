package middleware

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

const (
	TxKey = "Tx"
)

func TransactionHandler(db *dbr.Session) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			tx, _ := db.Begin()
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
