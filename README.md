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

# Examples

There are a few small examples of how to use the framework API on the [wiki](https://github.com/mkenney/go-chrome/wiki) and in the [`/_examples`](https://github.com/mkenney/go-chrome/tree/master/_examples) directory.

# TODO

Contributions of any kind are very welcome!

* Resolve [race condition issues](https://github.com/mkenney/go-chrome/pull/119). Any assistance is appreciated!
* Add framework API examples to the `/_examples` directory and wiki to showcase various ways people are using the package.

  Any example scripts showing various ways people are using the framework would be outstanding! The [screenshot script](https://github.com/mkenney/go-chrome/tree/master/_examples/screenshot-url) and several others are available there.

* Refactoring to implement standard interfaces where applicable and review current use of interfaces in the API. Some aren't needed at all and others are used to support test mocks.
* Add more tests, particularly for error cases.
* Add integrated tests to stablize package interactions raised in various issues.

If you would like to contribute but aren't sure how, take a look at the [issue tracker](https://github.com/mkenney/go-chrome/issues). Issues are labeled as bug reports, feature requests, feedback requests, help wanted, etc.

There are also always [tests that could be written](https://codecov.io/gh/mkenney/go-chrome). There are [many](https://github.com/mkenney/go-chrome/blob/master/tot/socket/cdtp.animation_test.go) [examples](https://github.com/mkenney/go-chrome/blob/master/tot/cdtp/animation/enum.animation.type_test.go) of tests in the package.

# [`CHANGELOG`](CHANGELOG.md)**

All notable changes to this project are documented in the [`CHANGELOG`](CHANGELOG.md). The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
