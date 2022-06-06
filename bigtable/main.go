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

	vaccineClient := vaccineclient.NewVaccineClient(projectID, instanceID, tableName)

	// Insert user data
	vaccineClient.CreateOrUpdateUser("A123456789", map[string]string{
		"name":         "Alice",
		"healthCardID": "000011112222",
		"gender":       "Female",
		"birthday":     "2022/05/23",
		"address":      "No.2, Sec. 4, Roosevelt Road, Taipei, 10617 Taiwan",
		"phone":        "0912345678",
		"vaccines":     "BNT,AZ",
	})
	vaccineClient.CreateOrUpdateUser("B223456789", map[string]string{
		"name":         "Bob",
		"healthCardID": "012311113333",
		"gender":       "Male",
		"birthday":     "2022/01/25",
		"address":      "No.2, Sec. 4, Roosevelt Road, Taipei, 10617 Taiwan",
		"phone":        "0987654321",
		"vaccines":     "Moderna,BNT",
	})

	// Insert hospital data
	vaccineClient.CreateOrUpdateHospital("1", map[string]string{
		"name":       "Hospital1",
		"city":       "Taipei",
		"township":   "Da'an District",
		"address":    "example address 1",
		"vaccineCnt": "BNT:1,AZ:2",
	})

	// Insert Reservation data
	vaccineClient.CreateOrUpdateReservation("1", map[string]string{
		"userID":      "A123456789",
		"hospitalID":  "1",
		"vaccineType": "BNT",
		"datetime":    "1654524434",
		"completed":   "1",
	})

	// Debug
	// row, _ := vaccineClient.GetUser("A123456789")
	// vaccineclient.PrintRow(row)
	// vaccineClient.CreateOrUpdateUser("A123456789", map[string]string{
	// 	"name": "Alice1",
	// })
	// row, _ = vaccineClient.GetUser("A123456789")
	// vaccineclient.PrintRow(row)
	// vaccineClient.DeleteUser("A123456789")
	// row, _ = vaccineClient.GetUser("A123456789")
	// vaccineclient.PrintRow(row)
}
