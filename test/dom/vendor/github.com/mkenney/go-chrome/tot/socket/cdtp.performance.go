package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/performance"
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
func (protocol *PerformanceProtocol) Disable() <-chan *performance.DisableResult {
	resultChan := make(chan *performance.DisableResult)
	command := NewCommand(protocol.Socket, "Performance.disable", nil)
	result := &performance.DisableResult{}

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
Enable enables collecting and reporting metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-enable
*/
func (protocol *PerformanceProtocol) Enable() <-chan *performance.EnableResult {
	resultChan := make(chan *performance.EnableResult)
	command := NewCommand(protocol.Socket, "Performance.enable", nil)
	result := &performance.EnableResult{}

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
GetMetrics retrieves current values of run-time metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-getMetrics
*/
func (protocol *PerformanceProtocol) GetMetrics() <-chan *performance.GetMetricsResult {
	resultChan := make(chan *performance.GetMetricsResult)
	command := NewCommand(protocol.Socket, "Performance.getMetrics", nil)
	result := &performance.GetMetricsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
