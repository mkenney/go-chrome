package protocol

import (
	"encoding/json"

	performance "github.com/mkenney/go-chrome/protocol/performance"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Performance provides a namespace for the Chrome Performance protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/
*/
var Performance = PerformanceProtocol{}

/*
PerformanceProtocol defines the Performance protocol methods.
*/
type PerformanceProtocol struct{}

/*
Disable disables collecting and reporting metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-disable
*/
func (PerformanceProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Performance.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables collecting and reporting metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-enable
*/
func (PerformanceProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Performance.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetMetrics retrieves current values of run-time metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-getMetrics
*/
func (PerformanceProtocol) GetMetrics(
	socket sock.Socketer,
) (*performance.GetMetricsResult, error) {
	command := sock.NewCommand("Performance.getMetrics", nil)
	result := &performance.GetMetricsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
OnMetrics adds a handler to the Performance.metrics event. Performance.metrics returns current
values of the metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#event-metrics
*/
func (PerformanceProtocol) OnMetrics(
	socket sock.Socketer,
	callback func(event *performance.MetricsEvent),
) {
	handler := sock.NewEventHandler(
		"Performance.metrics",
		func(response *sock.Response) {
			event := &performance.MetricsEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
