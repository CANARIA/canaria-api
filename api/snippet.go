package api

// import (
// 	"github.com/Sirupsen/logrus"
// 	"github.com/SnippetsBucket/snicket/model"
// 	"github.com/gocraft/dbr"
// 	"github.com/labstack/echo"
// 	"github.com/valyala/fasthttp"
// )

// func Create() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		s := new(model.SnippetJson)
// 		if err := c.Bind(s); err != nil {
// 			return err
// 		}

// 		tx := c.Get("Tx").(*dbr.Tx)
// 		snippet := model.CreateSnippet(s.SnippetTitle, s.SnippetText)

// 		if err := snippet.Save(tx); err != nil {
// 			logrus.Debug(err)
// 			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
// 		}

// 		return c.JSON(fasthttp.StatusCreated, snippet)
// 	}
// }
