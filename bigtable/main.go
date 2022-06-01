package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/bigtable"
	"github.com/joho/godotenv"

	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
)

const tableName = "vaccine-reservation-system"

func main() {
	// Read environment variables
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

	// Setup tables and column families
	ctx := context.Background()

	adminClient, err := bigtable.NewAdminClient(ctx, projectID, instanceID)
	if err != nil {
		log.Fatalf("Could not create admin client: %v", err)
	}

	adminClient.DeleteTable(ctx, tableName)
	adminClient.CreateTable(ctx, tableName)

	adminClient.CreateColumnFamily(ctx, tableName, "user")
	adminClient.CreateColumnFamily(ctx, tableName, "hospital")
	adminClient.CreateColumnFamily(ctx, tableName, "reservation")

	// Close client
	if err = adminClient.Close(); err != nil {
		log.Fatalf("Could not close admin client: %v", err)
	}

	// Insert data
	vaccineClient := vaccineclient.NewVaccineClient(projectID, instanceID, tableName)

	vaccineClient.CreateOrUpdateUser("AliceNationID", "Alice", "AliceHealthCardID", "Female", "1970/1/1", "AliceAddress", "AlicePhone", "BNT,AZ")

	// row, _ := vaccineClient.GetUser("AliceNationID")
	// vaccineclient.PrintRow(row)
	// vaccineClient.CreateOrUpdateUser("AliceNationID", "Alice1", "", "", "", "", "", "")
	// row, _ = vaccineClient.GetUser("AliceNationID")
	// vaccineclient.PrintRow(row)
	// vaccineClient.DeleteUser("AliceNationID")
	// row, _ = vaccineClient.GetUser("AliceNationID")
	// vaccineclient.PrintRow(row)
}
