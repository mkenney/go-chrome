# go-chrome
<table><tbody><tr>
    <td>
        <a href="https://github.com/mkenney/stability-badges#unstable"><img src="https://img.shields.io/badge/stability-unstable-yellow.svg" alt="Unstable"></a>
    </td>
    <td rowspan="6">
        This package aims to be a complete <a href="https://chromedevtools.github.io/devtools-protocol/">Chrome DevTools Protocol</a> implementation. My primary use-case is interacting with <a href="https://developers.google.com/web/updates/2017/04/headless-chrome">headless Google Chrome</a> in a container environment, but this should be appropriate for developing server side and desktop applications as well.
        <br><br>
        This is a work in progress. I think most things are workable now but I still need to add tests and happy-path test more of the APIs. I have <a href="https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot">`Page.captureScreenshot`</a> working fairly well and have been using that for validating changes so far. This implementation is based on the <a href="https://chromedevtools.github.io/devtools-protocol/tot/">Tip-of-Tree</a> documentation and may be prone to change. At some point I will implement stable versions as well.
        <br><br>
        The API should be stable but is not yet production ready.
    </td>
</tr>
<tr>
    <td width="150">
        <a href="https://travis-ci.org/mkenney/go-chrome"><img src="https://travis-ci.org/mkenney/go-chrome.svg?branch=master" alt="Build status"></a>
    </td>
</tr><tr>
    <td width="150">
        <a href="https://codecov.io/gh/mkenney/go-chrome"><img src="https://codecov.io/gh/mkenney/go-chrome/branch/master/graph/badge.svg" alt="Coverage status"></a>
    </td>
</tr><tr>
    <td>
        <a href="https://github.com/mkenney/go-chrome/issues"><img src="https://img.shields.io/github/issues-raw/mkenney/go-chrome.svg" alt="Github issues"></a>
    </td>
</tr>
<tr>
    <td>
        <a href="https://goreportcard.com/report/github.com/mkenney/go-chrome"><img src="https://goreportcard.com/badge/github.com/mkenney/go-chrome" alt="Go Report Card"></a>
    </td>
</tr>
<tr>
    <td>
        <a href="https://godoc.org/github.com/mkenney/go-chrome"><img src="https://godoc.org/github.com/mkenney/go-chrome?status.svg" alt="GoDoc"></a>
    </td>
</tr></tbody></table>

# Examples

I added a few small examples to the [wiki](https://github.com/mkenney/go-chrome/wiki).

# TODO

Add more tests. So many tests... If you would like to contribute but aren't sure how, there are several [open issues](https://github.com/mkenney/go-chrome/issues?q=is%3Aopen+is%3Aissue+project%3Amkenney%2Fgo-chrome%2F1) for tests that need to be written, I'll try to keep that list up to date. There are also [several examples](https://github.com/mkenney/go-chrome/blob/master/socket/protocol.animation_test.go) of tests that have already be written.

Any contributions to tests or otherwise are very welcome!
