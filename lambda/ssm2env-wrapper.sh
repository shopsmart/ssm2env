#!/usr/bin/env bash

set -euo pipefail

if [ -n "${SSM2ENV_PATH:-}" ]; then
  echo "[INFO ] Pulling ssm parameters from $SSM2ENV_PATH" >&2
  eval "$(ssm2env --export "$SSM2ENV_PATH")"
fi

exec "$@"
