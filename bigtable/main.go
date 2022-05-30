// Quickstart is a sample program demonstrating use of the Cloud Bigtable client
// library to read a row from an existing table.
package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/bigtable"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env file not found!")
	}

	projectID, present := os.LookupEnv("PROJECT_ID")
	if !present {
		log.Fatal("Environment variable not found: PROJECT_ID")
	}

	instanceID, present := os.LookupEnv("INSTANCE_ID")
	if !present {
		log.Fatal("Environment variable not found: INSTANCE_ID")
	}

	ctx := context.Background()

	adminClient, err := bigtable.NewAdminClient(ctx, projectID, instanceID)
	if err != nil {
		log.Fatalf("Could not create admin client: %v", err)
	}
	client, err := bigtable.NewClient(ctx, projectID, instanceID)
	if err != nil {
		log.Fatalf("Could not create data operations client: %v", err)
	}

	adminClient.DeleteTable(ctx, "User")
	adminClient.DeleteTable(ctx, "Reservation")
	adminClient.DeleteTable(ctx, "Hospital")

	adminClient.CreateTable(ctx, "User")
	adminClient.CreateTable(ctx, "Reservation")
	adminClient.CreateTable(ctx, "Hospital")

	adminClient.CreateColumnFamily(ctx, "User", "user")

	var mutation *bigtable.Mutation

	userTable := client.Open("User")
	mutation = bigtable.NewMutation()
	mutation.Set("user", "name", bigtable.Now(), []byte("Alice"))
	err = userTable.Apply(ctx, "alice-id", mutation)
	if err != nil {
		log.Fatalf("Could not apply row mutation: %v", err)
	}

	if err = adminClient.Close(); err != nil {
		log.Fatalf("Could not close admin client: %v", err)
	}
	if err = client.Close(); err != nil {
		log.Fatalf("Could not close data operations client: %v", err)
	}
}
