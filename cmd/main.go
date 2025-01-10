package main

import (
	"encoding/json"
	"log"

	"github.com/ghostriderdev/housierBack/internal/models"
)

func main() {
	log.Println("Starting app")
	user := models.User{
		Email: "Olvadis",
		Password: "12345",
		Role: models.ROLE_USER,
	}

	decoded, err := json.Marshal(user)

	if err != nil {
		panic("Error to marshal user to json")
	}

	log.Printf("User created: %s", string(decoded))
}
