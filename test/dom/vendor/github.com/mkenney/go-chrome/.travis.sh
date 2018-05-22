#!/bin/sh
set -e
docker pull golang:1.10-alpine
docker run \
    --rm \
    -v $(pwd):/go/src/github.com/mkenney/go-chrome \
    --entrypoint="/go/src/github.com/mkenney/go-chrome/.travis.entrypoint.sh" \
    golang:1.10-alpine