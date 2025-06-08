package storage

import (
  "fmt"
  configutilities "github.com/kieranroneill/yggdrasil/internal/utilities/config"
  "github.com/ostafen/clover/v2"
  "log/slog"
)

func SetupDatabase() (*clover.DB, error) {
  configPath, err := configutilities.DefaultConfigPath()
  if err != nil {
    return nil, err
  }

  // open connection to the database
  db, err := clover.Open(configPath)
  if err != nil {
    return nil, err
  }

  // defer db connect closure
  defer func(db *clover.DB) {
    if err = db.Close(); err != nil {
      slog.Error(fmt.Sprintf("error while closing the db: %v", err))
    }
  }(db)

  return db, nil
}
