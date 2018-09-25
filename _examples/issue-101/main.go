// +build ignore

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

	// Create a chrome instance
	browser := chrome.New(
		&chrome.Flags{
			"addr":                     "0.0.0.0",
			"remote-debugging-address": "0.0.0.0",
			"remote-debugging-port":    9222,
		}, "", "", "", "",
	)

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
