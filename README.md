# go-chrome

[![GoDoc](https://godoc.org/github.com/mkenney/go-chrome?status.svg)](https://godoc.org/github.com/mkenney/go-chrome)
[![Go Report Card](https://goreportcard.com/badge/github.com/mkenney/go-chrome)](https://goreportcard.com/report/github.com/mkenney/go-chrome) [![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://github.com/emersion/stability-badges#experimental)

This package aims to be a complete [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/) implementation. My primary use-case is interacting with [headless Google Chrome](https://developers.google.com/web/updates/2017/04/headless-chrome) in a container environment, but this should be appropriate for developing server side and desktop applications as well.

This is a work in progress. I think most things are workable now but I still need to add tests and happy-path test more of the APIs. I have `Page.captureScreenshot` working fairly well and have been using that for validating changes so far.

This is not production ready by any stretch.
