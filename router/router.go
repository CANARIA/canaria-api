package router

import (
	"net/http"
	"runtime"

	"github.com/CANARIA/canaria-api/api"
	"github.com/CANARIA/canaria-api/db"
	appMw "github.com/CANARIA/canaria-api/middleware"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func init() {
	println("最初に実行される")
	// db, err := gorm.Open("mysql", "root:password@tcp(mysql:3306)/canaria?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	println("DB connection failed")
	// }
}

func Init() *echo.Echo {
	e := echo.New()

	runtime.GOMAXPROCS(runtime.NumCPU())

	e.Logger.Debug()

	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	// e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// set custome middleware
	e.Use(appMw.TransactionHandler(db.Init()))

	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		println("Redis connection failed")
	}

	defer c.Close()

	// testKey, err := redis.String(c.Do("GET", "testkey"))
	// if err != nil {
	// 	println("Can not get value")
	// }

	// view
	// e.GET("/", handler.Home)
	// e.GET("/snippet", handler.Snippet)
	// e.GET("/create", handler.SnippetCreate)
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!")
	})

	// api
	v1 := e.Group("/api/v1")
	{
		v1.POST("/auth/preregister", api.PreRegister())
		v1.POST("/auth/register", api.AuthRegister())
	}

	return e
}
