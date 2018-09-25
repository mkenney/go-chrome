// +build ignore

package main

import (
	"encoding/json"
	"fmt"

	"github.com/mkenney/go-chrome/tot/page"

	chrome "github.com/mkenney/go-chrome/tot"
	"github.com/mkenney/go-chrome/tot/dom"
)

func main() {
	var err error

	// Create a chrome instance
	browser := chrome.New(
		&chrome.Flags{
			"addr":                     "0.0.0.0",
			"remote-debugging-address": "0.0.0.0",
			"remote-debugging-port":    9222,
		}, "", "", "", "",
	)

	// Open a tab and navigate to the URL you want to screenshot.
	tab, err := browser.NewTab("https://www.google.com")
	if nil != err {
		panic(err)
	}

	// Enable Page events for this tab.
	if enableResult := <-tab.Page().Enable(); nil != enableResult.Err {
		panic(enableResult.Err)
	}

	// Enable the DOM agent for this tab.
	if enableResult := <-tab.DOM().Enable(); nil != enableResult.Err {
		panic(enableResult.Err)
	}

	// Create a channel to receive the DOM data.
	results := make(chan *dom.GetDocumentResult)

	// When the page load event fires, deliver the root DOM node.
	tab.Page().OnLoadEventFired(func(event *page.LoadEventFiredEvent) {
		results <- <-tab.DOM().GetDocument(&dom.GetDocumentParams{})
	})

	// Wait for the handler to fire
	result := <-results
	tmp, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("DOM: %s\n\n", string(tmp))
}
