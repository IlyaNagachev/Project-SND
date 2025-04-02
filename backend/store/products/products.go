package products

import (
	"context"
	"github.com/scylladb/gocqlx/v3"
	"snd/backend/store/models"
)

type Storage struct {
	ctx context.Context
	session gocqlx.Session
	Tables *models.Tables
}

func NewStorage(ctx context.Context, session gocqlx.Session, tables *models.Tables) (s *Storage, err error) {
	return &Storage{ctx: ctx, session: session, Tables: tables}, nil
}

