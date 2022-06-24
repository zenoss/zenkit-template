#!/usr/bin/env bash
: "${ROOTDIR:=$PWD/$1}"

if docker run --rm \
	-v "$GOPATH"/src:/go/src \
	-w /go/src/"$(dirname "${ROOTDIR#"$GOPATH"/src/}")" \
	-e LOCAL_USER_ID="$(id -u)" \
	-e IN_DOCKER=1 \
	zenoss/zenkit-build:1.12.0 \
	/usr/local/bin/create-zenkit-local.sh "$1"
then
    echo "Complete. You should vendor dependencies with the following commands."
    echo
    echo "    cd $1"
    echo "    make vendor"
    echo
fi
