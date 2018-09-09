package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/schema"
)

func TestSchemaGetDomains(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &schema.GetDomainsResult{
		Domains: []*schema.Domain{{
			Name:    "name",
			Version: "version",
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Schema().GetDomains()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Domains[0].Name != result.Domains[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Domains[0].Name, result.Domains[0].Name)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Schema().GetDomains()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
