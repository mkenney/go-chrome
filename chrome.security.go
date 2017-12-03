package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Security - https://chromedevtools.github.io/devtools-protocol/tot/Security/
*/
type Security struct{}

/*
Disable disables tracking security state changes.
*/
func (Security) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "Security.disable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable tracking security state changes.
*/
func (Security) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "Security.enable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
HandleCertificateError handles a certificate error that fired a certificateError event.
*/
func (Security) HandleCertificateError(socket *Socket, params *security.HandleCertificateErrorParams) error {
	command := &protocol.Command{
		method: "Security.handleCertificateError",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetOverrideCertificateErrors enables/disables overriding certificate errors. If enabled, all
certificate error events need to be handled by the DevTools client and should be answered with
handleCertificateError commands.
*/
func (Security) SetOverrideCertificateErrors(socket *Socket, params *security.SetOverrideCertificateErrorsParams) error {
	command := &protocol.Command{
		method: "Security.setOverrideCertificateErrors",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnCertificateError adds a handler to the Security.certificateError event. Security.certificateError
fires when there is a certificate error. If overriding certificate errors is enabled, then it should
be handled with the handleCertificateError command. Note: this event does not fire if the
certificate error has been allowed internally.
*/
func (Security) OnCertificateError(socket *Socket, callback func(event *security.CertificateErrorEvent)) error {
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
	return command.Err
}

/*
OnSecurityStateChanged adds a handler to the Security.securityStateChanged event.
Security.securityStateChanged fires when the security state of the page changed.
*/
func (Security) OnSecurityStateChanged(socket *Socket, callback func(event *security.SecurityStateChangedEvent)) error {
	handler := protocol.NewEventHandler(
		"Security.securityStateChanged",
		func(name string, params []byte) {
			event := &security.SecurityStateChangedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}
