<p align="center">
    <img src="https://github.com/mkenney/go-chrome/wiki/assets/images/gopher-logo.png" width="300px">
</p>

# go-chrome

<table><tbody><tr>
    <td width="150">
        <a href="https://travis-ci.org/mkenney/go-chrome"><img src="https://travis-ci.org/mkenney/go-chrome.svg?branch=master" alt="Build status"></a>
    </td>
    <td rowspan="6">
        This package aims to be a complete <a href="https://chromedevtools.github.io/devtools-protocol/">Chrome DevTools Protocol</a> implementation. The primary use-case behind this project is interacting with <a href="https://developers.google.com/web/updates/2017/04/headless-chrome">headless Google Chrome</a> in a container environment, but it should be appropriate for developing server side and desktop applications as well.
        <br><br>
        The API is fairly settled and most functions are workable now but more automated tests and real-world happy-path testing is needed. <a href="https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot">`Page.captureScreenshot`</a> is working fairly well and is regularly used for validating the workability of API changes.
        <br /><br />
        This implementation is based on the <a href="https://chromedevtools.github.io/devtools-protocol/tot/">Tip-of-Tree</a> documentation and may be prone to change. At some point stable versions will be implemented as well.
    </td>
</tr>
<tr>
    <td>
        <a href="https://github.com/mkenney/go-chrome/issues"><img src="https://img.shields.io/github/issues-raw/mkenney/go-chrome.svg" alt="Github issues"></a>
    </td>
</tr><tr>
    <td>
        <a href="https://github.com/mkenney/stability-badges#stable"><img src="https://img.shields.io/badge/stability-stable-33BBFF.svg" alt="Stable"></a>
    </td>
</tr><tr>
    <td width="150">
        <a href="https://codecov.io/gh/mkenney/go-chrome"><img src="https://img.shields.io/codecov/c/github/mkenney/go-chrome/master.svg" alt="Coverage status"></a>
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

Add more tests. So many tests... If you would like to contribute but aren't sure how, there are several [open issues](https://github.com/mkenney/go-chrome/issues?q=is%3Aopen+is%3Aissue+project%3Amkenney%2Fgo-chrome%2F1) for tests that need to be written, I'll try to keep that list up to date. There are also [several examples](https://github.com/mkenney/go-chrome/blob/master/socket/cdtp.animation_test.go) of tests that have already be written.

Any contributions to tests or otherwise are very welcome!

# Change Log

## 2017-12-05

In preparation for having stable versions of the library, I updated the directory structure of the repo adding a `tot` package for the current code (Tip-of-Tree). This should be the last major refactoring of this package and no further breaking changes are expected.

## 2017-01-05

I merged some changes that did change the API a bit. Mainly, all the protocol methods now return a channel instead of blocking until they get a result to better handle the nature of socket data streams. This makes unit testing easier and cleaner and the API more useful, but the race detector still finds false positives due to writing test data to a stack that mocks a socket data stream, which is being drained by an independent goroutine...

I'm not sure what to do with that, or if I care at this point. You can see it by running `go test -race ./...`.
