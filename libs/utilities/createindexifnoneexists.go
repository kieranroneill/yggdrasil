package utilities

import (
	"github.com/ostafen/clover/v2"
)

// CreateIndexIfNoneExists Checks if the index exists, if it doesn't it creates one.
//
// Parameters:
//   - db: An instance of the clover db.
//   - collection: The name of the collection.
//   - field: The name of the field to apply the index.
//
// Returns:
//   - An error if something went wrong.
func CreateIndexIfNoneExists(db *clover.DB, collection string, field string) error {
	hasIndex, err := db.HasIndex(collection, field)
	if err != nil {
		return err
	}

	// if the index exists ignore
	if hasIndex {
		return nil
	}

	return db.CreateIndex(collection, field)
}
