/*
Package chrome provides an interface to a headless Chrome instance.
*/
package chrome

import (
	"fmt"
	"net/url"
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
}

func NewTab(uri string) (*Tab, error) {
	var err error
	var browser *Browser

	browser, err = GetBrowser()
	if nil != err {
		return nil, err
	}

	params := url.Values{}
	params[uri] = nil

	tab := new(Tab)
	_, err = browser.Cmd("/new", params, tab)
	if nil != err {
		return nil, err
	}

	tab.Socket, err = newSocket(tab.WebSocketDebuggerURL)
	if nil != err {
		return nil, err
	}

	browser.Tabs = append(browser.Tabs, tab)
	return tab, nil
}

func (tab *Tab) Close() (interface{}, error) {
	var err error
	var browser *Browser
	var result interface{}

	browser, err = GetBrowser()
	if nil != err {
		return "", err
	}
	_, err = browser.Cmd(fmt.Sprintf("/close/%s", tab.ID), url.Values{}, &result)
	if nil != err {
		return "", err
	}
	return result, nil
}
