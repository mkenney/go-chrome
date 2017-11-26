package chrome

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

/*
Tab is a struct representing an individual Chrome tab
*/
type Tab struct {
	Description          string  `json:"description"`
	DevtoolsFrontendURL  string  `json:"devtoolsFrontendUrl"`
	ID                   string  `json:"id"`
	Socket               *Socket `json:"-"`
	Title                string  `json:"title"`
	Type                 string  `json:"type"`
	URL                  string  `json:"url"`
	WebSocketDebuggerURL string  `json:"webSocketDebuggerUrl"`

	loadEventFired bool
}

/*
New spawns a new tab and returns a reference to it
*/
func NewTab(uri string) (*Tab, error) {
	if "" == uri {
		uri = "about:blank"
	}
	var err error
	var browser *Browser

	browser, err = GetBrowser()
	if nil != err {
		return nil, err
	}

	tab := new(Tab)
	_, err = browser.Cmd(fmt.Sprintf("/json/new?%s", url.QueryEscape(uri)), url.Values{}, tab)
	if nil != err {
		return nil, err
	}

	tab.Socket, err = NewSocket(tab)
	if nil != err {
		return nil, err
	}

	//tab.SendCommand("Network.enable", nil)
	tab.SendCommand("Page.enable", nil)
	tab.SendCommand("Emulation.enable", nil)
	//tab.SendCommand("DOM.enable", nil)
	//tab.SendCommand("CSS.enable", nil)

	tab.Socket.AddEventHandler("Page.loadEventFired", func(name string, params []byte) {
		tab.loadEventFired = true
	})
	tab.loadEventFired = false

	browser.Tabs = append(browser.Tabs, tab)
	return tab, nil
}

type tabCommand struct {
	err    error
	result SocketScreenshotResult
	wg     sync.WaitGroup
	prop   interface{}
	method string
}

func (command *tabCommand) Done(result []byte, err error) {
	if err == nil {
		err = json.Unmarshal(result, &command.result)
	}
	command.err = err
	command.wg.Done()
}
func (command *tabCommand) Method() string {
	return command.method
}
func (command *tabCommand) Params() interface{} {
	return command.prop
}
func (command *tabCommand) Run(socket *Socket) error {
	command.wg.Add(1)
	socket.SendCommand(command)
	command.wg.Wait()
	return command.err
}

func (tab *Tab) SendCommand(cmd string, args interface{}) {
	command := &tabCommand{}
	command.method = cmd
	command.prop = args
	command.Run(tab.Socket)
}

/*
Close closes the referenced tab
*/
func (tab *Tab) Close() (interface{}, error) {
	var err error
	var browser *Browser
	var result interface{}

	browser, err = GetBrowser()
	if nil != err {
		return "", err
	}

	_, err = browser.Cmd(fmt.Sprintf("/json/close/%s", tab.ID), url.Values{}, &result)
	log.Debugf("Close result: %s: %s", result, err)
	if nil != err {
		log.Warnf("%s: %s", result, err)
		return "", err
	}

	return result, nil
}

