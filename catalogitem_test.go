package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
	"reflect"
	"testing"
	"time"
)

func TestCatalogItem_Create(t *testing.T) {
	key := "test api key"

	mockResponseId := "example"
	mockResponse := new(invdendpoint.CatalogItem)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "delivery"
	mockResponse.Type = "service"

	server, err := invdmockserver.New(200, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	requestEntity := conn.NewCatalogItem()

	requestEntity.Id = "example"
	requestEntity.Name = "delivery"
	requestEntity.Type = "service"

	requestEntity, err = requestEntity.Create(requestEntity)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(requestEntity.CatalogItem, mockResponse) {
		t.Fatal("entity was not created", requestEntity.CatalogItem, mockResponse)
	}

}

func TestCatalogItem_Save(t *testing.T) {
	key := "test api key"

	mockResponseId := "delivery"
	mockResponse := new(invdendpoint.CatalogItem)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"
	mockResponse.Type = "service"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)

	entityToUpdate := conn.NewCatalogItem()

	entityToUpdate.Name = "new-name"

	err = entityToUpdate.Save()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.CatalogItem) {
		t.Fatal("Error: entity not updated correctly")
	}

}

func TestCatalogItem_Delete(t *testing.T) {

	key := "api key"

	mockResponse := ""
	mockResponseId := "example"

	server, err := invdmockserver.New(204, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewCatalogItem()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error Occured Deleting Transaction")
	}

}

func TestCatalogItem_Retrieve(t *testing.T) {

	key := "test api key"

	mockResponseId := "example"
	mockResponse := new(invdendpoint.CatalogItem)
	mockResponse.Id = mockResponseId
	mockResponse.Name = "delivery"
	mockResponse.Type = "service"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewCatalogItem()

	retrievedTransaction, err := entity.Retrieve(mockResponseId)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(retrievedTransaction.CatalogItem, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}

func TestCatalogItem_ListAll(t *testing.T) {

	key := "test api key"

	var mockListResponse [1] invdendpoint.CatalogItem

	mockResponse := new(invdendpoint.CatalogItem)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewCatalogItem()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(result[0].CatalogItem, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}