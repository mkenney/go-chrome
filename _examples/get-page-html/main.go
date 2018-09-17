package main

import (
	"encoding/json"
	"fmt"
	"time"

	chrome "github.com/mkenney/go-chrome/tot"
	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/page"
)

func main() {
	var err error
	outer_html_chan := make(chan *dom.GetOuterHTMLResult)
	var document *dom.GetDocumentResult

	chrome_path := "/usr/bin/google-chrome"
	//chrome_path := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"

	// Define a chrome instance with remote debugging enabled.
	browser := chrome.New(
		// See https://developers.google.com/web/updates/2017/04/headless-chrome#cli
		// for details about startup flags
		&chrome.Flags{
			"addr":               "localhost",
			"disable-extensions": nil,
			"disable-gpu":        nil,
			"headless":           true,
			"hide-scrollbars":    nil,
			"no-first-run":       nil,
			"no-sandbox":         nil,
			"port":               9222,
			"remote-debugging-address": "0.0.0.0",
			"remote-debugging-port":    9222,
		},
		chrome_path, // Path to Chromeium binary
		"/tmp",      // Set the Chromium working directory
		"/dev/null", // Ignore internal Chromium output, set to empty string for os.Stdout
		"/dev/null", // Ignore internal Chromium errors, set to empty string for os.Stderr
	)

	// Start the chrome process.
	if err := browser.Launch(); nil != err {
		panic(err)
	}

	// Open a tab and navigate to the URL you want to screenshot.
	tab, err := browser.NewTab("http://www.brainjar.com/java/host/test.html")
	if nil != err {
		panic(err)
	}

	// When the page load event fires, deliver the root DOM node.
	tab.Page().OnDOMContentEventFired(func(event *page.DOMContentEventFiredEvent) {
		document := <-tab.DOM().GetDocument(&dom.GetDocumentParams{Depth: -1})
		outer_html_chan <- <-tab.DOM().GetOuterHTML(&dom.GetOuterHTMLParams{
			NodeID: document.Root.NodeID,
		})
	})

	// Enable the DOM agent for this tab.
	if enableResult := <-tab.DOM().Enable(); nil != enableResult.Err {
		panic(enableResult.Err)
	}

	var result *dom.GetOuterHTMLResult
	select {
	case result = <-outer_html_chan:
	case <-time.After(2 * time.Second):
		fmt.Println("timeout elapsed, requesting dom")
		document = <-tab.DOM().GetDocument(&dom.GetDocumentParams{Depth: -1})
		result = <-tab.DOM().GetOuterHTML(&dom.GetOuterHTMLParams{
			NodeID: document.Root.NodeID,
		})
	}
	tmp, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s\n\n", string(tmp))
}
