package chrome

import (
	"fmt"
	"net/url"

	"github.com/mkenney/go-chrome/tot/socket"
	errs "github.com/mkenney/go-errors"
	log "github.com/sirupsen/logrus"
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
		return nil, errs.Wrap(err, "invalid URL")
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
		return nil, errs.Wrap(err, fmt.Sprintf("/new?%s query failed", url.QueryEscape(uri)))
	}

	websocketURL, err := url.Parse(tab.Data().WebSocketDebuggerURL)
	if nil != err {
		return nil, errs.Wrap(err, fmt.Sprintf("invalid websocket URL '%s'", tab.Data().WebSocketDebuggerURL))
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
	chrome   *Chrome
	data     *TabData
	protocol socket.Protocoller
	socket   socket.Socketer
	url      *url.URL
}

/*
Chromium implements Tabber.
*/
func (tab *Tab) Chromium() *Chrome {
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
	log.Debugf("Close result: %v - %v", result, err)
	if nil != err {
		log.Warnf("%s: %s", result, err)
		return nil, errs.Wrap(err, fmt.Sprintf("close/%s query failed", tab.Data().ID))
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
