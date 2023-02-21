package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

func main() {
	// Connect Google Cloud
	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "ta-tracking-f43e5"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// Execute
	// addData(client, ctx)
	// readData(client, ctx)
	setData(client, ctx)
}

// Adding data
func addData(client *firestore.Client, ctx context.Context) {
	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Leon",
		"last":  "Harryman",
		"born":  1990,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

func setData(client *firestore.Client, ctx context.Context) {
		_, err := client.Collection("cities").Doc("SF").Set(ctx, map[string]interface{}{
			"name":    "San Francisco",
			"state":   "CA",
			"country": "USA",
			"long": "124",
			"lat": "322",
	})
	if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
	}
}

func readData(client *firestore.Client, ctx context.Context) {
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
						break
		}
		if err != nil {
						log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}