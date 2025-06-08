package main

import (
	"embed"
	"fmt"
	"github.com/kieranroneill/yggdrasil/internal/constants"
	"github.com/kieranroneill/yggdrasil/internal/routes"
	commontypes "github.com/kieranroneill/yggdrasil/internal/types/common"
	osutilities "github.com/kieranroneill/yggdrasil/internal/utilities/os"
	storageutilities "github.com/kieranroneill/yggdrasil/internal/utilities/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
)

//go:embed web/dist/*
var embeddedStaticAssets embed.FS
var Version string

func main() {
	var err error

	debug := osutilities.GetEnvWithDefault("DEBUG", "false") == "true"
	level := slog.LevelError

	// if debug is enabled, set the log level to debug
	if debug {
		level = slog.LevelDebug
	}

	slog.SetLogLoggerLevel(level)

	config := commontypes.Config{
		Debug:   debug,
		Version: Version,
	}

	// setup/open database
	db, err := storageutilities.SetupDatabase()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to open db: %v", err))

		os.Exit(1)
	}

	e := echo.New()

	e.HideBanner = true

	staticFiles, err := fs.Sub(embeddedStaticAssets, "web/dist")
	if err != nil {
		slog.Error(fmt.Sprintf("failed to load static assets: %v", err))

		os.Exit(1)
	}

	// middlewares
	e.Use(middleware.Logger())

	// routes
	// /app/*
	e.GET(constants.AppPath, routes.NewGetAppRoute(staticFiles))
	e.GET(fmt.Sprintf("%s/*", constants.AppPath), echo.WrapHandler(http.StripPrefix(fmt.Sprintf("%s/", constants.AppPath), http.FileServer(http.FS(staticFiles)))))
	// /versions
	e.GET(constants.VersionsPath, routes.NewGetVersionsRoute(config, db))

	// display a welcome message
	fmt.Println(fmt.Sprintf(constants.WelcomeMessage, config.Version))

	// start the server
	err = e.Start(fmt.Sprintf(":%s", osutilities.GetEnvWithDefault("PORT", "3000")))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to start server: %v", err))

		os.Exit(1)
	}
}
