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

  echo "$output" | sort > "$BATS_TEST_TMPDIR/ssm2env.env"
  diff -y "$BATS_TEST_DIRNAME/expected.env" "$BATS_TEST_TMPDIR/ssm2env.env"
}

@test "it should export the ssm params if the --export flag is provided" {
  run "$TEST_SSM2ENV_EXECUTABLE" --export /aws/service/global-infrastructure/regions

  [ $status -eq 0 ]

  echo "$output" | sort > "$BATS_TEST_TMPDIR/ssm2env.env"
  sed 's/^/export /' "$BATS_TEST_DIRNAME/expected.env" | sort > "$BATS_TEST_TMPDIR/sorted.env"
  diff -y "$BATS_TEST_TMPDIR/sorted.env" "$BATS_TEST_TMPDIR/ssm2env.env"
}

@test "it should export the ssm params if the SSM2ENV_EXPORT variable is set" {
  run env SSM2ENV_EXPORT=true "$TEST_SSM2ENV_EXECUTABLE" /aws/service/global-infrastructure/regions

  [ $status -eq 0 ]

  echo "$output" | sort > "$BATS_TEST_TMPDIR/ssm2env.env"
  sed 's/^/export /' "$BATS_TEST_DIRNAME/expected.env" | sort > "$BATS_TEST_TMPDIR/sorted.env"
  diff -y "$BATS_TEST_TMPDIR/sorted.env" "$BATS_TEST_TMPDIR/ssm2env.env"
}

# Validate that flags take precedence over environment variables

@test "it should export the ssm params if the SSM2ENV_EXPORT variable is false but the flag is provided" {
  run env SSM2ENV_EXPORT=false "$TEST_SSM2ENV_EXECUTABLE" --export /aws/service/global-infrastructure/regions

  [ $status -eq 0 ]

  echo "$output" | sort > "$BATS_TEST_TMPDIR/ssm2env.env"
  sed 's/^/export /' "$BATS_TEST_DIRNAME/expected.env" | sort > "$BATS_TEST_TMPDIR/sorted.env"
  diff -y "$BATS_TEST_TMPDIR/sorted.env" "$BATS_TEST_TMPDIR/ssm2env.env"
}
