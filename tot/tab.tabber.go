package chrome

import (
	"fmt"
	"net/url"

	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
	"github.com/mkenney/go-chrome/tot/socket"
)

/*
NewTab spawns a new Tab and returns a reference to it
*/
func (chrome *Chrome) NewTab(uri string) (*Tab, error) {
	var err error

	if "" == uri {
		uri = "about:blank"
	}
	targetURL, err := url.Parse(uri)
	if nil != err {
		return nil, errs.Wrap(err, 0, "invalid URL")
	}

	tab := &Tab{
		chrome: chrome,
		data:   &TabData{},
		url:    targetURL,
	}

	_, err = tab.Chromium().Query(
		fmt.Sprintf("/json/new?%s", url.QueryEscape(uri)),
		url.Values{},
		tab.data,
	)
	if nil != err {
		return nil, errs.Wrap(err, 0, fmt.Sprintf("/new?%s query failed", url.QueryEscape(uri)))
	}

	websocketURL, err := url.Parse(tab.Data().WebSocketDebuggerURL)
	if nil != err {
		return nil, errs.Wrap(err, 0, fmt.Sprintf("invalid websocket URL '%s'", tab.Data().WebSocketDebuggerURL))
	}

	socket := socket.New(websocketURL)
	tab.socket = socket
	tab.protocol = socket
	chrome.tabs = append(chrome.tabs, tab)

	return tab, nil
}

/*
Tab is a struct representing an individual Chrome tab
*/
type Tab struct {
	chrome   Chromium
	data     *TabData
	protocol socket.Protocoller
	socket   socket.Socketer
	url      *url.URL
}

/*
Chromium implements Tabber.
*/
func (tab *Tab) Chromium() Chromium {
	return tab.chrome
}

/*
Close implements Tabber.
*/
func (tab *Tab) Close() (interface{}, error) {
	var err error
	var result interface{}
	tab.Socket().Stop()
	_, err = tab.Chromium().Query(fmt.Sprintf("/json/close/%s", tab.Data().ID), url.Values{}, &result)
	if nil != err {
		log.WithFields(log.Fields{
			"result": result,
		}).Warn("%s", err)
		return nil, errs.Wrap(err, 0, fmt.Sprintf("close/%s query failed", tab.Data().ID))
	}
	tab.Chromium().RemoveTab(tab)
	return result, nil
}

/*
Data implements Tabber.
*/
func (tab *Tab) Data() *TabData {
	return tab.data
}

/*
Protocol implements Tabber.
*/
func (tab *Tab) Protocol() socket.Protocoller {
	return tab.protocol
}

/*
Socket implements Tabber.
*/
func (tab *Tab) Socket() socket.Socketer {
	return tab.socket
}

/*
URL implements Tabber.
*/
func (tab *Tab) URL() *url.URL {
	return tab.url
}
