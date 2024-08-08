#!/usr/bin/env bash

set -uo pipefail

ZENKIT_BUILD_TAG="zenoss/zenkit-build:1.18.2"

execute_template_in_docker() {
	local path="$1"
	shift
	local script="$1"
	shift
	local additional_docker_args=("$@")

	NEWPROJ_NAME="$(basename "$path")"
	NEWPROJ_PATH="$(realpath "$(dirname "$path")")"

	DOCKER_ARGS=(
		--rm
		-it
		-v "$NEWPROJ_PATH:/workspace/tmp/"
		-w "/workspace/tmp/"
		-e LOCAL_USER_ID="$(id -u)"
		-e IN_DOCKER=1
		"${additional_docker_args[@]}"
	)

	if
		docker run \
			"${DOCKER_ARGS[@]}" "$ZENKIT_BUILD_TAG" \
			"$script" "$NEWPROJ_NAME"
	then
		echo "Complete. You should vendor dependencies with the following commands."
		echo
		echo "    cd $NEWPROJ_PATH/$NEWPROJ_NAME"
		echo "    make vendor"
		echo
	fi
}

main() {
	execute_template_in_docker "$@"
}

[[ "$0" == "${BASH_SOURCE[0]}" ]] && main "$1" "create-zenkit.sh"
