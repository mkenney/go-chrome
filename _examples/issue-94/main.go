// +build ignore
package main

import (
	"fmt"
	"os"

	"github.com/bdlm/log"
	"github.com/mkenney/go-chrome/tot"
	"github.com/mkenney/go-chrome/tot/network"
	"github.com/mkenney/go-chrome/tot/page"
	"github.com/sanity-io/litter"
)

func init() {
	level, _ := log.ParseLevel("debug")
	log.SetLevel(level)
}

var flags = &chrome.Flags{
	"remote-debugging-port": 9222,

	// Puppeteer - Headless Mode
	"headless":        nil,
	"disable-gpu":     nil,
	"hide-scrollbars": nil,
	"mute-audio":      nil,

	// Puppeteer - Default Args
	"disable-background-networking":          nil,
	"disable-background-timer-throttling":    nil,
	"disable-client-side-phishing-detection": nil,
	"disable-default-apps":                   nil,
	"disable-dev-shm-usage":                  nil,
	"disable-extensions":                     nil,
	"disable-hang-monitor":                   nil,
	"disable-popup-blocking":                 nil,
	"disable-prompt-on-repost":               nil,
	"disable-sync":                           nil,
	"disable-translate":                      nil,
	"metrics-recording-only":                 nil,
	"no-first-run":                           nil,
	"safebrowsing-disable-auto-update":       nil,

	// Puppeteer - Automation Args
	"enable-automation": nil,
	"password-store":    "basic",
	"use-mock-keychain": nil,

	// From alpeware/chrome-headless-trunk (docker image) /usr/bin/start.sh
	"no-sandbox":    nil,
	"user-data-dir": os.TempDir(),
}

func main() {
	browser := chrome.New(flags, `/usr/bin/google-chrome`, os.TempDir(), "", "")
	if err := browser.Launch(); nil != err {
		panic(err)
	}

	defer func() {
		if err := browser.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	t, err := browser.NewTab("")
	if err != nil {
		panic(err)
	}

	if enableResult := <-t.Page().Enable(); enableResult.Err != nil {
		panic(enableResult.Err)
	}

	if enableResult := <-t.Network().Enable(&network.EnableParams{}); enableResult.Err != nil {
		panic(enableResult.Err)
	}

	ch1 := make(chan bool)
	t.Page().OnFrameNavigated(func(event *page.FrameNavigatedEvent) {
		litter.Dump(event)
		ch1 <- true
	})

	ch2 := make(chan bool)
	t.Network().OnRequestWillBeSent(func(event *network.RequestWillBeSentEvent) {
		litter.Dump(event)
		ch2 <- true
	})

	if res := <-t.Page().Navigate(&page.NavigateParams{URL: "https://www.google.com"}); res.Err != nil {
		panic(res.Err)
	}

	<-ch1
	<-ch2

	browser.Close()
	os.Exit(0)
}
