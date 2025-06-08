package routes

import (
	"github.com/labstack/echo/v4"
	"io/fs"
	"net/http"
)

func NewGetAppRoute(staticFiles fs.FS) echo.HandlerFunc {
	return func(c echo.Context) error {
		indexFile, err := staticFiles.Open("index.html")
		if err != nil {
			return c.NoContent(404)
		}

		defer indexFile.Close()

		return c.Stream(http.StatusOK, "text/html", indexFile)
	}
}
