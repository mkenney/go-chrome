#!/bin/sh
set -e

exit_code=0

go get -u golang.org/x/lint/golint
[ "0" = "$?" ] || exit 1

go get -u github.com/golang/dep/cmd/dep
[ "0" = "$?" ] || exit 2

dep ensure
[ "0" = "$?" ] || exit 3

for dir in $(go list ./... | grep -v vendor); do
    echo "golint -s $dir"
    result=$(golint -s $dir)
    if [ "" != "$result" ]; then
        echo $result
        exit_code=5
    fi
    if [ "0" != "$exit_code" ]; then
        exit $exit_code
    fi
done

rm -f coverage.txt
for dir in $(go list ./... | grep -v vendor); do
    go test -timeout 300s -coverprofile=profile.out $dir
    exit_code=$?
    if [ "0" != "$exit_code" ]; then
        exit $exit_code
    fi
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

exit $exit_code
