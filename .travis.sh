#!/bin/sh
set -e

go get -v github.com/golang/lint/golint
[ "0" = "$?" ] || exit 1

go get -u github.com/golang/dep/cmd/dep
[ "0" = "$?" ] || exit 2

dep ensure
[ "0" = "$?" ] || exit 3

rm -f coverage.txt
for dir in $(go list ./... | grep -v vendor); do
    go test -race -timeout 300s -coverprofile=profile.out $dir
    exit_code=$?
    if [ "0" != "$exit_code" ]; then
        exit $exit_code
    fi
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

exit_code=0
for dir in $(go list ./... | grep -v vendor); do
    echo "golint $dir"
    result=$(golint $dir)
    if [ "" != "$result" ]; then
        echo $result
        exit_code=5
    fi
done

exit $exit_code
