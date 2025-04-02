package network

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func CreateRouter() http.Handler {
	var router = chi.NewRouter()
	router.Use(validateRequest)
	router.Post("/user/info", userInfo())
	router.Post("/user/create", createProduct())
	router.Post("/user/get", getProduct())
	router.Post("/product/id", getProductById())
	return router
}

func userInfo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Info().Msg(request.RequestURI)
	}
}

func validateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var origin string
		var allowed_headers = strings.Join([]string{
			"origin",
			"content-type",
		}, ", ")
		if origin = r.Header.Get("Origin"); origin == "" {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", allowed_headers)
		w.Header().Set("Access-Control-Expose-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
