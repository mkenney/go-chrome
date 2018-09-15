package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/device/orientation"
)

func TestDeviceOrientationClearOverride(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &orientation.ClearOverrideResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DeviceOrientation().ClearOverride()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DeviceOrientation().ClearOverride()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDeviceOrientationSetOverride(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &orientation.SetOverrideResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DeviceOrientation().SetOverride(&orientation.SetOverrideParams{
		Alpha: 1,
		Beta:  1,
		Gamma: 1,
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DeviceOrientation().SetOverride(&orientation.SetOverrideParams{
		Alpha: 1,
		Beta:  1,
		Gamma: 1,
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
