#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Creates a .env file for each an application, if it don't exist.
#
# Examples
#
#   ./scripts/setup.sh
#
# Returns exit code 0.
function main() {
  local app_dir
  local env_example_file

  set_vars

  for app_dir in ./apps/*/; do
    env_example_file="${app_dir}/.env.example"

    if [[ ! -f "${env_example_file}" ]]; then
      printf "%b no .env.example at %b, ignoring \n" "${INFO_PREFIX}" "${env_example_file}"
      continue
    fi

    printf "%b creating file: %b.env \n" "${INFO_PREFIX}" "${app_dir}"
    # copies the .env.example if it doesn't exist (uses no clobber [-n])
    cp -n "${app_dir}/.env.example" "${app_dir}/.env"
  done

  exit 0
}

# and so, it begins...
main
