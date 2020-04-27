#!/usr/bin/env bash
if docker run --rm -it \
	-v $GOPATH/src:/go/src \
	-w /go/src/${PWD#$GOPATH/src/} \
	-e LOCAL_USER_ID=$(id -u) \
	-e IN_DOCKER=1 \
	zenoss/zenkit-build:1.8.0 \
	/usr/local/bin/create-zenkit.sh $1
then
    echo "Complete. You should vendor dependencies with the following commands."
    echo
    echo "    cd $1"
    echo "    make vendor"
    echo
fi
