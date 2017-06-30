#!/bin/bash

USER=$1
NAME=$2

cleanup () {
	chown -R $USER $NAME
}
trap cleanup EXIT

boilr template download zenoss/zenkit-template tpl
boilr template use tpl $NAME
