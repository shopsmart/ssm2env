#!/usr/bin/env bats

[ -n "$TEST_SSM2ENV_EXECUTABLE" ] || {
  TEST_SSM2ENV_EXECUTABLE="$BATS_TEST_DIRNAME/../bin/ssm2env"
  if [ $TEST_SSM2ENV_INTEGRATION_MOCK_CLIENT -eq 0 ]; then
    TEST_SSM2ENV_EXECUTABLE="${TEST_SSM2ENV_EXECUTABLE}mock"
  fi
}
export TEST_SSM2ENV_EXECUTABLE


load helpers/integration.bash

function setup() {
  integration_setup
}

function teardown() {
  integration_teardown
}

@test "it should log a warning when using mock client" {
  [ $TEST_SSM2ENV_INTEGRATION_MOCK_CLIENT -eq 0 ] || skip

  run "$TEST_SSM2ENV_EXECUTABLE"

  echo "$output"

  [ $status -eq 0 ]
  [[ "$output" =~ ^.*level=warning\ msg=\"Using\ ssm\ mock\ for\ params\".*$ ]]
}

@test "it should print help if no args are provided" {
  [ $TEST_SSM2ENV_INTEGRATION -eq 0 ] || skip

  run "$TEST_SSM2ENV_EXECUTABLE"

  [ $status -eq 0 ]
}

@test "it should print the version" {
  [ $TEST_SSM2ENV_INTEGRATION -eq 0 ] || skip

  run "$TEST_SSM2ENV_EXECUTABLE" --version

  [ $status -eq 0 ]
}

@test "it should pull the ssm params from the prefix" {
  [ $TEST_SSM2ENV_INTEGRATION -eq 0 ] || skip

  run "$TEST_SSM2ENV_EXECUTABLE" "$TEST_SSM2ENV_PREFIX"

  [ $status -eq 0 ]

  echo "$output" > "$BATS_TEST_TMPDIR/ssm2env.env"
  compare_envs "$INTEGRATION_EXPECTED_ENV_FILE" "$BATS_TEST_TMPDIR/ssm2env.env"
}
