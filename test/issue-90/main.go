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

	// Define a chrome instance with remote debugging enabled.
	browser := chrome.New(
		// See https://developers.google.com/web/updates/2017/04/headless-chrome#cli
		// for details about startup flags
		&chrome.Flags{
			"addr":               "localhost",
			"disable-extensions": nil,
			"disable-gpu":        nil,
			"headless":           nil,
			"hide-scrollbars":    nil,
			"no-first-run":       nil,
			"no-sandbox":         nil,
			"port":               9222,
			"remote-debugging-address": "0.0.0.0",
			"remote-debugging-port":    9222,
		},
		"/usr/bin/google-chrome", // Path to Chromeium binary
		"/tmp",      // Set the Chromium working directory
		"/dev/null", // Ignore internal Chromium output, set to empty string for os.Stdout
		"/dev/null", // Ignore internal Chromium errors, set to empty string for os.Stderr
	)

	// Start the chrome process.
	if err := browser.Launch(); nil != err {
		panic(err)
	}

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
