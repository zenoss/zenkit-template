#!/usr/bin/env bash

docker run --rm -it -v $GOPATH/src:/go/src -w /go/src/${PWD#$GOPATH/src/} zenoss/zenkit-template:latest $(id -u):$(id -g) $1
