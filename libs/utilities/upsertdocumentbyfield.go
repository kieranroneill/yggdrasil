package utilities

import (
	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
)

func UpsertDocumentByField(db *clover.DB, collection string, field string, value interface{}, data interface{}) (*document.Document, error) {
	_document, err := db.FindFirst(query.NewQuery(collection).Where(query.Field(field).Eq(value)))
	if err != nil {
		return nil, err
	}

	// if no document exists, insert a new one
	if _document == nil {
		_document = document.NewDocumentOf(data)
		if err = db.Insert(collection, _document); err != nil {
			return nil, err
		}

		return _document, nil
	}

	if err = db.Update(query.NewQuery(collection).Where(query.Field(field).Eq(value)), data.(map[string]interface{})); err != nil {
		return nil, err
	}

	return _document, nil
}
