#!/usr/bin/env bash

set -euo pipefail

if [ -n "${SSM2ENV_PATH:-}" ]; then
  echo "[INFO ] Pulling ssm parameters from $SSM2ENV_PATH" >&2
  eval "$(ssm2env --export "$SSM2ENV_PATH")"
fi

# Lambda allows one exec wrapper, which this takes; hand off to another (such as OpenTelemetry's
#  /opt/otel-handler) after loading the parameters
if [ -n "${SSM2ENV_EXEC_WRAPPER:-}" ]; then
  exec "$SSM2ENV_EXEC_WRAPPER" "$@"
fi

exec "$@"
