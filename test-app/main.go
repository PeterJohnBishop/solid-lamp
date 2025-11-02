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
	task, err := client.GetTask("868fntvay", nil)
	if err != nil {
		fmt.Println("Error!")
	}
	fmt.Printf("Task %s found", task.Name)
}
