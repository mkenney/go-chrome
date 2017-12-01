package chrome

import (
	"encoding/json"
	"sync"
)

/*
SocketScreenshotCmd is a representation of the CaptureScreenshot command
*/
type SocketScreenshotCmd struct {
	err    error
	params *socketScreenshotParams
	result SocketScreenshotResult
	wg     sync.WaitGroup
}
type socketScreenshotParams struct {
	Format      string      `json:"format,omitempty"`
	Quality     int         `json:"quality,omitempty"`
	Clip        interface{} `json:"clip,omitempty"`
	FromSurface bool        `json:"fromSurface,omitempty"`
}
type socketScreenshotViewport struct {
	XOffset int `json:"x"`
	YOffset int `json:"y"`
	Width   int `json:"width"`
	Height  int `json:"height"`
	Scale   int `json:"scale"`
}

/*
SocketScreenshotResult represents the result of the screenshot capture
*/
type SocketScreenshotResult struct {
	Data string `json:"data"`
}

/*
Done is a socketCmdIface interface implementation
*/
func (cmd *SocketScreenshotCmd) Done(result []byte, err error) {
	if err == nil {
		err = json.Unmarshal(result, &cmd.result)
	}
	cmd.err = err
	cmd.wg.Done()
}

/*
Method is a socketCmdIface interface implementation
*/
func (cmd *SocketScreenshotCmd) Method() string {
	return "Page.captureScreenshot"
}

/*
Params is a socketCmdIface interface implementation
*/
func (cmd *SocketScreenshotCmd) Params() (params interface{}) {
	return cmd.params
}

/*
Run is a socketCmdIface interface implementation
*/
func (cmd *SocketScreenshotCmd) Run(socket *Socket) error {
	cmd.wg.Add(1)
	socket.SendCommand(cmd)
	cmd.wg.Wait()
	return cmd.err
}

/*
ScreenshotCommand tells the socket to capture a screenshot of the current page
*/
func ScreenshotCommand(socket *Socket, handler func(result interface{})) {
	OnLoadEvent(socket, func(evt *LoadEvent) {
		cmd := &SocketScreenshotCmd{}
		cmd.Run(socket)
		handler(cmd.result)
	})
}
