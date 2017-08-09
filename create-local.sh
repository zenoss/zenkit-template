#!/usr/bin/env bash
: ${ROOTDIR:=$PWD/$1}

docker run --rm -i \
	-v $GOPATH/src:/go/src \
	-w /go/src/`dirname ${ROOTDIR#$GOPATH/src/}` \
	-e LOCAL_USER_ID=$(id -u) \
	-e IN_DOCKER=1 \
	zenoss/zenkit-build:1.5 \
	/usr/local/bin/create-zenkit-local.sh $1
(cd $1; make)
