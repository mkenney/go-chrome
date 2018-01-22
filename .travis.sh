#!/bin/sh
set -e

docker run \
    --rm \
    -v $(pwd):/go/src/github.com/mkenney/go-chrome \
    --entrypoint="/go/src/github.com/mkenney/go-chrome/.travis.entrypoint.sh" \
    golang:1.9-alpine