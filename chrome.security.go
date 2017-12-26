package Chrome

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/protocol"
	security "github.com/mkenney/go-chrome/security"

	log "github.com/sirupsen/logrus"
)

/*
Security - https://chromedevtools.github.io/devtools-protocol/tot/Security/
*/
type Security struct{}

/*
Disable disables tracking security state changes.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-disable
*/
func (Security) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Security.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable tracking security state changes.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-enable
*/
func (Security) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Security.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetIgnoreCertificateErrors enables/disables whether all certificate errors should be ignored.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors
*/
func (Security) SetIgnoreCertificateErrors(
	socket *Socket,
	params *security.SetIgnoreCertificateErrorsParams,
) error {
	command := &protocol.Command{
		Method: "Security.setIgnoreCertificateErrors",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HandleCertificateError handles a certificate error that fired a certificateError event.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-handleCertificateError
*/
func (Security) HandleCertificateError(
	socket *Socket,
	params *security.HandleCertificateErrorParams,
) error {
	command := &protocol.Command{
		Method: "Security.handleCertificateError",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetOverrideCertificateErrors enables/disables overriding certificate errors. If enabled, all
certificate error events need to be handled by the DevTools client and should be answered with
handleCertificateError commands.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setOverrideCertificateErrors
*/
func (Security) SetOverrideCertificateErrors(
	socket *Socket,
	params *security.SetOverrideCertificateErrorsParams,
) error {
	command := &protocol.Command{
		Method: "Security.setOverrideCertificateErrors",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnCertificateError adds a handler to the Security.certificateError event. Security.certificateError
fires when there is a certificate error. If overriding certificate errors is enabled, then it should
be handled with the handleCertificateError command. Note: this event does not fire if the
certificate error has been allowed internally.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-certificateError
*/
func (Security) OnCertificateError(socket *Socket, callback func(event *security.CertificateErrorEvent)) {
	handler := protocol.NewEventHandler(
		"Security.certificateError",
		func(name string, params []byte) {
			event := &security.CertificateErrorEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (Security) OnSecurityStateChanged(socket *Socket, callback func(event *security.StateChangedEvent)) {
	handler := protocol.NewEventHandler(
		"Security.securityStateChanged",
		func(name string, params []byte) {
			event := &security.StateChangedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
