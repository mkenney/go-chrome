package chrome

import (
	performance "app/chrome/performance"
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
func (Performance) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Performance.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables collecting and reporting metrics.
*/
func (Performance) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Performance.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetMetrics retrieves current values of run-time metrics.
*/
func (Overlay) GetMetrics(
	socket *Socket,
) (performance.GetMetricsResult, error) {
	command := &protocol.Command{
		Method: "Performance.getMetrics",
	}
	socket.SendCommand(command)
	return command.Result.(performance.GetMetricsResult), command.Err
}

/*
OnMetrics adds a handler to the Performance.metrics event. Performance.metrics returns current
values of the metrics.
*/
func (Overlay) OnMetrics(
	socket *Socket,
	callback func(event *performance.MetricsEvent),
) {
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
}
