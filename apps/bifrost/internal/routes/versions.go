package routes

import (
	"github.com/kieranroneill/yggdrasil/libs/types"
	"github.com/labstack/echo/v4"
	"github.com/ostafen/clover/v2"
	"net/http"
	"os"
)

func NewGetVersionsRoute(db *clover.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, types.VersionsResponse{
			Environment: os.Getenv("ENVIRONMENT"),
			Name:        os.Getenv("NAME"),
			Version:     os.Getenv("VERSION"),
		})
	}
}
