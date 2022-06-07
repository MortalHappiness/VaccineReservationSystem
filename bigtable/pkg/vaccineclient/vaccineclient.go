package vaccineclient

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"cloud.google.com/go/bigtable"
)

type VaccineClient struct {
	projectID  string
	instanceID string
	tableName  string
}

func NewVaccineClient(projectID string, instanceID string, tableName string) *VaccineClient {
	vaccineclient := &VaccineClient{
		projectID:  projectID,
		instanceID: instanceID,
		tableName:  tableName,
	}
	return vaccineclient
}

// Utility function to print out a row
//
// Example Usage:
//
// vaccineclient.PrintRow(row)
func PrintRow(row bigtable.Row) {
	if row == nil {
		println("Empty Row\n")
		return
	}
	fmt.Printf("Reading data for %s:\n", row.Key())
	for columnFamily, cols := range row {
		fmt.Printf("Column Family %s\n", columnFamily)
		for _, col := range cols {
			qualifier := col.Column[strings.IndexByte(col.Column, ':')+1:]
			fmt.Printf("\t%s: %s @%d\n", qualifier, col.Value, col.Timestamp)
		}
	}
	println()
}

func (vaccineClient *VaccineClient) createOrUpdate(rowKey string, columnFamilyName string, attributes map[string]string) error {
	if len(attributes) == 0 {
		return errors.New("attributes cannot be empty")
	}

	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, vaccineClient.projectID, vaccineClient.instanceID)
	if err != nil {
		return fmt.Errorf("bigtable.NewClient: %v", err)
	}
	defer client.Close()
	tbl := client.Open(vaccineClient.tableName)

	timestamp := bigtable.Now()

	mut := bigtable.NewMutation()

	for columnName, value := range attributes {
		mut.Set(columnFamilyName, columnName, timestamp, []byte(value))
	}

	if err := tbl.Apply(ctx, rowKey, mut); err != nil {
		return fmt.Errorf("apply: %v", err)
	}

	return nil
}

func (vaccineClient *VaccineClient) get(rowKey string) (bigtable.Row, error) {
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, vaccineClient.projectID, vaccineClient.instanceID)
	if err != nil {
		return nil, fmt.Errorf("bigtable.NewClient: %w", err)
	}
	defer client.Close()
	tbl := client.Open(vaccineClient.tableName)
	row, err := tbl.ReadRow(ctx, rowKey, bigtable.RowFilter(bigtable.LatestNFilter(1)))
	if err != nil {
		return nil, fmt.Errorf("could not read row with key %s: %w", rowKey, err)
	}

	if len(row) == 0 {
		return nil, nil
	}

	return row, nil
}

func (vaccineClient *VaccineClient) getPrefix(rowKeyPrefix string) ([]bigtable.Row, error) {
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, vaccineClient.projectID, vaccineClient.instanceID)
	if err != nil {
		return nil, fmt.Errorf("bigtable.NewClient: %w", err)
	}
	defer client.Close()
	tbl := client.Open(vaccineClient.tableName)
	var rows []bigtable.Row
	err = tbl.ReadRows(ctx, bigtable.PrefixRange(rowKeyPrefix),
		func(row bigtable.Row) bool {
			rows = append(rows, row)
			return true
		},
		bigtable.RowFilter(bigtable.LatestNFilter(1)))

	if err != nil {
		return nil, fmt.Errorf("could not read rows with prefix %s: %w", rowKeyPrefix, err)
	}

	if err = client.Close(); err != nil {
		return nil, fmt.Errorf("client.Close(): %v", err)
	}

	return rows, nil
}

func (vaccineClient *VaccineClient) delete(rowKey string) error {
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, vaccineClient.projectID, vaccineClient.instanceID)
	if err != nil {
		return fmt.Errorf("bigtable.NewClient: %v", err)
	}
	defer client.Close()
	tbl := client.Open(vaccineClient.tableName)

	mut := bigtable.NewMutation()
	mut.DeleteRow()

	if err := tbl.Apply(ctx, rowKey, mut); err != nil {
		return fmt.Errorf("apply: %v", err)
	}

	return nil
}

