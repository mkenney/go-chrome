package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/security"
)

func TestSecurityDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &security.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Security().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Security().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestSecurityEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &security.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Security().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Security().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestSecurityHandleCertificateError(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &security.HandleCertificateErrorParams{
		EventID: 1,
		Action:  security.CertificateErrorAction.Continue,
	}
	mockResult := &security.HandleCertificateErrorResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Security().HandleCertificateError(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Security().HandleCertificateError(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestSecuritySetIgnoreCertificateErrors(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &security.SetIgnoreCertificateErrorsParams{
		Ignore: true,
	}
	mockResult := &security.SetIgnoreCertificateErrorsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Security().SetIgnoreCertificateErrors(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Security().SetIgnoreCertificateErrors(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestSecuritySetOverrideCertificateErrors(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &security.SetOverrideCertificateErrorsParams{
		Override: true,
	}
	mockResult := &security.SetOverrideCertificateErrorsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Security().SetOverrideCertificateErrors(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Security().SetOverrideCertificateErrors(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestSecurityOnCertificateError(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *security.CertificateErrorEvent)
	soc.Security().OnCertificateError(func(eventData *security.CertificateErrorEvent) {
		resultChan <- eventData
	})

	mockResult := &security.CertificateErrorEvent{
		EventID:    1,
		ErrorType:  "error-type",
		RequestURL: "http://some.url",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Security.certificateError",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.EventID != result.EventID {
		t.Errorf("Expected %d, got %d", mockResult.EventID, result.EventID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Security.certificateError",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestSecurityOnSecurityStateChanged(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *security.StateChangedEvent)
	soc.Security().OnSecurityStateChanged(func(eventData *security.StateChangedEvent) {
		resultChan <- eventData
	})

	mockResult := &security.StateChangedEvent{
		State:                 security.State.Unknown,
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Security.securityStateChanged",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.State != result.State {
		t.Errorf("Expected %s, got %s", mockResult.State, result.State)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Security.securityStateChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
