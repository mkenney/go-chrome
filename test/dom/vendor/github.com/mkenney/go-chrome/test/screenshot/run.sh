#!/bin/bash
docker run -it --rm \
	-v $(pwd):/go/src/test \
	-v $(pwd):/tmp/test \
	-v $(pwd)/../:/go/src/github.com/mkenney/go-chrome:ro \
	-e LOG_LEVEL=info \
	mkenney/chromium-headless:latest \
		sh -cexu "cd /go/src/test \
			&& go build -o /go/bin/app \
			&& /go/bin/app"
