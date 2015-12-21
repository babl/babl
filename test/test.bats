#!/usr/bin/env bats

load test_helper


@test "execution w/o parameter" {
  run babl
  assert_fail
  assert_equal "Incorrect Usage." "${lines[0]}"
}

@test "show help" {
  run babl help
  assert_success
  assert_equal "${lines[2]}" "USAGE:"
}

@test "show version" {
  run babl --version
  assert_success
  assert_output "babl version 0.2.0"
  # [[ "$output" =~ "^babl version [0-9]\.[0-9]\.[0-9]$" ]]
  # [ $output =~ "babl version 0.2.0" ]
}
