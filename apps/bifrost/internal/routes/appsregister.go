package routes

import (
  "fmt"
  "github.com/fatih/structs"
  _constants "github.com/kieranroneill/yggdrasil/apps/bifrost/internal/constants"
  "github.com/kieranroneill/yggdrasil/libs/types"
  "github.com/labstack/echo/v4"
  "github.com/ostafen/clover/v2"
  "github.com/ostafen/clover/v2/document"
  "github.com/ostafen/clover/v2/query"
  "log/slog"
  "net/http"
)

func NewPostAppsRegisterRoute(db *clover.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    var metadata types.AppMetadata

    if err := c.Bind(metadata); err != nil {
      slog.Error(fmt.Sprintf("failed to parse request body: %v", err))

      return c.String(http.StatusInternalServerError, err.Error())
    }

    _document, err := db.FindFirst(query.NewQuery(_constants.AppsCollection).Where(query.Field("name").Eq(metadata.Name)))
    if err != nil {
      slog.Error(fmt.Sprintf("failed to get metadata for app %s: %v", metadata.Name, err))

      return c.String(http.StatusInternalServerError, err.Error())
    }

    // if no app exists in the registry insert a new one
    if _document == nil {
      _document = document.NewDocumentOf(metadata)
      if err = db.Insert(_constants.AppsCollection, _document); err != nil {
        slog.Error(fmt.Sprintf("failed to insert app %s: %v", metadata.Name, err))

        return c.String(http.StatusInternalServerError, err.Error())
      }

      return c.NoContent(http.StatusCreated)
    }

    // if it exists, simply update the document
    if err = db.Update(query.NewQuery(_constants.AppsCollection).Where(query.Field("name").Eq(metadata.Name)), structs.Map(metadata)); err != nil {
      slog.Error(fmt.Sprintf("failed to update app %s: %v", metadata.Name, err))

      return c.String(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusOK)
  }
}
