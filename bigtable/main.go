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

	tableID := "my-table"

	ctx := context.Background()

	// Set up Bigtable data operations client.
	client, err := bigtable.NewClient(ctx, projectID, instanceID)
	if err != nil {
		log.Fatalf("Could not create data operations client: %v", err)
	}

	tbl := client.Open(tableID)

	// Read data in a row using a row key
	rowKey := "r1"
	columnFamilyName := "cf1"

	log.Printf("Getting a single row by row key:")
	row, err := tbl.ReadRow(ctx, rowKey)
	if err != nil {
		log.Fatalf("Could not read row with key %s: %v", rowKey, err)
	}
	log.Printf("Row key: %s\n", rowKey)
	log.Printf("Data: %s\n", string(row[columnFamilyName][0].Value))

	if err = client.Close(); err != nil {
		log.Fatalf("Could not close data operations client: %v", err)
	}
}
