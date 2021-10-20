#!/usr/bin/env bash

set -euo pipefail

__DIR__="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"

pushd "$__DIR__" >/dev/null || exit 1

{
  {
    terraform init -no-color
    terraform refresh -no-color
  } >"$__DIR__/parameters.stdout.log"

  terraform output -json -no-color parameters

} 2>"$__DIR__/parameters.stderr.log"

popd >/dev/null || exit 1
