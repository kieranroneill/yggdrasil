package main

import (
	"fmt"
	"github.com/kieranroneill/yggdrasil/apps/mimir/internal/routes"
	"github.com/kieranroneill/yggdrasil/libs/constants"
	"github.com/kieranroneill/yggdrasil/libs/types"
	"github.com/kieranroneill/yggdrasil/libs/utilities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ostafen/clover/v2"
	"log/slog"
	"os"
)

func main() {
	// open connection to the database
	db, err := clover.Open(".data")
	if err != nil {
		slog.Error(fmt.Sprintf("failed to open db: %v", err))
	}

	// defer db connect closure
	defer func(db *clover.DB) {
		if err = db.Close(); err != nil {
			slog.Error(fmt.Sprintf("error while closing the db: %v", err))
		}
	}(db)

	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())

	// routes
	// /auth
	e.GET(constants.AuthPath, routes.NewGetVersionsRoute())

	// send a request to register the app with the registry service
	if err = utilities.RegisterApp(os.Getenv("BIFROST_URL"), types.AppMetadata{
		Environment: os.Getenv("ENVIRONMENT"),
		Name:        os.Getenv("NAME"),
		URL:         fmt.Sprintf("http://%s:%s", os.Getenv("NAME"), os.Getenv("PORT")),
		Version:     os.Getenv("VERSION"),
	}); err != nil {

	}

	// start the server
	err = e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		e.Logger.Fatal(err)
	}
}
