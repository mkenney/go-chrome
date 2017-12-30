package socket

import (
	"encoding/json"

	performance "github.com/mkenney/go-chrome/protocol/performance"
	log "github.com/sirupsen/logrus"
)

/*
PerformanceProtocol provides a namespace for the Chrome Performance protocol
methods.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/
*/
type PerformanceProtocol struct {
	Socket Socketer
}

/*
Disable disables collecting and reporting metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-disable
*/
func (protocol *PerformanceProtocol) Disable() error {
	command := NewCommand("Performance.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables collecting and reporting metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-enable
*/
func (protocol *PerformanceProtocol) Enable() error {
	command := NewCommand("Performance.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetMetrics retrieves current values of run-time metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-getMetrics
*/
func (protocol *PerformanceProtocol) GetMetrics() (*performance.GetMetricsResult, error) {
	command := NewCommand("Performance.getMetrics", nil)
	result := &performance.GetMetricsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
OnMetrics adds a handler to the Performance.metrics event. Performance.metrics
returns current values of the metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#event-metrics
*/
func (protocol *PerformanceProtocol) OnMetrics(
	callback func(event *performance.MetricsEvent),
) {
	handler := NewEventHandler(
		"Performance.metrics",
		func(response *Response) {
			event := &performance.MetricsEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
