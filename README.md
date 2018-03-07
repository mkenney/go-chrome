<p align="center">
    <a href="https://gopherize.me/gopher/255e20ee48c85f3b4701446e2513c100f22129f3"><img src="https://github.com/mkenney/go-chrome/wiki/assets/images/gopher-logo.png" width="300px"></a>
</p>

# go-chrome

<table><tbody><tr>
    <td>
        <a href="https://github.com/mkenney/go-chrome/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mkenney/go-chrome.svg" alt="MIT License"></a>
    </td>
    <td rowspan="7">
        This package aims to be a complete <a href="https://chromedevtools.github.io/devtools-protocol/">Chrome DevTools Protocol</a> implementation. The primary use-case behind this project is interacting with <a href="https://developers.google.com/web/updates/2017/04/headless-chrome">headless Google Chrome</a> in a container environment, but it should be appropriate for developing server side and desktop applications as well.
        <br><br>
        The API is fairly settled and basic code-coverage tests have been implemented but real-world testing is needed. <a href="https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot">`Page.captureScreenshot`</a> and related calls are working well and are regularly used for validating the workability of code refactoring.
        <br /><br />
        This implementation is based on the <a href="https://chromedevtools.github.io/devtools-protocol/tot/">Tip-of-Tree</a> documentation and may be prone to change. At some point stable versions will be implemented as well, hopefully beginning with v1.3.
    </td>
</tr><tr>
    <td>
        <a href="https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#beta"><img src="https://img.shields.io/badge/stability-beta-33bbff.svg" alt="Beta"></a>
    </td>
</tr><tr>
    <td width="150">
        <a href="https://travis-ci.org/mkenney/go-chrome"><img src="https://travis-ci.org/mkenney/go-chrome.svg?branch=master" alt="Build status"></a>
    </td>
</tr><tr>
    <td width="150">
        <a href="https://codecov.io/gh/mkenney/go-chrome"><img src="https://img.shields.io/codecov/c/github/mkenney/go-chrome/master.svg" alt="Coverage status"></a>
    </td>
</tr><tr>
    <td>
        <a href="https://github.com/mkenney/go-chrome/issues"><img src="https://img.shields.io/github/issues-raw/mkenney/go-chrome.svg" alt="Github issues"></a>
    </td>
</tr><tr>
    <td>
        <a href="https://goreportcard.com/report/github.com/mkenney/go-chrome"><img src="https://goreportcard.com/badge/github.com/mkenney/go-chrome" alt="Go Report Card"></a>
    </td>
</tr><tr>
    <td>
        <a href="https://godoc.org/github.com/mkenney/go-chrome"><img src="https://godoc.org/github.com/mkenney/go-chrome?status.svg" alt="GoDoc"></a>
    </td>
</tr></tbody></table>

# Examples

There are a few small examples of how to use the framework API on the [wiki](https://github.com/mkenney/go-chrome/wiki).

# TODO

Add more tests, especially for error cases. If you would like to contribute but aren't sure how, take a look at [codecov](https://codecov.io/gh/mkenney/go-chrome) for any tests that could be written. There are [many](https://github.com/mkenney/go-chrome/blob/master/tot/socket/cdtp.animation_test.go) [examples](https://github.com/mkenney/go-chrome/blob/master/tot/cdtp/animation/enum.animation.type_test.go)of tests in the repo.

Any contributions of any kind are very welcome!

# Change Log

## 2017-02-06

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

## 2017-12-21

The core data types have been updated with proper enum support for string properties that only allow specific values.

## 2017-12-05

In preparation for having stable versions of the library, I updated the directory structure of the repo adding a `tot` package for the current code (Tip-of-Tree).

## 2017-01-05

I merged some changes that did change the API a bit. Mainly, all the protocol methods now return a channel instead of blocking until they get a result to better handle the nature of socket data streams. This makes unit testing easier and cleaner and the API more useful, but the race detector still finds false positives due to writing test data to a stack that mocks a socket data stream, which is being drained by an independent goroutine...

I'm not sure what to do with that, or if I care at this point. You can see it by running `go test -race ./...`.
