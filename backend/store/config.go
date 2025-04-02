package store

import "snd/backend/store/adapters/scylla"

type Config struct {
	Scylla *scylla.Config `yaml:"scylla"`
}
