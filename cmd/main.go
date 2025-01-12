package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ghostriderdev/housierBack/cfg"
)

func main() {
	config, err := cfg.LoadConfig()

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ServerCfg.Port),
		Handler: http.HandlerFunc(router),
	}

	log.Printf("Running server on http://127.0.0.1:%d\n", config.ServerCfg.Port)

	log.Fatal(server.ListenAndServe())
}
