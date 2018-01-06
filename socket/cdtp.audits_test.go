package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	audits "github.com/mkenney/go-chrome/cdtp/audits"
	network "github.com/mkenney/go-chrome/cdtp/network"
)

func TestAuditsGetEncodedResponse(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Audits().GetEncodedResponse(&audits.GetEncodedResponseParams{
		RequestID: network.RequestID("audit-id"),
		Encoding:  "encoding",
		Quality:   1,
		SizeOnly:  true,
	})
	mockResult := &audits.GetEncodedResponseResult{
		Body:         "Response body",
		OriginalSize: 1,
		EncodedSize:  2,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Audits.getEncodedResponse",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.Body != mockResult.Body {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.Body,
			result.Body,
		)
	}
}
