package main

import (
	"fmt"
	"log"
	"os"

	solidlamp "github.com/PeterJohnBishop/solid-lamp"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("API_KEY")
	client := solidlamp.NewClickUpClient(key)
	user, err := client.GetAuthorizedUser()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	fmt.Printf("Found user %s", user.Username)
}
