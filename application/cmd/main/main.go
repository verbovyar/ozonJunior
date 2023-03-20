package main

import (
	"log"
	"pkgs/internal/app"
	"pkgs/pkg/config"
)

func main() {
	conf, err := config.LoadConfig("././pkg/config")

	if err != nil {
		log.Fatalf("%v", err)
	}

	app.Run(conf)
}
