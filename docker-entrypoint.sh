#!/bin/bash

USER=$1
NAME=$2

cleanup () {
	chown -R $USER $NAME
}
trap cleanup EXIT

boilr template download zenoss/zenkit-template tpl

/usr/bin/expect -f <(cat <<EOF
spawn -noecho boilr template use tpl $NAME
expect -re ".*Please choose a value for \"Name\".*"
send "$NAME\n"

interact

exit
EOF
)
