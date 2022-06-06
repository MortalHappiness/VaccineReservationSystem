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
		return nil, fmt.Errorf("Could not read row with key %s: %w", rowKey, err)
	}

	if len(row) == 0 {
		return nil, nil
	}

	return row, nil
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

func (vaccineClient *VaccineClient) CreateOrUpdateUser(ID string, attributes map[string]string) error {
	return vaccineClient.createOrUpdate("user#"+ID, "user", attributes)
}

func (vaccineClient *VaccineClient) GetUser(ID string) (bigtable.Row, error) {
	return vaccineClient.get("user#" + ID)
}

func (vaccineClient *VaccineClient) DeleteUser(ID string) error {
	return vaccineClient.delete("user#" + ID)
}

func (vaccineClient *VaccineClient) CreateOrUpdateHospital(ID string, attributes map[string]string) error {
	return vaccineClient.createOrUpdate("hospital#"+ID, "hospital", attributes)
}

func (vaccineClient *VaccineClient) GetHospital(ID string) (bigtable.Row, error) {
	return vaccineClient.get("hospital#" + ID)
}

func (vaccineClient *VaccineClient) DeleteHospital(ID string) error {
	return vaccineClient.delete("hospital#" + ID)
}

func (vaccineClient *VaccineClient) CreateOrUpdateReservation(ID string, attributes map[string]string) error {
	return vaccineClient.createOrUpdate("reservation#"+ID, "reservation", attributes)
}

func (vaccineClient *VaccineClient) GetReservation(ID string) (bigtable.Row, error) {
	return vaccineClient.get("reservation#" + ID)
}

func (vaccineClient *VaccineClient) DeleteReservation(ID string) error {
	return vaccineClient.delete("reservation#" + ID)
}
