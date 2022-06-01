package vaccineclient

import (
	"context"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/bigtable"
	"github.com/joho/godotenv"
)

func setup() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Print(".env file not found!")
	}
}

func TestUserOperations(t *testing.T) {
	setup()

	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")
	instanceID := os.Getenv("INSTANCE_ID")
	if projectID == "" || instanceID == "" {
		t.Fatal("Skipping vaccineclient test. Set PROJECT_ID and INSTANCE_ID.")
	}

	adminClient, err := bigtable.NewAdminClient(ctx, projectID, instanceID)
	if err != nil {
		t.Fatalf("bigtable.NewAdminClient: %v", err)
	}

	tableName := "vaccine-reservation-system-test"
	adminClient.DeleteTable(ctx, tableName)
	adminClient.CreateTable(ctx, tableName)

	columnFamilyName := "user"
	if err := adminClient.CreateColumnFamily(ctx, tableName, columnFamilyName); err != nil {
		adminClient.DeleteTable(ctx, tableName)
		t.Fatalf("CreateColumnFamily(%s): %v", columnFamilyName, err)
	}

	vaccineClient := NewVaccineClient(projectID, instanceID, tableName)

	if err = vaccineClient.CreateUser("AliceNationID", "Alice", "AliceHealthCardID", "Female", "1970/1/1", "AliceAddress", "AlicePhone", "BNT,AZ"); err != nil {
		t.Errorf("TestCreateUser: %v", err)
	}

	adminClient.DeleteTable(ctx, tableName)
}
