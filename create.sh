#!/usr/bin/env bash
docker run --rm -it \
	-v $GOPATH/src:/go/src \
	-w /go/src/${PWD#$GOPATH/src/} \
	-e LOCAL_USER_ID=$(id -u) \
	zenoss/zenkit-build:1.0 \
	/usr/local/bin/create-zenkit.sh $1
