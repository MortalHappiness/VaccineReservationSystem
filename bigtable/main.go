package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"cloud.google.com/go/bigtable"
	"github.com/joho/godotenv"

	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func insertDataFromCsvFile(vaccineClient *vaccineclient.VaccineClient, filePath string, collection string) {
	records := readCsvFile(filePath)
	headers := records[0]
	rows := records[1:]
	for _, row := range rows {
		ID := row[0]
		attributes := make(map[string]string)
		for j, value := range row[1:] {
			key := headers[j+1]
			attributes[key] = value
		}
		vaccineClient.CreateOrUpdate(collection+"#"+ID, collection, attributes)
	}
}

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

	tableName, present := os.LookupEnv("TABLE_NAME")
	if !present {
		log.Fatal("TABLE_NAME")
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

	insertDataFromCsvFile(vaccineClient, "data/user.csv", "user")
	insertDataFromCsvFile(vaccineClient, "data/hospital.csv", "hospital")
	insertDataFromCsvFile(vaccineClient, "data/reservation.csv", "reservation")
}
