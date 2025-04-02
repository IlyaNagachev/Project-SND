package store

import (
	"context"
	"snd/backend/store/adapters/scylla"
	"snd/backend/store/products"
)

func NewStore(config *Config) (s *Store, err error) {
	s = &Store{}
	if s.scylla, err = scylla.NewAdapter(config.Scylla); err != nil {
		return
	}
	if s.Products, err = products.NewStorage(context.Background(),s.scylla.GetSession(), s.scylla.GetTables()); err != nil {
		return
	}
	return
}

type Store struct {
	scylla *scylla.Adapter
	Products *products.Storage
}

var Default *Store