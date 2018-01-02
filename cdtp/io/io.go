/*
Package io provides type definitions for use with the Chrome IO protocol

https://chromedevtools.github.io/devtools-protocol/tot/IO/
*/
package io

/*
StreamHandle is either obtained from another method or specified as blob:<uuid> where <uuid> is an
UUID of a Blob.

https://chromedevtools.github.io/devtools-protocol/tot/IO/#type-StreamHandle
*/
type StreamHandle string
