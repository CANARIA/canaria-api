package api

// import (
// 	"github.com/SnippetsBucket/snicket/model"

// 	"github.com/labstack/echo"
// 	"github.com/valyala/fasthttp"
// )

// func Preview() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		rqJson := new(model.PreviewJson)
// 		if err := c.Bind(rqJson); err != nil {
// 			return err
// 		}

// 		parsed := model.ParseMarkdown(rqJson)
// 		built := model.BuildPreviewJson(parsed)
// 		return c.JSON(fasthttp.StatusCreated, built)
// 	}
// }
