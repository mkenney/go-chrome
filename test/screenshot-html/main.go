package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

	chrome "github.com/mkenney/go-chrome/tot"
	"github.com/mkenney/go-chrome/tot/cdtp/emulation"
	"github.com/mkenney/go-chrome/tot/cdtp/page"
	logfmt "github.com/mkenney/go-log-fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
	level, _ := log.ParseLevel("debug")
	log.SetLevel(level)
	log.SetFormatter(&logfmt.TextFormat{})
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

	// Open a tab and navigate to the URL you want to screenshot.
	tab, err := browser.NewTab("")
	if nil != err {
		panic(err)
	}

	// Enable Page events for this tab.
	if enableResult := <-tab.Page().Enable(); nil != enableResult.Err {
		panic(enableResult.Err)
	}

	ftResult := <-tab.Page().GetFrameTree()
	tmp, _ := json.MarshalIndent(ftResult, "", "    ")
	log.Debugf("Frame Tree: %s", string(tmp))
	log.Debugf("Err: %+v", ftResult.Err)

	setContentResult := <-tab.Page().SetDocumentContent(&page.SetDocumentContentParams{
		FrameID: page.FrameID(ftResult.FrameTree.Frame.ID),
		HTML:    htmlString,
	})
	if nil != setContentResult.Err {
		log.Fatalf("%+v", setContentResult.Err)
	}

	// Set the device emulation parameters.
	overrideResult := <-tab.Emulation().SetDeviceMetricsOverride(
		&emulation.SetDeviceMetricsOverrideParams{
			Width:  1440,
			Height: 1440,
			ScreenOrientation: &emulation.ScreenOrientation{
				Type:  emulation.OrientationType.PortraitPrimary,
				Angle: 90,
			},
		},
	)
	if nil != overrideResult.Err {
		panic(overrideResult.Err)
	}

	// Capture a screenshot of the current state of the current
	// page.
	screenshotResult := <-tab.Page().CaptureScreenshot(
		&page.CaptureScreenshotParams{
			Format:  page.Format.Jpeg,
			Quality: 50,
		},
	)
	if nil != screenshotResult.Err {
		panic(screenshotResult.Err)
	}

	// Decode the base64 encoded image data
	data, err := base64.StdEncoding.DecodeString(screenshotResult.Data)
	if nil != err {
		panic(err)
	}

	// Write the generated image to a file
	err = ioutil.WriteFile("/tmp/test/example.jpg", data, 0644)
	if nil != err {
		panic(err)
	}

	fmt.Println("Finished rendering example.jpg")
}
