package chrome

import (
	"fmt"
	"net/url"

	chrome_error "github.com/mkenney/go-chrome/error"
	"github.com/mkenney/go-chrome/socket"
	log "github.com/sirupsen/logrus"
)

/*
NewTab spawns a new tab and returns a reference to it
*/
func (browser *Browser) NewTab(uri string) (*Tab, *chrome_error.Error) {
	if "" == uri {
		uri = "about:blank"
	}

	tab := &Tab{
		browser: browser,
		data:    &TabData{},
	}

	_, err := tab.Browser().Cmd(fmt.Sprintf("/json/new?%s", url.QueryEscape(uri)), url.Values{}, tab.data)
	if nil != err {
		return nil, chrome_error.NewFromErr(err)
	}

	tab.socket, err = socket.New(tab.Data().WebSocketDebuggerURL)
	if nil != err {
		return nil, err
	}
	browser.tabs = append(browser.tabs, tab)

	return tab, nil
}

/*
Tab is a struct representing an individual Chrome tab
*/
type Tab struct {
	browser *Browser
	data    *TabData
	socket  socket.Socketer
}

type TabData struct {
	Description          string `json:"description"`
	DevtoolsFrontendURL  string `json:"devtoolsFrontendURL"`
	ID                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	URL                  string `json:"url"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerURL"`
}

/*
Browser implements Tabber.
*/
func (tab *Tab) Browser() *Browser {
	return tab.browser
}

/*
Close implements Tabber.
*/
func (tab *Tab) Close() (interface{}, *chrome_error.Error) {
	var err error
	var result interface{}

	_, err = tab.Browser().Cmd(fmt.Sprintf("/json/close/%s", tab.Data().ID), url.Values{}, &result)
	log.Debugf("Close result: %s - %s", result, err)
	if nil != err {
		log.Warnf("%s: %s", result, err)
		return nil, chrome_error.NewFromErr(err)
	}

	return result, nil
}

/*
Data implements Tabber.
*/
func (tab *Tab) Data() *TabData {
	return tab.data
}

/*
Socket implements Tabber.
*/
func (tab *Tab) Socket() socket.Socketer {
	return tab.socket
}
