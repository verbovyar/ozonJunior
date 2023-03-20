package app

import (
	"log"
	"net/http"
	"pkgs/internal/handlers"
	"pkgs/internal/repositories/datas"
	"pkgs/pkg/config"
)

func Run(config config.Config) {
	repository := datas.New()
	server := handlers.New(repository)
	runHTTPServer(server, config.Port)
}

func runHTTPServer(server *handlers.Server, port string) {
	err := http.ListenAndServe(port, server.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
