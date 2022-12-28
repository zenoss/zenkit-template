#!/usr/bin/env bash

set -uo pipefail

# shellcheck source=create.sh
source "${BASH_SOURCE%/*}/create.sh"

main() {
	execute_template_in_docker "$@"
}

docker_args=(
	-v "$PWD:/workspace/zenkit-template"
)

main "$1" "create-zenkit-local.sh" "${docker_args[@]}"
