package routes

import (
	"fmt"
	apitypes "github.com/kieranroneill/yggdrasil/internal/types/api"
	commontypes "github.com/kieranroneill/yggdrasil/internal/types/common"
	configutilities "github.com/kieranroneill/yggdrasil/internal/utilities/config"
	"github.com/labstack/echo/v4"
	"github.com/ostafen/clover/v2"
	"log/slog"
	"net/http"
)

func NewGetVersionsRoute(config commontypes.Config, db *clover.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		configPath, err := configutilities.DefaultConfigPath()
		if err != nil {
			slog.Error(fmt.Sprintf("failed to get config path: %v", err))
		}

		return c.JSON(http.StatusOK, apitypes.VersionsResponse{
			ConfigPath: configPath,
			Version:    config.Version,
		})
	}
}
