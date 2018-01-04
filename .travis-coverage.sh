#!/bin/bash
set -e
echo "" > coverage.txt

go get -v github.com/golang/lint/golint
[ "0" = "$?" ] || exit 1

go get -u github.com/golang/dep/cmd/dep
[ "0" = "$?" ] || exit 2

cd /go/src/github.com/mkenney/go-chrome
[ "0" = "$?" ] || exit 3

dep ensure
[ "0" = "$?" ] || exit 4

for d in $(go list ./... | grep -v vendor); do
    go test -coverprofile=profile.out $d
    exit_code=$?
    if [ "0" != "$exit_code" ]; then
        exit $exit_code
    fi
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

bash <(curl -s https://codecov.io/bash)
[ "0" = "$?" ] || exit 5

exit 0

