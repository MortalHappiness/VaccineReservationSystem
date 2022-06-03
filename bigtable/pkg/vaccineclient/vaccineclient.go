package vaccineclient

import (
	"context"
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

func setIfNotEmptyString(mut *bigtable.Mutation, columnFamilyName string, columnName string, timestamp bigtable.Timestamp, value string) {
	if value != "" {
		mut.Set(columnFamilyName, columnName, timestamp, []byte(value))
	}
}

func (vaccineClient *VaccineClient) CreateOrUpdateUser(nationID string, name string, healthCardID string, gender string, birthday string, address string, phone string, vaccines string) error {
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, vaccineClient.projectID, vaccineClient.instanceID)
	if err != nil {
		return fmt.Errorf("bigtable.NewClient: %v", err)
	}
	defer client.Close()
	tbl := client.Open(vaccineClient.tableName)

	columnFamilyName := "user"
	timestamp := bigtable.Now()

	mut := bigtable.NewMutation()

	setIfNotEmptyString(mut, columnFamilyName, "name", timestamp, name)
	setIfNotEmptyString(mut, columnFamilyName, "healthCardID", timestamp, healthCardID)
	setIfNotEmptyString(mut, columnFamilyName, "gender", timestamp, gender)
	setIfNotEmptyString(mut, columnFamilyName, "birthday", timestamp, birthday)
	setIfNotEmptyString(mut, columnFamilyName, "address", timestamp, address)
	setIfNotEmptyString(mut, columnFamilyName, "phone", timestamp, phone)
	setIfNotEmptyString(mut, columnFamilyName, "vaccines", timestamp, vaccines)

	rowKey := "user#" + nationID
	if err := tbl.Apply(ctx, rowKey, mut); err != nil {
		return fmt.Errorf("apply: %v", err)
	}

	return nil
}

func (vaccineClient *VaccineClient) GetUser(nationID string) (bigtable.Row, error) {
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, vaccineClient.projectID, vaccineClient.instanceID)
	if err != nil {
		return nil, fmt.Errorf("bigtable.NewClient: %w", err)
	}
	defer client.Close()
	tbl := client.Open(vaccineClient.tableName)
	rowkey := "user#" + nationID
	row, err := tbl.ReadRow(ctx, rowkey, bigtable.RowFilter(bigtable.LatestNFilter(1)))
	if err != nil {
		return nil, fmt.Errorf("Could not read row with key %s: %w", rowkey, err)
	}

	if len(row) == 0 {
		return nil, nil
	}

	return row, nil
}

func (vaccineClient *VaccineClient) DeleteUser(nationID string) error {
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, vaccineClient.projectID, vaccineClient.instanceID)
	if err != nil {
		return fmt.Errorf("bigtable.NewClient: %v", err)
	}
	defer client.Close()
	tbl := client.Open(vaccineClient.tableName)

	mut := bigtable.NewMutation()
	mut.DeleteRow()

	rowKey := "user#" + nationID
	if err := tbl.Apply(ctx, rowKey, mut); err != nil {
		return fmt.Errorf("apply: %v", err)
	}

	return nil
}
