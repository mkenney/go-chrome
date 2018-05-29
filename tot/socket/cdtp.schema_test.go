package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	schema "github.com/mkenney/go-chrome/tot/cdtp/schema"
)

func TestSchemaGetDomains(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestSchemaGetDomains")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Schema().GetDomains()
	mockResult := &schema.GetDomainsResult{
		Domains: []*schema.Domain{{
			Name:    "name",
			Version: "version",
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Domains[0].Name != result.Domains[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Domains[0].Name, result.Domains[0].Name)
	}

	resultChan = mockSocket.Schema().GetDomains()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
