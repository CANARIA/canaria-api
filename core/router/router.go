package router

import (
	"net/http"

	"github.com/CANARIA/canaria-api/core/api"
	"github.com/CANARIA/canaria-api/core/db"
	appMw "github.com/CANARIA/canaria-api/core/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/CANARIA/canaria-api/core/env"
	"google.golang.org/appengine/log"
	"github.com/CANARIA/canaria-api/core/config"
	"golang.org/x/net/context"
	"os"
	"fmt"
)


func Init() *echo.Echo {
	app := echo.New()

	env.SetUp()

	app.Debug = true
	app.Logger.Debug()

	app.Use(appMw.AppEngineContext())
	app.Use(mw.Logger())
	app.Use(mw.Recover())
	app.Use(mw.Gzip())
	app.Use(mw.CORS())
	app.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:  []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAccessControlExposeHeaders, echo.HeaderAuthorization},
		ExposeHeaders: []string{"access_token", "Authorization"},
	}))
	// e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// set custome middleware
	app.Use(appMw.TransactionHandler(db.Init()))

	// c, err := redis.Dial("tcp", "redis:6379")
	// if err != nil {
	// 	println("Redis connection failed")
	// }

	// defer c.Close()

	// testKey, err := redis.String(c.Do("GET", "testkey"))
	// if err != nil {
	// 	println("Can not get value")
	// }

	app.GET("/", func(c echo.Context) error {
		ctx := c.Get(config.AppEngineContextName).(*context.Context)
		log.Debugf(*ctx, "from echo.Context")
		apiEnv := os.Getenv("API_ENV")
		return c.JSON(http.StatusOK, fmt.Sprintf("hello! %s", apiEnv))
	})
	app.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	// api
	v1 := app.Group("/api/v1")
	{
		v1.POST("/auth/checktoken", api.CheckToken())
		v1.POST("/auth/preregister", api.PreRegister())
		v1.POST("/auth/register", api.AuthRegister())
		v1.POST("/auth/login", api.Login())
		v1.POST("/auth/check", api.CheckAuth(), appMw.AuthFilter())
		v1.GET("/populartags", api.PopularTags())
		v1.GET("/tags", api.Tags())
		v1.GET("/tags/create", api.Create())
	}

	return app
}
