#!/usr/bin/env bash

set -euo pipefail

echo "[DEBUG] ssm2env layer start" >&2

if [ -z "${SSM2ENV_PATH:-}" ]; then
  echo "[INFO ] No ssm2env path provided" >&2
  exit 0
fi

echo "[INFO ] Pulling ssm parameters from $SSM2ENV_PATH" >&2
eval "$(ssm2env --export "$SSM2ENV_PATH")"

echo "[DEBUG] ssm2env layer end" >&2
