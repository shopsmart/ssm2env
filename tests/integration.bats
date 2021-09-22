#!/usr/bin/env bats

export TEST_SSM2ENV_EXECUTABLE="${TEST_SSM2ENV_EXECUTABLE:-$BATS_TEST_DIRNAME/../bin/ssm2env}"

load helpers/integration.bash

function setup() {
  integration_setup
}

function teardown() {
  integration_teardown
}

@test "it should require 1 arg" {
  [ $TEST_SSM2ENV_INTEGRATION -eq 0 ] || skip

  run "$EXECUTABLE"

  [ $status -ne 0 ]
}

@test "it should pull the ssm params from the prefix" {
  [ $TEST_SSM2ENV_INTEGRATION -eq 0 ] || skip

  run "$EXECUTABLE" "$TEST_SSM2ENV_PREFIX"

  [ $status -eq 0 ]

  echo "$output" > "$BATS_TEST_TMPDIR/ssm2env.env"
  compare_envs "$INTEGRATION_EXPECTED_ENV_FILE" "$BATS_TEST_TMPDIR/ssm2env.env"
}
