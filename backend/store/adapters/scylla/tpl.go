package scylla

import (
	"github.com/scylladb/gocqlx/v3/table"
	"snd/backend/store/models"
)

var productMetadata = table.Metadata{
	Name		: models.ProductTable,
	Columns		: []string{"uuid", "class", "name", "price", "img", "description", "specs"},
	PartKey     : []string{"uuid"},
	SortKey     : []string{"uuid"},
}