/*
RenderScreenshot takes a screenshot of the referenced tab
*/
//func (tab *Tab) RenderScreenshot(format string, width, height int, maxWait int, handler func(result SocketScreenshotResult)) {
func RenderScreenshots(params url.Values, handle func(results []SocketScreenshotResult)) {
	var takeScreenshot func(tab *Tab) SocketScreenshotResult
	takeScreenshot = func(tab *Tab) SocketScreenshotResult {
		log.Debugf("Screenshot params: %v", params)

		var viewportParams interface{}
		height, _ := strconv.Atoi(params["height"][0])
		scale, _ := strconv.Atoi(params["scale"][0])
		width, _ := strconv.Atoi(params["width"][0])
		x, _ := strconv.Atoi(params["x-offset"][0])
		y, _ := strconv.Atoi(params["y-offset"][0])
		viewportParams = nil
		if height != 0 || width != 0 || x != 0 || y != 0 {
			viewportParams = &socketScreenshotViewport{
				x,
				y,
				width,
				height,
				scale,
			}
		}

		quality := 0
		if len(params["quality"]) > 0 {
			quality, _ = strconv.Atoi(params["quality"][0])
		}

		cmd := &SocketScreenshotCmd{}
		cmd.params = &socketScreenshotParams{
			params["format"][0],
			quality,
			viewportParams,
			false,
		}

		cmd.Run(tab.Socket)

		return cmd.result
	}

	tabs := make([]*Tab, 0)
	errors := make([]error, 0)
	for _, url := range params["url"] {
		tab, err := NewTab(url)
		if nil != err {
			errors = append(errors, err)
		} else {
			tabs = append(tabs, tab)
		}
	}

	results := make([]SocketScreenshotResult, 0)
	start := time.Now()
	timeout, _ := strconv.Atoi(params["timeout"][0])
	for {
		splice := make([]int, 0)
		for k, tab := range tabs {
			if tab.loadEventFired || (timeout > 0 && time.Since(start) > (time.Duration(timeout)*time.Second)) {
				if tab.loadEventFired {
					log.Info("Page loaded, sending screenshot command to socket")
				} else {
					log.Infof("%d second timeout exceeded, sending screenshot command to socket", timeout)
				}
				results = append(results, takeScreenshot(tab))
				splice = append(splice, k)
			}
		}

		for a := len(splice) - 1; a >= 0; a-- {
			tabs = append(tabs[:splice[a]], tabs[splice[a]+1:]...)
		}
		if 0 == len(tabs) {
			break
		}
		time.Sleep(1 * time.Second)
	}

	handle(results)
}

func RenderScreenshotsTest(params url.Values, handle func(results []SocketResult)) {
	takeScreenshot := func(tab *Tab) SocketResult {
		log.Debugf("Screenshot params: %v", params)

		var viewportParams interface{}
		height, _ := strconv.Atoi(params["height"][0])
		scale, _ := strconv.Atoi(params["scale"][0])
		width, _ := strconv.Atoi(params["width"][0])
		x, _ := strconv.Atoi(params["x-offset"][0])
		y, _ := strconv.Atoi(params["y-offset"][0])
		viewportParams = nil
		if height != 0 || width != 0 || x != 0 || y != 0 {
			viewportParams = &socketScreenshotViewport{
				x,
				y,
				width,
				height,
				scale,
			}
		}

		quality := 0
		if len(params["quality"]) > 0 {
			quality, _ = strconv.Atoi(params["quality"][0])
		}

		cmd := NewSocketCmd(
			"Page.captureScreenshot",
			&socketScreenshotParams{
				params["format"][0],
				quality,
				viewportParams,
				false,
			})
		cmd.Run(tab.Socket)

		return cmd.result
	}

	tabs := make([]*Tab, 0)
	errors := make([]error, 0)
	for _, url := range params["url"] {
		tab, err := NewTab(url)
		if nil != err {
			errors = append(errors, err)
		} else {
			tabs = append(tabs, tab)
		}
	}

	results := make([]SocketResult, 0)
	start := time.Now()
	timeout, _ := strconv.Atoi(params["timeout"][0])
	for {
		splice := make([]int, 0)
		for k, tab := range tabs {
			if tab.loadEventFired || (timeout > 0 && time.Since(start) > (time.Duration(timeout)*time.Second)) {
				if tab.loadEventFired {
					log.Info("Page loaded, sending screenshot command to socket")
				} else {
					log.Infof("%d second timeout exceeded, sending screenshot command to socket", timeout)
				}
				log.Debugf("****************************************************")
				log.Debugf("****************************************************")
				log.Debugf("****************************************************")
				results = append(results, takeScreenshot(tab))
				splice = append(splice, k)
			}
		}

		for a := len(splice) - 1; a >= 0; a-- {
			tabs = append(tabs[:splice[a]], tabs[splice[a]+1:]...)
		}
		if 0 == len(tabs) {
			break
		}
		time.Sleep(1 * time.Second)
	}

	handle(results)
}
