package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/security"
)

/*
SecurityProtocol provides a namespace for the Chrome Security protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Security/
*/
type SecurityProtocol struct {
	Socket Socketer
}

/*
Disable disables tracking security state changes.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-disable
*/
func (protocol *SecurityProtocol) Disable() <-chan *security.DisableResult {
	resultChan := make(chan *security.DisableResult)
	command := NewCommand(protocol.Socket, "Security.disable", nil)
	result := &security.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable tracking security state changes.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-enable
*/
func (protocol *SecurityProtocol) Enable() <-chan *security.EnableResult {
	resultChan := make(chan *security.EnableResult)
	command := NewCommand(protocol.Socket, "Security.enable", nil)
	result := &security.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
HandleCertificateError handles a certificate error that fired a certificateError
event.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-handleCertificateError
*/
func (protocol *SecurityProtocol) HandleCertificateError(
	params *security.HandleCertificateErrorParams,
) <-chan *security.HandleCertificateErrorResult {
	resultChan := make(chan *security.HandleCertificateErrorResult)
	command := NewCommand(protocol.Socket, "Security.handleCertificateError", params)
	result := &security.HandleCertificateErrorResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetIgnoreCertificateErrors enables/disables whether all certificate errors
should be ignored.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors
EXPERIMENTAL.
*/
func (protocol *SecurityProtocol) SetIgnoreCertificateErrors(
	params *security.SetIgnoreCertificateErrorsParams,
) <-chan *security.SetIgnoreCertificateErrorsResult {
	resultChan := make(chan *security.SetIgnoreCertificateErrorsResult)
	command := NewCommand(protocol.Socket, "Security.setIgnoreCertificateErrors", params)
	result := &security.SetIgnoreCertificateErrorsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetOverrideCertificateErrors enables/disables overriding certificate errors. If
enabled, all certificate error events need to be handled by the DevTools client
and should be answered with HandleCertificateError commands.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setOverrideCertificateErrors
*/
func (protocol *SecurityProtocol) SetOverrideCertificateErrors(
	params *security.SetOverrideCertificateErrorsParams,
) <-chan *security.SetOverrideCertificateErrorsResult {
	resultChan := make(chan *security.SetOverrideCertificateErrorsResult)
	command := NewCommand(protocol.Socket, "Security.setOverrideCertificateErrors", params)
	result := &security.SetOverrideCertificateErrorsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
OnCertificateError adds a handler to the Security.certificateError event.
Security.certificateError fires when there is a certificate error. If overriding
certificate errors is enabled, then it should be handled with the
HandleCertificateError command. Note: this event does not fire if the
certificate error has been allowed internally.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-certificateError
*/
func (protocol *SecurityProtocol) OnCertificateError(
	callback func(event *security.CertificateErrorEvent),
) {
	handler := NewEventHandler(
		"Security.certificateError",
		func(response *Response) {
			event := &security.CertificateErrorEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnSecurityStateChanged adds a handler to the Security.StateChanged event.
Security.StateChanged fires when the security state of the page changed.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-securityStateChanged
*/
func (protocol *SecurityProtocol) OnSecurityStateChanged(
	callback func(event *security.StateChangedEvent),
) {
	handler := NewEventHandler(
		"Security.securityStateChanged",
		func(response *Response) {
			event := &security.StateChangedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
