package main

import (
	"fmt"
	"log"
	"os"

	solidlamp "github.com/PeterJohnBishop/solid-lamp"
	"github.com/PeterJohnBishop/solid-lamp/cu"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("API_KEY")
	client := solidlamp.NewClickUpClient(key)
	params := cu.GetTaskQueryParams{}
	task, err := client.GetTask("868fntvay", params)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	fmt.Printf("Found task %s", task)
}
