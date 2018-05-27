<p align="center">
    <a href="https://gopherize.me/gopher/255e20ee48c85f3b4701446e2513c100f22129f3"><img src="https://github.com/mkenney/go-chrome/wiki/assets/images/gopher-logo.png" width="300px"></a>
</p>

# go-chrome

[![MIT license](https://img.shields.io/github/license/mkenney/go-chrome.svg)](https://github.com/mkenney/go-chrome/blob/master/LICENSE) [![Beta](https://img.shields.io/badge/stability-beta-33bbff.svg)](https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#beta) [![Build status](https://travis-ci.org/mkenney/go-chrome.svg?branch=master)](https://travis-ci.org/mkenney/go-chrome) [![Issues](https://img.shields.io/github/issues-raw/mkenney/go-chrome.svg)](https://github.com/mkenney/go-chrome/issues) [![Pull requests](https://img.shields.io/github/issues-pr/mkenney/go-chrome.svg)](https://github.com/mkenney/go-chrome/pulls) [![Test coverage](https://img.shields.io/codecov/c/github/mkenney/go-chrome/master.svg)](https://codecov.io/gh/mkenney/go-chrome) [![Report card](https://goreportcard.com/badge/github.com/mkenney/go-chrome)](https://goreportcard.com/report/github.com/mkenney/go-chrome) [![Documentation](https://godoc.org/github.com/mkenney/go-chrome?status.svg)](https://godoc.org/github.com/mkenney/go-chrome)

This package aims to be a complete [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/) implementation. The primary use-case behind this project is interacting with [headless Google Chrome](https://developers.google.com/web/updates/2017/04/headless-chrome) in a container environment, but it should be appropriate for developing server side and desktop applications in any browser that supports the devtools protocol.

The API is fairly settled and basic code-coverage tests have been implemented but real-world testing is needed. [`Page.captureScreenshot`](https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot) and related calls are working well and are regularly used for validating the workability of code refactoring.

This implementation is based on the [Tip-of-Tree](https://chromedevtools.github.io/devtools-protocol/tot/) documentation and may be prone to change. At some point stable versions will be implemented as well, hopefully beginning with v1.3.

# Examples

There are a few small examples of how to use the framework API on the [wiki](https://github.com/mkenney/go-chrome/wiki).

# TODO

Any contributions of any kind are very welcome!

* Refactoring to implement standard interfaces where applicable and review current use of interfaces in the API. Some aren't needed at all and others are used to support test mocks.
* Add more tests, especially for error cases. If you would like to contribute but aren't sure how, take a look at [codecov](https://codecov.io/gh/mkenney/go-chrome) for any tests that could be written. There are [many](https://github.com/mkenney/go-chrome/blob/master/tot/socket/cdtp.animation_test.go) [examples](https://github.com/mkenney/go-chrome/blob/master/tot/cdtp/animation/enum.animation.type_test.go) of tests in the repo.
* Add integration test scripts to the `test/` directory to exercise various functions. The [screenshot script](https://github.com/mkenney/go-chrome/wiki/Example%3A-Capture-A-Screenshot) is already setup there.

# Change Log

## 2018-03-20

Removed the unnecessary array notation for the flags passed to the Chrome process. This is a breaking change but it's simple to update client code and easier to use in general.

When defining your Chrome instance, remove any `[]interface{}{}` declarations in the `Flags` parameter. To demonstrate, the [example code](https://github.com/mkenney/go-chrome/wiki/Example%3A-Capture-A-Screenshot) has changed from:
```go
	// Define a chrome instance with remote debugging enabled.
	browser := chrome.New(
		&chrome.Flags{
			"addr":               []interface{}{"localhost"},
			"disable-extensions": nil,
			"disable-gpu":        nil,
			"headless":           nil,
			"hide-scrollbars":    nil,
			"no-first-run":       nil,
			"no-sandbox":         nil,
			"port":               []interface{}{9222},
			"remote-debugging-address": []interface{}{"0.0.0.0"},
			"remote-debugging-port":    []interface{}{9222},
		},
		"/usr/bin/google-chrome",
		"",
		"",
		"",
	)
```
to simply:
```go
	// Define a chrome instance with remote debugging enabled.
	browser := chrome.New(
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
		"/usr/bin/google-chrome",
		"",
		"",
		"",
	)
```

It's easier to deal with, simpler to understand, and supporting multiple values wasn't ever useful. An [argument could be made](https://github.com/mkenney/go-chrome/issues/new) for supporting something like [`pkg/flag`](https://golang.org/pkg/flag/) I suppose but I don't generally need a CLI interface.

## 2018-02-06

To provide for more idiomatic package naming the following packages have been renamed:

* tot/cdtp/application_cache -> tot/cdtp/application/cache
* tot/cdtp/cache_storage -> tot/cdtp/cache/storage
* tot/cdtp/device_orientation -> tot/cdtp/device/orientation
* tot/cdtp/dom_debugger -> tot/cdtp/dom/debugger
* tot/cdtp/dom_snapshot -> tot/cdtp/dom/snapshot
* tot/cdtp/dom_storage -> tot/cdtp/dom/storage
* tot/cdtp/headless_experimental -> tot/cdtp/headless/experimental
* tot/cdtp/heap_profiler -> tot/cdtp/heap/profiler
* tot/cdtp/indexed_db -> tot/cdtp/indexed/db
* tot/cdtp/layer_tree -> tot/cdtp/layer/tree
* tot/cdtp/service_worker -> tot/cdtp/service/worker
* tot/cdtp/system_info -> tot/cdtp/system/info

## 2018-01-21

The core data types have been updated with proper enum support for string properties that only allow specific values.

## 2018-01-12

In preparation for having stable versions of the library, I updated the directory structure of the repo adding a `tot` package for the current code (Tip-of-Tree).

## 2018-01-05

I merged some changes that did change the API a bit. Mainly, all the protocol methods now return a channel instead of blocking until they get a result to better handle the nature of socket data streams. This makes unit testing easier and cleaner and the API more useful, but the race detector still finds false positives due to writing test data to a stack that mocks a socket data stream, which is being drained by an independent goroutine...

I'm not sure what to do with that, or if I care at this point. You can see it by running `go test -race ./...`.
