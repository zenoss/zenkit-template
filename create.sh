#!/usr/bin/env bash
docker run --rm -it \
	-v $GOPATH/src:/go/src \
	-w /go/src/${PWD#$GOPATH/src/} \
	-e LOCAL_USER_ID=$(id -u) \
	-e IN_DOCKER=1 \
	zenoss/zenkit-build:1.6.0 \
	/usr/local/bin/create-zenkit.sh $1
(cd $1; make)
