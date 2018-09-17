package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bdlm/log"
	chrome "github.com/mkenney/go-chrome/tot"
	"github.com/mkenney/go-chrome/tot/emulation"
	"github.com/mkenney/go-chrome/tot/page"
	"github.com/mkenney/go-chrome/tot/socket"
)

func init() {
	level, _ := log.ParseLevel("debug")
	log.SetLevel(level)
}

type testHandler struct{}

func (handler testHandler) Name() string {
	return "Page.setDocumentContent"
}
func (handler testHandler) Handle(response *socket.Response) {
	tmp, _ := json.MarshalIndent(response, "`", "    ")
	log.Debugf("****testHandler*****\n\n%s\n\n****testHandler*****", string(tmp))
	return
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

	// Get the root frame ID.
	ftResult := <-tab.Page().GetFrameTree()
	if nil != ftResult.Err {
		panic(ftResult.Err)
	}

	start := time.Now()
	// Write data to the frame ID.
	setContentResult := <-tab.Page().SetDocumentContent(&page.SetDocumentContentParams{
		FrameID: page.FrameID(ftResult.FrameTree.Frame.ID),
		HTML:    htmlString,
	})
	if nil != setContentResult.Err {
		log.Fatalf("%+v", setContentResult.Err)
	}

	// Set the device emulation parameters, make the page tall enough
	// for this image.
	overrideResult := <-tab.Emulation().SetDeviceMetricsOverride(
		&emulation.SetDeviceMetricsOverrideParams{
			Width:  300,
			Height: 300,
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
			Format: page.Format.Png,
			//Quality: 100,
		},
	)
	if nil != screenshotResult.Err {
		panic(screenshotResult.Err)
	}
	stop := time.Now()
	elapsed := stop.Sub(start)
	log.Debugf("Elapsed time: %s", elapsed)

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
