#!/usr/bin/env bats

function setup() {
  export WRAPPER="$BATS_TEST_DIRNAME/../lambda/ssm2env-wrapper.sh"
  unset SSM2ENV_PATH SSM2ENV_EXEC_WRAPPER

  # Stands in for another exec wrapper: prints that it ran, then runs the rest
  export NEXT_WRAPPER="$BATS_TEST_TMPDIR/next-wrapper.sh"
  printf '#!/usr/bin/env bash\necho next-wrapper\nexec "$@"\n' > "$NEXT_WRAPPER"
  chmod +x "$NEXT_WRAPPER"
}

@test "it should run the command" {
  run "$WRAPPER" echo runtime

  [ $status -eq 0 ]
  [ "$output" = "runtime" ]
}

@test "it should hand off to SSM2ENV_EXEC_WRAPPER when set" {
  SSM2ENV_EXEC_WRAPPER="$NEXT_WRAPPER" run "$WRAPPER" echo runtime

  [ $status -eq 0 ]
  [ "${lines[0]}" = "next-wrapper" ]
  [ "${lines[1]}" = "runtime" ]
}
