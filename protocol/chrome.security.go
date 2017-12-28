package protocol

import (
	"encoding/json"

	security "github.com/mkenney/go-chrome/protocol/security"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Security is a struct that provides a namespace for the Chrome Security protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Security/
*/
type Security struct{}

/*
Disable disables tracking security state changes.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-disable
*/
func (Security) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Security.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable tracking security state changes.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-enable
*/
func (Security) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Security.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetIgnoreCertificateErrors enables/disables whether all certificate errors should be ignored.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors
*/
func (Security) SetIgnoreCertificateErrors(
	socket sock.Socketer,
	params *security.SetIgnoreCertificateErrorsParams,
) error {
	command := sock.NewCommand("Security.setIgnoreCertificateErrors", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
HandleCertificateError handles a certificate error that fired a certificateError event.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-handleCertificateError
*/
func (Security) HandleCertificateError(
	socket sock.Socketer,
	params *security.HandleCertificateErrorParams,
) error {
	command := sock.NewCommand("Security.handleCertificateError", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetOverrideCertificateErrors enables/disables overriding certificate errors. If enabled, all
certificate error events need to be handled by the DevTools client and should be answered with
handleCertificateError commands.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setOverrideCertificateErrors
*/
func (Security) SetOverrideCertificateErrors(
	socket sock.Socketer,
	params *security.SetOverrideCertificateErrorsParams,
) error {
	command := sock.NewCommand("Security.setOverrideCertificateErrors", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnCertificateError adds a handler to the Security.certificateError event. Security.certificateError
fires when there is a certificate error. If overriding certificate errors is enabled, then it should
be handled with the handleCertificateError command. Note: this event does not fire if the
certificate error has been allowed internally.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-certificateError
*/
func (Security) OnCertificateError(
	socket sock.Socketer,
	callback func(event *security.CertificateErrorEvent),
) {
	handler := sock.NewEventHandler(
		"Security.certificateError",
		func(response *sock.Response) {
			event := &security.CertificateErrorEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnSecurityStateChanged adds a handler to the Security.StateChanged event. Security.StateChanged
fires when the security state of the page changed.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-securityStateChanged
*/
func (Security) OnSecurityStateChanged(
	socket sock.Socketer,
	callback func(event *security.StateChangedEvent),
) {
	handler := sock.NewEventHandler(
		"Security.securityStateChanged",
		func(response *sock.Response) {
			event := &security.StateChangedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
