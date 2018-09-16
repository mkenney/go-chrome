<p align="center">
    <a href="https://gopherize.me/gopher/255e20ee48c85f3b4701446e2513c100f22129f3"><img src="https://github.com/mkenney/go-chrome/wiki/assets/images/gopher-logo.png" width="300px"></a>
</p>

# go-chrome

<p align="center">
	<a href="https://github.com/mkenney/go-chrome/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-BSD-green.svg" alt="BSD-2-Clause"></a>
	<a href="https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#release-candidate"><img src="https://img.shields.io/badge/stability-pre--release-48c9b0.svg" alt="Release Candidate"></a>
	<a href="https://travis-ci.org/mkenney/go-chrome"><img src="https://travis-ci.org/mkenney/go-chrome.svg?branch=master" alt="Build status"></a>
	<a href="https://codecov.io/gh/mkenney/go-chrome"><img src="https://img.shields.io/codecov/c/github/mkenney/go-chrome/master.svg" alt="Coverage status"></a>
	<a href="https://goreportcard.com/report/github.com/mkenney/go-chrome"><img src="https://goreportcard.com/badge/github.com/mkenney/go-chrome" alt="Go Report Card"></a>
	<a href="https://github.com/mkenney/go-chrome/issues"><img src="https://img.shields.io/github/issues-raw/mkenney/go-chrome.svg" alt="Github issues"></a>
	<a href="https://github.com/mkenney/go-chrome/pulls"><img src="https://img.shields.io/github/issues-pr/mkenney/go-chrome.svg" alt="Github pull requests"></a>
	<a href="https://godoc.org/github.com/mkenney/go-chrome"><img src="https://godoc.org/github.com/mkenney/go-chrome?status.svg" alt="GoDoc"></a>
</p>

This package aims to be a complete [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/) implementation. The primary use-case behind this project is interacting with [headless Google Chrome](https://developers.google.com/web/updates/2017/04/headless-chrome) in a container environment, but it should be appropriate for developing server side and desktop applications for any browser that supports the devtools protocol.

The API is fairly settled and basic code-coverage tests have been implemented but real-world testing is needed. [`Page.captureScreenshot`](https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot) and related calls are working well and are regularly used for validating the viability of code changes.

This implementation is based on the [Tip-of-Tree](https://chromedevtools.github.io/devtools-protocol/tot/) documentation and may be prone to change. At some point stable versions will be implemented as well, hopefully beginning with v1.3.

# Announcements

## v1.0.0-rc5 released

[`v1.0.0-rc5`](https://github.com/mkenney/go-chrome/releases/tag/v1.0.0-rc5) has been released.

* Fixes a `dep` issue with transitive dependencies

Please [open an issue](https://github.com/mkenney/go-chrome/issues/new/choose) to report any problems or suggest any changes.

## v1.0.0-rc4 released

[`v1.0.0-rc4`](https://github.com/mkenney/go-chrome/releases/tag/v1.0.0-rc4) has been released.

* Updates the log package
* Cleans up and clarifies some log messages

Please [open an issue](https://github.com/mkenney/go-chrome/issues/new/choose) to report any problems or suggest any changes.

## v1.0.0-rc3 released

[`v1.0.0-rc3`](https://github.com/mkenney/go-chrome/releases/tag/v1.0.0-rc3) has been released.

* Fixes an issue with an upstream dependency

Please [open an issue](https://github.com/mkenney/go-chrome/issues/new/choose) to report any problems or suggest any changes.

## v1.0.0-rc2 released

[`v1.0.0-rc2`](https://github.com/mkenney/go-chrome/releases/tag/v1.0.0-rc2) has been released.

* Fixes an issue with zombie listen process
* Fixes an issue with zombie stop process
* Adds test coverage for socket timeouts

Please [open an issue](https://github.com/mkenney/go-chrome/issues/new/choose) to report any problems or suggest any changes.

```toml
[[constraint]]
  name = "github.com/mkenney/go-chrome"
  version = "1.0.0-rc2"
```

## v1.0.0-rc1 released

[`v1.0.0-rc1`](https://github.com/mkenney/go-chrome/releases/tag/v1.0.0-rc1) has been released. Please [open an issue](https://github.com/mkenney/go-chrome/issues/new/choose) to report any problems or suggest any changes.

```toml
[[constraint]]
  name = "github.com/mkenney/go-chrome"
  version = "1.0.0-rc1"
```

# Examples

There are a few small examples of how to use the framework API on the [wiki](https://github.com/mkenney/go-chrome/wiki) and in the [`/test`](https://github.com/mkenney/go-chrome/tree/master/test) directory.

# TODO

Contributions of any kind are very welcome!

* Add framework API examples to the `/test` directory and wiki.

  Any example scripts showing various ways people are using the framework would be outstanding!

* Refactoring to implement standard interfaces where applicable and review current use of interfaces in the API. Some aren't needed at all and others are used to support test mocks.
* Add more tests, especially for error cases. If you would like to contribute but aren't sure how, take a look at [codecov](https://codecov.io/gh/mkenney/go-chrome) for any tests that could be written. There are [many](https://github.com/mkenney/go-chrome/blob/master/tot/socket/cdtp.animation_test.go) [examples](https://github.com/mkenney/go-chrome/blob/master/tot/cdtp/animation/enum.animation.type_test.go) of tests in the repo.
* Add integration test scripts to the `test/` directory to exercise various functions. The [screenshot script](https://github.com/mkenney/go-chrome/blob/master/test/screenshot/main.go) is already setup there.

# Change Log

## 2018-05-30

I have [refactored the package layout](https://github.com/mkenney/go-chrome/pull/98), deleting the `cdtp` package and moving all the packages within it to the `tot` package. This cleans up import statements and make the layout a bit more natural to use. The only refactoring this requires is removing that folder from your import path, changing them from (for example):

```go
import (
	"github.com/mkenney/go-chrome/tot/cdtp/accessibility"
)
```

to simply:

```go
import (
	"github.com/mkenney/go-chrome/tot/accessibility"
)
```

The [`v1.0.0-rc1`](https://github.com/mkenney/go-chrome/releases/tag/v1.0.0-rc1) pre-release candidate coincides with this change. Please [open an issue](https://github.com/mkenney/go-chrome/issues/new/choose) to report any problems or suggest any additional changes to further the [`v1`](https://github.com/mkenney/go-chrome/milestone/1) milestone.



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

It's easier to deal with, simpler to understand, and supporting multiple values wasn't ever useful. An [argument could be made](https://github.com/mkenney/go-chrome/issues/new/choose) for supporting something like [`pkg/flag`](https://golang.org/pkg/flag/) I suppose but I don't generally need a CLI interface.

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
