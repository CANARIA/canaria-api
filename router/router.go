package router

import

// "github.com/SnippetsBucket/snicket/db"
// "github.com/SnippetsBucket/snicket/handler"
// sckMw "github.com/SnippetsBucket/snicket/middleware"

(
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	runtime.GOMAXPROCS(runtime.NumCPU())

	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	// e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// set custome middleware
	// e.Use(sckMw.TransactionHandler(db.Init()))

	// view
	// e.GET("/", handler.Home)
	// e.GET("/snippet", handler.Snippet)
	// e.GET("/create", handler.SnippetCreate)

	// api
	// v1 := e.Group("/api/v1")
	// {
	// 	v1.POST("/snippet/create", api.Create())
	// 	v1.POST("/preview", api.Preview())
	// }

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!")
	})
	return e
}
