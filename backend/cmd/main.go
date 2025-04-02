package main

import (
	"encoding/json"
	"flag"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	snd "snd/backend"
	"snd/backend/logs"
	"snd/backend/network"
	"snd/backend/store"
	"snd/backend/store/models"
)

var (
	config *snd.Config
	configPath = flag.String("config", "./snd/templates/config.yaml", "config File Path")
	err error
)

func init() {
	flag.Parse()
	if config, err = snd.NewConfig(configPath); err != nil {
		panic(err)
	}
	if err = logs.NewLogger(config.Logger); err != nil {
		panic(err)
	}
}


func main() {
	log.Info().Msg("Hello")
	var server http.Server
	var router http.Handler
	router = network.CreateRouter()
	server = network.NewServer(config.Network, router)
	if store.Default, err = store.NewStore(config.Store); err != nil {
		panic(err)
	}

	if p,_ := store.Default.Products.GetAll(); len(p)<4 {
		var file []byte
		if file, err = os.ReadFile("./snd/products.json"); err != nil {
			panic(err)
		}
		var products []models.Product
		if err = json.Unmarshal(file, &products); err != nil {
			panic(err)
		}
		for _,p :=range products {
			if err = store.Default.Products.New(&p); err != nil {
				panic(err)
			}
		}
	}




	server.ListenAndServe()
}

