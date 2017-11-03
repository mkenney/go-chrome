/*
Package chrome provides an interface to a headless Chrome instance.
*/
package chrome

import (
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
	tab := new(Tab)
	params[uri] = nil
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

func (tab *Tab) Close() {

}
