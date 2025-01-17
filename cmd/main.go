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

  serverPort := config.ServerCfg.Port
  
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: http.HandlerFunc(router),
	}

	log.Printf("Running server on http://127.0.0.1%s\n", server.Addr)

	log.Fatal(server.ListenAndServe())
}
