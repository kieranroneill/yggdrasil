package storage

import (
	"github.com/ostafen/clover/v2"
)

// CreateCollectionIfNoneExists Checks if the collection exists, if it doesn't it creates one.
//
// Parameters:
//   - db: An instance of the clover db.
//   - collection: The name of the collection to create.
//
// Returns:
//   - An error if something went wrong.
func CreateCollectionIfNoneExists(db *clover.DB, collection string) error {
	hasCollection, err := db.HasCollection(collection)
	if err != nil {
		return err
	}

	// if the collection exists ignore
	if hasCollection {
		return nil
	}

	return db.CreateCollection(collection)
}
