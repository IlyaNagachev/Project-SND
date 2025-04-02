package snd

import (
	"gopkg.in/yaml.v3"
	"os"
	"snd/backend/logs"
	"snd/backend/network"
	"snd/backend/store"
)

type Config struct {
	Network *network.Config `yaml:"network"`
	Store   *store.Config `yaml:"store"`
	Server  any `yaml:"server"`
	Logger *logs.Config `yaml:"logger"`
}

func NewConfig(path *string) (config *Config, err error) {
	var bytes []byte
	if bytes, err = os.ReadFile(*path); err != nil {return}

	if err = yaml.Unmarshal(bytes, &config); err != nil {return}

	return
}

