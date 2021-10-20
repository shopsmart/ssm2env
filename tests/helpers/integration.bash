#!/usr/bin/env bash

function integration_setup() {
  INTEGRATION_EXPECTED_ENV_FILE="$BATS_TEST_TMPDIR/expected.env"
  integration_expected_env "$INTEGRATION_EXPECTED_ENV_FILE"
  export INTEGRATION_EXPECTED_ENV_FILE
}

function integration_teardown() {
  :
}

function integration_expected_value() {
  echo "$TEST_SSM2ENV_PARAMETERS" | jq -r . >/dev/null || {
    echo "[ERROR] Unable to decode parameters as json" >&2
    return 255
  }

  val="$(echo "$TEST_SSM2ENV_PARAMETERS" | jq -r .'["'"$1"'"]')"
  printf "%q" "$val"
}

function integration_expected_env() {
  local file="$1"

  echo "$TEST_SSM2ENV_PARAMETERS" | jq -r . >/dev/null || {
    echo "[ERROR] Unable to decode parameters as json" >&2
    return 255
  }

  echo "$TEST_SSM2ENV_PARAMETERS" |
    jq -rc 'keys[]' |
    while read -r key; do

      value="$(integration_expected_value "$key")"
      KEY="${key^^}"

      echo "$KEY=$value"

    done > "$file"
}

function compare_envs() {
  local file1="$1"
  local file2="$2"

  local tmp_file1="$BATS_TEST_TMPDIR/file1.txt"
  local tmp_file2="$BATS_TEST_TMPDIR/file2.txt"

  # reset files
  echo -n '' > "$tmp_file1"
  echo -n '' > "$tmp_file2"

  echo "$TEST_SSM2ENV_PARAMETERS" |
    jq -rc 'keys[]' |
    while read -r key; do

      KEY=${key^^}

      bash -c "source '$file1' && echo $KEY=\$$KEY" >> "$tmp_file1"
      bash -c "source '$file2' && echo $KEY=\$$KEY" >> "$tmp_file2"

    done

  diff -y "$tmp_file1" "$tmp_file2"
}
