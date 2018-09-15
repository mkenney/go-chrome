package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/database"
)

func TestDatabaseDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &database.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Database().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Database().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDatabaseEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &database.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Database().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Database().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDatabaseExecuteSQL(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &database.ExecuteSQLResult{
		ColumnNames: []string{"column1", "column2"},
		Values:      []interface{}{"value1", 2},
		SQLError: &database.Error{
			Code:    1,
			Message: "error message",
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Database().ExecuteSQL(&database.ExecuteSQLParams{
		ID:    database.ID("db-id"),
		Query: "SELECT * FROM table_name",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ColumnNames[0] != result.ColumnNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.ColumnNames[0], result.ColumnNames[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Database().ExecuteSQL(&database.ExecuteSQLParams{
		ID:    database.ID("db-id"),
		Query: "SELECT * FROM table_name",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDatabaseGetTableNames(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &database.GetTableNamesResult{
		TableNames: []string{"table1", "table2"},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Database().GetTableNames(&database.GetTableNamesParams{
		ID: database.ID("db-id"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.TableNames[0] != result.TableNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.TableNames[0], result.TableNames[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Database().GetTableNames(&database.GetTableNamesParams{
		ID: database.ID("db-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDatabaseOnAdd(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *database.AddEvent)
	soc.Database().OnAdd(func(eventData *database.AddEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Database.addDatabase",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Database.addDatabase",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
