package network

import (
	"fmt"
	"net/http"
)

func NewServer(config *Config, h http.Handler) http.Server {
	return http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Address, config.Port),
		Handler: h,
	}
}