// Create or update a user, attributes map cannot be empty
//
// Example Usage (Create):
//
// err := vaccineClient.CreateOrUpdateUser("A123456789", map[string]string{
// 	"name":         "Alice",
// 	"healthCardID": "000011112222",
// 	"gender":       "Female",
// 	"birthday":     "2022/05/23",
// 	"address":      "No.2, Sec. 4, Roosevelt Road, Taipei, 10617 Taiwan",
// 	"phone":        "0912345678",
// 	"vaccines":     "BNT,AZ",
// })
//
// Example Usage (Update):
//
// err := vaccineClient.CreateOrUpdateUser("A123456789", map[string]string{
// 	"name":         "Alice1",
// })
func (vaccineClient *VaccineClient) CreateOrUpdateUser(ID string, attributes map[string]string) error {
	rowKey := "user#" + ID
	return vaccineClient.createOrUpdate(rowKey, "user", attributes)
}

// Get a user
//
// Example Usage:
//
// row, err := vaccineClient.GetUser("A123456789")
func (vaccineClient *VaccineClient) GetUser(ID string) (bigtable.Row, error) {
	rowKey := "user#" + ID
	return vaccineClient.get(rowKey)
}

// Delete a user
//
// Example Usage:
//
// err := vaccineClient.DeleteUser("A123456789")
func (vaccineClient *VaccineClient) DeleteUser(ID string) error {
	rowKey := "user#" + ID
	return vaccineClient.delete(rowKey)
}

func (vaccineClient *VaccineClient) CreateOrUpdateHospital(ID string, county string, township string, attributes map[string]string) error {
	rowKey := fmt.Sprintf("hospital#%s#%s#%s", county, township, ID)
	return vaccineClient.createOrUpdate(rowKey, "hospital", attributes)
}

func (vaccineClient *VaccineClient) GetHospital(ID string, county string, township string) (bigtable.Row, error) {
	rowKey := fmt.Sprintf("hospital#%s#%s#%s", county, township, ID)
	return vaccineClient.get(rowKey)
}

func (vaccineClient *VaccineClient) GetHospitals(county string, township string) ([]bigtable.Row, error) {
	rowKeyPrefix := fmt.Sprintf("hospital#%s#%s", county, township)
	return vaccineClient.getPrefix(rowKeyPrefix)
}

func (vaccineClient *VaccineClient) ListHospitals() ([]bigtable.Row, error) {
	rowKeyPrefix := fmt.Sprintf("hospital")
	return vaccineClient.getPrefix(rowKeyPrefix)
}

func (vaccineClient *VaccineClient) DeleteHospital(ID string, county string, township string) error {
	rowKey := fmt.Sprintf("hospital#%s#%s#%s", county, township, ID)
	return vaccineClient.delete(rowKey)
}

func (vaccineClient *VaccineClient) CreateOrUpdateReservation(ID string, userID string, attributes map[string]string) error {
	rowKey := fmt.Sprintf("reservation#%s#%s", userID, ID)
	return vaccineClient.createOrUpdate(rowKey, "reservation", attributes)
}

func (vaccineClient *VaccineClient) GetReservation(ID string, userID string) (bigtable.Row, error) {
	rowKey := fmt.Sprintf("reservation#%s#%s", userID, ID)
	return vaccineClient.get(rowKey)
}

func (vaccineClient *VaccineClient) GetReservations(userID string) ([]bigtable.Row, error) {
	rowKeyPrefix := fmt.Sprintf("reservation#%s", userID)
	return vaccineClient.getPrefix(rowKeyPrefix)
}

func (vaccineClient *VaccineClient) DeleteReservation(ID string, userID string) error {
	rowKey := fmt.Sprintf("reservation#%s#%s", userID, ID)
	return vaccineClient.delete(rowKey)
}
