package main

import (
  "fmt"
  _constants "github.com/kieranroneill/yggdrasil/apps/bifrost/internal/constants"
  "github.com/kieranroneill/yggdrasil/apps/bifrost/internal/routes"
  "github.com/kieranroneill/yggdrasil/libs/constants"
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

    os.Exit(1)
  }

  // create collections
  if err = utilities.CreateCollectionIfNoneExists(db, _constants.AppsCollection); err != nil {
    slog.Error(fmt.Sprintf("failed to create %s collection: %v", _constants.AppsCollection, err))

    os.Exit(1)
  }

  // create indexes
  if err = utilities.CreateIndexIfNoneExists(db, _constants.AppsCollection, "name"); err != nil {
    slog.Error(fmt.Sprintf("failed to create %s index in collection %s: %v", "name", _constants.AppsCollection, err))

    os.Exit(1)
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
  // /versions
  e.GET(constants.VersionsPath, routes.NewGetVersionsRoute(db))
  // /apps
  // /apps/register
  e.POST(fmt.Sprintf("%s/%s", constants.AppsPath, constants.RegisterPath), routes.NewPostAppsRegisterRoute(db))

  // start the server
  err = e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
  if err != nil {
    slog.Error(fmt.Sprintf("failed to start app %s: %v", os.Getenv("NAME"), err))

    os.Exit(1)
  }
}
