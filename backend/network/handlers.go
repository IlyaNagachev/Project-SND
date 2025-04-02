package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"snd/backend/store"
	"snd/backend/store/models"
	"strconv"
)

func createProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var err error
		var bytes []byte
		if bytes, err = io.ReadAll(request.Body); err != nil {
			writer.Write([]byte("Ошибка" + err.Error()))
			return
		}
		var product *models.Product

		if err = json.Unmarshal(bytes, &product); err != nil {
			writer.Write([]byte("Ошибка" + err.Error()))
			return
		}

		if err = store.Default.Products.New(product); err != nil {
			writer.Write([]byte("Ошибка" + err.Error()))
			return
		}
	}
}

func getProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var err error
		var products []models.Product

		if products, err = store.Default.Products.GetAll(); err != nil {
			writer.Write([]byte("Ошибка" + err.Error()))
			return
		}

		var bytes []byte

		if bytes, err = json.Marshal(products); err != nil {
			writer.Write([]byte("Ошибка" + err.Error()))
			return
		}

		writer.Write(bytes)
	}
}

func getProductById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var err error
		var product models.Product
		var in = struct {
			Uuid int `json:"uuid"`
		}{}
		var inbytes []byte
		if inbytes, err = io.ReadAll(request.Body); err != nil {
			writer.Write([]byte("Ошибка" + err.Error()))
			return
		}

		if err = json.Unmarshal(inbytes, &in); err != nil {
			writer.Write([]byte("Ошибка" + err.Error()))
			return
		}

		fmt.Println(in)

		if product, err = store.Default.Products.Get(strconv.Itoa(in.Uuid)); err != nil {
			writer.Write([]byte("Ошибка" + err.Error()))
			return
		}

		var bytes []byte

		if bytes, err = json.Marshal(product); err != nil {
			writer.Write([]byte("Ошибка" + err.Error()))
			return
		}

		writer.Write(bytes)
	}
}
