package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Performance - https://chromedevtools.github.io/devtools-protocol/tot/Performance/
*/
type Performance struct{}

/*
Disable disables collecting and reporting metrics.
*/
func (Performance) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "Performance.disable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables collecting and reporting metrics.
*/
func (Performance) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "Performance.enable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetMetrics retrieves current values of run-time metrics.
*/
func (Overlay) GetMetrics(socket *Socket, params *performance.GetMetricsParams) error {
	command := &protocol.Command{
		method: "Performance.getMetrics",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnMetrics adds a handler to the Performance.metrics event. Performance.metrics returns current
values of the metrics.
*/
func (Overlay) OnMetrics(socket *Socket, callback func(event *performance.MetricsEvent)) error {
	handler := protocol.NewEventHandler(
		"Performance.metrics",
		func(name string, params []byte) {
			event := &performance.MetricsEvent{}
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
