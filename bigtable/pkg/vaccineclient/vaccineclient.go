package vaccineclient

import (
	"context"
	"fmt"

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

func (vaccineClient *VaccineClient) CreateUser(nationID string, name string, healthCardID string, gender string, birthday string, address string, phone string, vaccines string) error {
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

	mut.Set(columnFamilyName, "name", timestamp, []byte(name))
	mut.Set(columnFamilyName, "healthCardID", timestamp, []byte(healthCardID))
	mut.Set(columnFamilyName, "gender", timestamp, []byte(gender))
	mut.Set(columnFamilyName, "birthday", timestamp, []byte(birthday))
	mut.Set(columnFamilyName, "address", timestamp, []byte(address))
	mut.Set(columnFamilyName, "phone", timestamp, []byte(phone))
	mut.Set(columnFamilyName, "vaccines", timestamp, []byte(vaccines))

	rowKey := "user#" + nationID
	if err := tbl.Apply(ctx, rowKey, mut); err != nil {
		return fmt.Errorf("apply: %v", err)
	}

	return nil
}
