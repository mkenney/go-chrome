package socket

import (
	"testing"

	audits "github.com/mkenney/go-chrome/tot/audits"
	network "github.com/mkenney/go-chrome/tot/network"
)

func TestAuditsGetEncodedResponse(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &audits.GetEncodedResponseResult{
		Body:         "Response body",
		OriginalSize: 1,
		EncodedSize:  2,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Audits().GetEncodedResponse(&audits.GetEncodedResponseParams{
		RequestID: network.RequestID("audit-id"),
		Encoding:  audits.Encoding.Webp,
		Quality:   1,
		SizeOnly:  true,
	})
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Audits().GetEncodedResponse(&audits.GetEncodedResponseParams{
		RequestID: network.RequestID("audit-id"),
		Encoding:  audits.Encoding.Webp,
		Quality:   1,
		SizeOnly:  true,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
