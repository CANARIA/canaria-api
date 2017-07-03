package middleware

import (
	"github.com/labstack/echo"
	"github.com/CANARIA/canaria-api/core/config"
	"google.golang.org/appengine"
)

func AppEngineContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			ctx := appengine.NewContext(c.Request())

			c.Set(config.AppEngineContextName, &ctx)
			//c.Echo().DefaultHTTPErrorHandler(err, c)

			return next(c)

		})
	}
}
