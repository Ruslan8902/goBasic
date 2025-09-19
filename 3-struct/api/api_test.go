package api_test

import (
	"encoding/json"
	"gobasics/api"
	"gobasics/bins"
	"gobasics/config"
	"gobasics/storage"
	"testing"
)

type CreateResponse struct {
	Record   map[string]any `json:"record"`
	Metadata map[string]any `json:"metadata"`
}

type DeleteResponse struct {
	Metadata map[string]any `json:"metadata"`
	Message  string         `json:"message"`
}

func TestCreateBin(t *testing.T) {
	configStuct := config.NewConfig()
	db := storage.NewStorage("bins2.json")
	binListWithDb := bins.NewBinListWithDb(db)
	filepath := "data.json"
	binName := "test-bin"

	resBody, err := api.CreateBin(configStuct, binListWithDb, &filepath, &binName)

	var createdBinBody CreateResponse
	json.Unmarshal(resBody, &createdBinBody)
	created_id, _ := createdBinBody.Metadata["id"].(string)

	_, _ = api.DeleteBin(configStuct, binListWithDb, &created_id)

	if err != nil {
		t.Errorf("Unexpected %v", err)
	}

	if createdBinBody.Record["login"] != "Rus" {
		t.Errorf("Expected %v got %v", "Rus", createdBinBody.Record["login"])
	}
	if createdBinBody.Record["email"] != "some@mail.ru" {
		t.Errorf("Expected %v got %v", "some@mail.ru", createdBinBody.Record["email"])
	}
}

func TestGetBin(t *testing.T) {
	configStuct := config.NewConfig()
	db := storage.NewStorage("bins2.json")
	binListWithDb := bins.NewBinListWithDb(db)
	filepath := "data.json"
	binName := "test-bin"

	resBody, _ := api.CreateBin(configStuct, binListWithDb, &filepath, &binName)

	var createdBinBody CreateResponse
	json.Unmarshal(resBody, &createdBinBody)
	created_id, _ := createdBinBody.Metadata["id"].(string)

	gerResp, err := api.GetBin(configStuct, &created_id)

	var getBinBody CreateResponse
	json.Unmarshal(gerResp, &getBinBody)

	_, _ = api.DeleteBin(configStuct, binListWithDb, &created_id)

	if err != nil {
		t.Errorf("Unexpected %v", err)
	}
	if getBinBody.Record["login"] != "Rus" {
		t.Errorf("Expected %v got %v", "Rus", getBinBody.Record["login"])
	}
	if getBinBody.Record["email"] != "some@mail.ru" {
		t.Errorf("Expected %v got %v", "some@mail.ru", getBinBody.Record["email"])
	}
}

func TestDeleteBin(t *testing.T) {
	configStuct := config.NewConfig()
	db := storage.NewStorage("bins2.json")
	binListWithDb := bins.NewBinListWithDb(db)
	filepath := "data.json"
	binName := "test-bin"

	resBody, _ := api.CreateBin(configStuct, binListWithDb, &filepath, &binName)

	var createdBinBody CreateResponse
	json.Unmarshal(resBody, &createdBinBody)
	created_id, _ := createdBinBody.Metadata["id"].(string)

	delResp, err := api.DeleteBin(configStuct, binListWithDb, &created_id)

	var deleteBinBody DeleteResponse
	json.Unmarshal(delResp, &deleteBinBody)

	if err != nil {
		t.Errorf("Unexpected %v", err)
	}
	if deleteBinBody.Message != "Bin deleted successfully" {
		t.Errorf("Expected %v got %v", "Bin deleted successfully", deleteBinBody)
	}
	if deleteBinBody.Metadata["id"] != created_id {
		t.Errorf("Expected %v got %v, wrong bin is deleted", created_id, deleteBinBody.Metadata["id"])
	}
}

func TestUpdateBin(t *testing.T) {
	configStuct := config.NewConfig()
	db := storage.NewStorage("bins2.json")
	binListWithDb := bins.NewBinListWithDb(db)
	filepath := "data.json"
	binName := "test-bin"

	resBody, _ := api.CreateBin(configStuct, binListWithDb, &filepath, &binName)

	var createdBinBody CreateResponse
	json.Unmarshal(resBody, &createdBinBody)
	created_id, _ := createdBinBody.Metadata["id"].(string)

	resBody2, err := api.UpdateBin(configStuct, binListWithDb, &created_id, &filepath)
	var updatedBinBody CreateResponse
	json.Unmarshal(resBody2, &updatedBinBody)
	t.Log(string(resBody2))
	_, _ = api.DeleteBin(configStuct, binListWithDb, &created_id)

	if err != nil {
		t.Errorf("Unexpected %v", err)
	}
	if updatedBinBody.Record["login"] != "Rus" {
		t.Errorf("Expected %v got %v", "Rus", updatedBinBody.Record["login"])
	}
	if updatedBinBody.Record["email"] != "some@mail.ru" {
		t.Errorf("Expected %v got %v", "some@mail.ru", updatedBinBody.Record["email"])
	}
}
