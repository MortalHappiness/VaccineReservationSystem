package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"cloud.google.com/go/bigtable"
	"github.com/joho/godotenv"

	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
)

func readCsvFile(filePath string) ([]string, [][]string) {
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

	headers := records[0]
	rows := records[1:]

	return headers, rows
}

func insertUserFromCsvFile(vaccineClient *vaccineclient.VaccineClient, filePath string) {
	fmt.Println("inserting users from path: ", filePath)
	headers, rows := readCsvFile(filePath)
	for index, row := range rows {
		ID := row[0]
		attributes := make(map[string]string)
		for i := 1; i < len(row); i++ {
			key := headers[i]
			value := row[i]
			attributes[key] = value
		}
		vaccineClient.CreateOrUpdateUser(ID, attributes)
		if index % 100 == 0 {
			fmt.Println("user inserted: ", index)
		}
	}
}

func batchInsertUserFromCsvFile(vaccineClient *vaccineclient.VaccineClient, filePath string) {
	fmt.Println("batch inserting users from path: ", filePath)
	headers, rows := readCsvFile(filePath)

	IDs := []string{}
	attributes := []map[string]string{}

	for _, row := range rows {
		ID := row[0]
		attribute := make(map[string]string)
		for i := 1; i < len(row); i++ {
			key := headers[i]
			value := row[i]
			attribute[key] = value
		}
		IDs = append(IDs, ID)
		attributes = append(attributes, attribute)
	}
	vaccineClient.BatchCreateUsers(IDs, attributes)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func insertUserRandom(vaccineClient *vaccineclient.VaccineClient, N int) {
	headers := [7]string{"name", "healthCardID", "gender", "birthDay", "address", "phone", "vaccines"}
	gender := [2]string{"Female", "Male"}
	addresses := [3]string{"花蓮縣鳳林鎮中長路17號", "高雄市前鎮區中三路15號", "桃園市中壢區吳鳳二街15號"}
	for i := 0; i < N; i += 1 {
		ID := string(byte(randomInt(65, 90))) + strconv.Itoa(rand.Intn(2))
		for j := 0; j < 8; j += 1 {
			ID = ID + strconv.Itoa(rand.Intn(10))
		}
		fmt.Printf("ID: %s\n", ID)
		attributes := make(map[string]string)
		attributes[headers[0]] = "Allen"
		HealthCardID := ""
		for j := 0; j < 12; j += 1 {
			HealthCardID = HealthCardID + strconv.Itoa(rand.Intn(10))
		}
		attributes[headers[1]] = HealthCardID
		attributes[headers[2]] = gender[rand.Intn(2)]
		attributes[headers[3]] = fmt.Sprintf("%d/%02d/%02d", randomInt(1925, 2020), randomInt(1, 12), randomInt(1, 28))
		attributes[headers[4]] = addresses[rand.Intn(3)]
		phone := "09"
		for j := 0; j < 8; j += 1 {
			phone = phone + strconv.Itoa(rand.Intn(9))
		}
		attributes[headers[5]] = phone
		attributes[headers[6]] = ""

		vaccineClient.CreateOrUpdateUser(ID, attributes)
	}
}

func insertHospitalFromCsvFile(vaccineClient *vaccineclient.VaccineClient, filePath string) {
	fmt.Println("inserting hospitals from path: ", filePath)
	headers, rows := readCsvFile(filePath)
	for index, row := range rows {
		ID := row[0]
		county := row[1]
		township := row[2]
		attributes := make(map[string]string)
		for i := 3; i < len(row); i++ {
			key := headers[i]
			value := row[i]
			attributes[key] = value
		}
		vaccineClient.CreateOrUpdateHospital(ID, county, township, attributes)
		if index % 100 == 0 {
			fmt.Println("hospital inserted: ", index)
		}
	}
}

func batchInsertHospitalFromCsvFile(vaccineClient *vaccineclient.VaccineClient, filePath string) {
	fmt.Println("batch inserting hospitals from path: ", filePath)
	headers, rows := readCsvFile(filePath)
	IDs := []string{}
	counties := []string{}
	townships := []string{}
	attributes := []map[string]string{}
	for _, row := range rows {
		ID := row[0]
		county := row[1]
		township := row[2]
		attribute := make(map[string]string)
		for i := 3; i < len(row); i++ {
			key := headers[i]
			value := row[i]
			attribute[key] = value
		}
		IDs = append(IDs, ID)
		counties = append(counties, county)
		townships = append(townships, township)
		attributes = append(attributes, attribute)
	}
	vaccineClient.BatchCreateHospitals(IDs, counties, townships, attributes)
}

func insertAllHospitalFromCsvFile(vaccineClient *vaccineclient.VaccineClient) {
	filePath := "data/hospitals_all_clean.csv"
	headers, rows := readCsvFile(filePath)
	for _, row := range rows {
		ID := row[0]
		county := row[1]
		township := row[2]
		attributes := make(map[string]string)
		for i := 3; i < len(row); i++ {
			key := headers[i]
			value := row[i]
			attributes[key] = value
		}
		vaccineClient.CreateOrUpdateHospital(ID, county, township, attributes)
	}
}

func insertReservationFromCsvFile(vaccineClient *vaccineclient.VaccineClient) {
	filePath := "data/reservation.csv"
	headers, rows := readCsvFile(filePath)
	for _, row := range rows {
		ID := row[0]
		userID := row[1]
		attributes := make(map[string]string)
		for i := 2; i < len(row); i++ {
			key := headers[i]
			value := row[i]
			attributes[key] = value
		}
		vaccineClient.CreateOrUpdateReservation(ID, userID, attributes)
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

	// insertUserFromCsvFile(vaccineClient, "data/user.csv")
	// insertHospitalFromCsvFile(vaccineClient, "data/hospital.csv")
	// insertReservationFromCsvFile(vaccineClient)
	// insertUserRandom(vaccineClient, 10)
	
	batchInsertUserFromCsvFile(vaccineClient, "data/user_10000.csv")
	batchInsertHospitalFromCsvFile(vaccineClient, "data/hospitals_all_clean.csv")
}
