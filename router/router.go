package router

import (
	"net/http"
	"runtime"

	"github.com/CANARIA/canaria-api/api"
	"github.com/CANARIA/canaria-api/db"
	appMw "github.com/CANARIA/canaria-api/middleware"
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

	e.Debug = true
	e.Logger.Debug()

	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())
	e.Use(mw.CORS())
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:  []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAccessControlExposeHeaders, echo.HeaderAuthorization},
		ExposeHeaders: []string{"access_token", "Authorization"},
	}))
	// e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// set custome middleware
	e.Use(appMw.TransactionHandler(db.Init()))

	// c, err := redis.Dial("tcp", "redis:6379")
	// if err != nil {
	// 	println("Redis connection failed")
	// }

	// defer c.Close()

	// testKey, err := redis.String(c.Do("GET", "testkey"))
	// if err != nil {
	// 	println("Can not get value")
	// }

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!")
	})

	// api
	v1 := e.Group("/api/v1")
	{
		v1.POST("/auth/checktoken", api.CheckToken())
		v1.POST("/auth/preregister", api.PreRegister())
		v1.POST("/auth/register", api.AuthRegister())
		v1.POST("/auth/login", api.Login())
		v1.POST("/auth/check", api.CheckAuth(), appMw.AuthFilter())
		v1.GET("/populartags", api.PopularTags())
	}

	return e
}
