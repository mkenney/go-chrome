# go-chrome

[![Build status](https://travis-ci.org/mkenney/go-chrome.svg?branch=master)](https://travis-ci.org/mkenney/go-chrome) [![stability-unstable](https://img.shields.io/badge/stability-unstable-yellow.svg)](https://github.com/mkenney/stability-badges#unstable) [![Github issues](https://img.shields.io/github/issues-raw/mkenney/go-chrome.svg)](https://github.com/mkenney/go-chrome/issues) [![Go Report Card](https://goreportcard.com/badge/github.com/mkenney/go-chrome)](https://goreportcard.com/report/github.com/mkenney/go-chrome) [![GoDoc](https://godoc.org/github.com/mkenney/go-chrome?status.svg)](https://godoc.org/github.com/mkenney/go-chrome)


This package aims to be a complete [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/) implementation. My primary use-case is interacting with [headless Google Chrome](https://developers.google.com/web/updates/2017/04/headless-chrome) in a container environment, but this should be appropriate for developing server side and desktop applications as well.

This is a work in progress. I think most things are workable now but I still need to add tests and happy-path test more of the APIs. I have [`Page.captureScreenshot`](https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot) working fairly well and have been using that for validating changes so far. This implementation is based on the [Tip-of-Tree](https://chromedevtools.github.io/devtools-protocol/tot/) documentation and may be prone to change. At some point I will implement stable versions as well.

I added a few small examples to the [wiki](https://github.com/mkenney/go-chrome/wiki).

The API should be stable but is not yet production ready.

# TODO

Add more tests. So many tests... If you would like to contribute but aren't sure how, there are several [open issues](https://github.com/mkenney/go-chrome/issues?q=is%3Aopen+is%3Aissue+project%3Amkenney%2Fgo-chrome%2F1) for tests that need to be written, I'll try to keep that list up to date. There are also [several examples](https://github.com/mkenney/go-chrome/blob/master/socket/protocol.animation_test.go) of tests that have already be written.

Any contributions to tests or otherwise are very welcome!
