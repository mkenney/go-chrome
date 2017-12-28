package chrome

import (
	"fmt"
	"net/url"

	"github.com/mkenney/go-chrome/socket"
	log "github.com/sirupsen/logrus"
)

/*
Tab is a struct representing an individual Chrome tab
*/
type Tab struct {
	Chrome               *Chrome         `json:"-"`
	Description          string          `json:"description"`
	DevtoolsFrontendURL  string          `json:"devtoolsFrontendUrl"`
	ID                   string          `json:"id"`
	Socket               socket.Socketer `json:"-"`
	Title                string          `json:"title"`
	Type                 string          `json:"type"`
	URL                  string          `json:"url"`
	WebSocketDebuggerURL string          `json:"webSocketDebuggerUrl"`
}

/*
NewTab spawns a new tab and returns a reference to it
*/
func (chrome *Chrome) NewTab(uri string) (*Tab, error) {
	if "" == uri {
		uri = "about:blank"
	}
	var err error

	tab := &Tab{}
	tab.Chrome = chrome
	_, err = chrome.Cmd(fmt.Sprintf("/json/new?%s", url.QueryEscape(uri)), url.Values{}, tab)
	if nil != err {
		return nil, err
	}

	tab.Socket, err = socket.New(tab.WebSocketDebuggerURL)
	if nil != err {
		return nil, err
	}

	chrome.Tabs = append(chrome.Tabs, tab)
	return tab, nil
}

/*
Close closes the referenced tab
*/
func (tab *Tab) Close() (interface{}, error) {
	var err error
	var result interface{}

	if nil != err {
		return "", err
	}

	_, err = tab.Chrome.Cmd(fmt.Sprintf("/json/close/%s", tab.ID), url.Values{}, &result)
	log.Debugf("Close result: %s: %s", result, err)
	if nil != err {
		log.Warnf("%s: %s", result, err)
		return "", err
	}

	return result, nil
}
