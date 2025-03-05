#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Injects the version and runs the Go app.
#
# $1 - [optional] a version to inject, otherwise the version from the VERSION file is read.
#
# Examples
#
#   ./bin/dev.sh # reads the version in the VERSION file
#   ./bin/dev.sh "1.2.3"
#
# Returns exit code 0.
function main() {
  local build_dir
  local version

  set_vars

  build_dir="./bin"

  # get the version in the version file
  version=$(<VERSION)

  # if the version argument exists, use it instead of the one on file
  if [ -n "$1" ]; then
    version="$1"
  fi

  VERSION=$version

  # export the version as an env var
  export VERSION

  printf "%b starting app...\n" "${INFO_PREFIX}"
  CompileDaemon -build="go build -o ${build_dir} ./cmd/main.go" -command="${build_dir}"

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main "$@"
