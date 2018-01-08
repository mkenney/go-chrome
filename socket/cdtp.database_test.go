package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	database "github.com/mkenney/go-chrome/cdtp/database"
)

func TestDatabaseDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Database().Disable()
	mockResult := &database.DisableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Database.disable",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Database().Disable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Database.disable",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDatabaseEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Database().Enable()
	mockResult := &database.EnableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Database.enable",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Database().Enable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Database.enable",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDatabaseExecuteSQL(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Database().ExecuteSQL(&database.ExecuteSQLParams{
		ID:    database.ID("db-id"),
		Query: "SELECT * FROM table_name",
	})
	mockResult := &database.ExecuteSQLResult{
		ColumnNames: []string{"column1", "column2"},
		Values:      []interface{}{"value1", 2},
		SQLError: &database.Error{
			Code:    1,
			Message: "error message",
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Database.executeSQL",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ColumnNames[0] != result.ColumnNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.ColumnNames[0], result.ColumnNames[0])
	}

	resultChan = mockSocket.Database().ExecuteSQL(&database.ExecuteSQLParams{
		ID:    database.ID("db-id"),
		Query: "SELECT * FROM table_name",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Database.executeSQL",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDatabaseGetTableNames(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Database().GetTableNames(&database.GetTableNamesParams{
		ID: database.ID("db-id"),
	})
	mockResult := &database.GetTableNamesResult{
		TableNames: []string{"table1", "table2"},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Database.getTableNames",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.TableNames[0] != result.TableNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.TableNames[0], result.TableNames[0])
	}

	resultChan = mockSocket.Database().GetTableNames(&database.GetTableNamesParams{
		ID: database.ID("db-id"),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Database.getTableNames",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDatabaseOnAdd(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *database.AddEvent)
	mockSocket.Database().OnAdd(func(eventData *database.AddEvent) {
		resultChan <- eventData
	})
	mockResult := &database.AddEvent{
		Database: &database.Database{
			ID:      database.ID("db-id"),
			Domain:  "database.domain",
			Name:    "database_name",
			Version: "v1.0",
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Database.addDatabase",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *database.AddEvent)
	mockSocket.Database().OnAdd(func(eventData *database.AddEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Database.addDatabase",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
