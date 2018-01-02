package socket

import (
	"net/url"
	"testing"

	audits "github.com/mkenney/go-chrome/protocol/audits"
	network "github.com/mkenney/go-chrome/protocol/network"
	log "github.com/sirupsen/logrus"
)

func init() {
	level, err := log.ParseLevel("debug")
	if nil == err {
		log.SetLevel(level)
	}
}

func TestAuditsGetEncodedResponse(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockResult := &audits.GetEncodedResponseResult{
		Body:         "Response body",
		OriginalSize: 1,
		EncodedSize:  2,
	}
	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"Audits.GetEncodedResponse",
		mockResult,
	)
	go func() {
		result, err := mockSocket.Audits().GetEncodedResponse(&audits.GetEncodedResponseParams{
			RequestID: network.RequestID("audit-id"),
			Encoding:  "encoding",
			Quality:   1,
			SizeOnly:  true,
		})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
		if result.Body != mockResult.Body {
			t.Errorf(
				"Expected %s, received %s",
				mockResult.Body,
				result.Body,
			)
		}
	}()
}
