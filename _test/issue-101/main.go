package main

import (
	"fmt"
	"time"

	"github.com/bdlm/log"
	"github.com/mkenney/go-chrome/tot"
)

func init() {
	level, _ := log.ParseLevel("debug")
	log.SetLevel(level)
}

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

	// Open a tab
	tab, err := browser.NewTab("")
	if nil != err {
		panic(err)
	}

	// Do nothing so things can settle a bit
	cnt := 0
	for {
		fmt.Printf(".")
		time.Sleep(1 * time.Second)
		cnt++
		if cnt > 5 {
			break
		}
	}

	tab.Socket().Stop()
	fmt.Printf("GOT HERE")

	// Don't exit before the logs are written
	time.Sleep(5 * time.Second)
}
