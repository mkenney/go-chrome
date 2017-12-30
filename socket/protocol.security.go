package socket

import (
	"encoding/json"

	security "github.com/mkenney/go-chrome/protocol/security"
	log "github.com/sirupsen/logrus"
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
func (protocol *SecurityProtocol) Disable() error {
	command := NewCommand("Security.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable tracking security state changes.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-enable
*/
func (protocol *SecurityProtocol) Enable() error {
	command := NewCommand("Security.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetIgnoreCertificateErrors enables/disables whether all certificate errors
should be ignored.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors
EXPERIMENTAL.
*/
func (protocol *SecurityProtocol) SetIgnoreCertificateErrors(
	params *security.SetIgnoreCertificateErrorsParams,
) error {
	command := NewCommand("Security.setIgnoreCertificateErrors", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
HandleCertificateError handles a certificate error that fired a certificateError
event.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-handleCertificateError
*/
func (protocol *SecurityProtocol) HandleCertificateError(
	params *security.HandleCertificateErrorParams,
) error {
	command := NewCommand("Security.handleCertificateError", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetOverrideCertificateErrors enables/disables overriding certificate errors. If
enabled, all certificate error events need to be handled by the DevTools client
and should be answered with HandleCertificateError commands.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setOverrideCertificateErrors
*/
func (protocol *SecurityProtocol) SetOverrideCertificateErrors(
	params *security.SetOverrideCertificateErrorsParams,
) error {
	command := NewCommand("Security.setOverrideCertificateErrors", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
