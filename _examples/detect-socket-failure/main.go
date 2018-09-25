package main

import (
	"fmt"

	errs "github.com/bdlm/errors"
	"github.com/mkenney/go-chrome/codes"
	chrome "github.com/mkenney/go-chrome/tot"
)

func main() {
	browser := chrome.New(
		&chrome.Flags{
			"addr":                     "0.0.0.0",
			"remote-debugging-address": "0.0.0.0",
			"remote-debugging-port":    9222,
		}, "", "", "", "",
	)

	tab, err := browser.NewTab("https://www.google.com")
	if nil != err {
		fmt.Printf("%+v\n", err)
	}

	// Enable Page events for this tab.
	if enableResult := <-tab.Page().Enable(); nil != enableResult.Err {
		fmt.Printf("%+v\n", enableResult.Err)
	}

	// While this is waiting, stop chrome
	for err := range tab.Socket().Errors() {
		fmt.Printf("%+v\n", err)
		if e, ok := err.(errs.Err); ok {
			if codes.SocketPanic == e.Code() {
				return
			}
		}
	}
}
