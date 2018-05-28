package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	security "github.com/mkenney/go-chrome/tot/cdtp/security"
)

func TestSecurityDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Security().Disable()
	mockResult := &security.DisableResult{}
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

	resultChan = mockSocket.Security().Disable()
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

func TestSecurityEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Security().Enable()
	mockResult := &security.EnableResult{}
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

	resultChan = mockSocket.Security().Enable()
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

func TestSecurityHandleCertificateError(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &security.HandleCertificateErrorParams{
		EventID: 1,
		Action:  security.CertificateErrorAction.Continue,
	}
	resultChan := mockSocket.Security().HandleCertificateError(params)
	mockResult := &security.HandleCertificateErrorResult{}
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

	resultChan = mockSocket.Security().HandleCertificateError(params)
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

func TestSecuritySetIgnoreCertificateErrors(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &security.SetIgnoreCertificateErrorsParams{
		Ignore: true,
	}
	resultChan := mockSocket.Security().SetIgnoreCertificateErrors(params)
	mockResult := &security.SetIgnoreCertificateErrorsResult{}
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

	resultChan = mockSocket.Security().SetIgnoreCertificateErrors(params)
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

func TestSecuritySetOverrideCertificateErrors(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &security.SetOverrideCertificateErrorsParams{
		Override: true,
	}
	resultChan := mockSocket.Security().SetOverrideCertificateErrors(params)
	mockResult := &security.SetOverrideCertificateErrorsResult{}
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

	resultChan = mockSocket.Security().SetOverrideCertificateErrors(params)
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

func TestSecurityOnCertificateError(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *security.CertificateErrorEvent)
	mockSocket.Security().OnCertificateError(func(eventData *security.CertificateErrorEvent) {
		resultChan <- eventData
	})
	mockResult := &security.CertificateErrorEvent{
		EventID:    1,
		ErrorType:  "error-type",
		RequestURL: "http://some.url",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Security.certificateError",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.EventID != result.EventID {
		t.Errorf("Expected %d, got %d", mockResult.EventID, result.EventID)
	}

	resultChan = make(chan *security.CertificateErrorEvent)
	mockSocket.Security().OnCertificateError(func(eventData *security.CertificateErrorEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Security.certificateError",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestSecurityOnSecurityStateChanged(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *security.StateChangedEvent)
	mockSocket.Security().OnSecurityStateChanged(func(eventData *security.StateChangedEvent) {
		resultChan <- eventData
	})
	mockResult := &security.StateChangedEvent{
		State: security.State.Unknown,
		SchemeIsCryptographic: true,
		Explanations: []*security.StateExplanation{{
			State:            security.State.Unknown,
			Summary:          "summary",
			Description:      "description",
			MixedContentType: security.MixedContentType.Blockable,
			Certificate:      []string{"string1", "string2"},
		}},
		InsecureContentStatus: &security.InsecureContentStatus{
			RanMixedContent:                true,
			DisplayedMixedContent:          true,
			ContainedMixedForm:             true,
			RanContentWithCertErrors:       true,
			DisplayedContentWithCertErrors: true,
			RanInsecureContentStyle:        security.State.Unknown,
			DisplayedInsecureContentStyle:  security.State.Unknown,
		},
		Summary: "summary",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Security.securityStateChanged",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.State != result.State {
		t.Errorf("Expected %s, got %s", mockResult.State, result.State)
	}

	resultChan = make(chan *security.StateChangedEvent)
	mockSocket.Security().OnSecurityStateChanged(func(eventData *security.StateChangedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Security.securityStateChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
