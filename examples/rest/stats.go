package main

//go run stats.go utils.go constants.go

import (
	"context"
	"fmt"
	"os"

	"github.com/ably/ably-go/ably"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Connect to Ably using the API key and ClientID
	client, err := ably.NewREST(
		ably.WithKey(os.Getenv(AblyKey)),
		ably.WithClientID(UserName))
	if err != nil {
		panic(err)
	}

	printApplicationStats(client)
}

func printApplicationStats(client *ably.REST) {
	page, err := client.Stats(context.Background(), &ably.PaginateParams{})
	for ; err == nil && page != nil; page, err = page.Next(context.Background()) {
		for _, stat := range page.Stats() {
			fmt.Println(jsonify(stat))
		}
	}
	if err != nil {
		err := fmt.Errorf("error getting application stats %w", err)
		panic(err)
	}
}
