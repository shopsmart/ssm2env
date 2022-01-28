#!/usr/bin/env bats

function setup() {
  export TEST_SSM2ENV_EXECUTABLE="${TEST_SSM2ENV_EXECUTABLE:-$BATS_TEST_DIRNAME/../bin/ssm2env}"
}

function teardown() {
  :
}

@test "it should print help if no args are provided" {
  run "$TEST_SSM2ENV_EXECUTABLE"

  echo "$output"

  [ $status -eq 0 ]
}

@test "it should print the version" {
  run "$TEST_SSM2ENV_EXECUTABLE" --version

  echo "$output"

  [ $status -eq 0 ]
}

@test "it should pull the ssm params from the prefix" {
  run "$TEST_SSM2ENV_EXECUTABLE" /aws/service/global-infrastructure/regions

  [ $status -eq 0 ]

  echo "$output"

  echo "$output" | sort > "$BATS_TEST_TMPDIR/ssm2env.env"
  diff -y "$BATS_TEST_DIRNAME/expected.env" "$BATS_TEST_TMPDIR/ssm2env.env"
}
