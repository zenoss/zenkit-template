#!/usr/bin/env bash

set -uo pipefail

# shellcheck source=create.sh
source create.sh

main() {
	execute_template_in_docker "$@"
}

main "$1" "create-zenkit-local.sh"
