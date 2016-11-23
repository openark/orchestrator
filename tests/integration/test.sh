#!/bin/bash

# Local integration tests. To be used by CI.
# See https://github.com/github/orchestrator/tree/doc/local-tests.md
#

# Usage: localtests/test/sh [filter]
# By default, runs all tests. Given filter, will only run tests matching given regep

tests_path=$(dirname $0)
test_logfile=/tmp/orchestrator-test.log
test_outfile=/tmp/orchestrator-test.out
test_diff_file=/tmp/orchestrator-test.diff
orchestrator_binary=/tmp/orchestrator-test
exec_command_file=/tmp/orchestrator-test.bash

test_pattern="${1:-.}"

verify_mysql() {
  if [ "$(mysql test -e "select 1" -ss)" != "1" ] ; then
    echo "Cannot verify mysql"
    exit 1
  fi
}

exec_cmd() {
  echo "$@"
  command "$@" 1> $test_outfile 2> $test_logfile
  return $?
}

echo_dot() {
  echo -n "."
}

test_single() {
  local test_name
  test_name="$1"

  echo -n "Testing: $test_name"

  echo_dot
  mysql --default-character-set=utf8mb4 test < $tests_path/create-per-test.sql

  echo_dot
  if [ -f $tests_path/$test_name/create.sql ] ; then
    mysql --default-character-set=utf8mb4 test < $tests_path/$test_name/create.sql
  fi

  extra_args=""
  if [ -f $tests_path/$test_name/extra_args ] ; then
    extra_args=$(cat $tests_path/$test_name/extra_args)
  fi
  #
  cmd="$orchestrator_binary \
    --config=${tests_path}/orchestrator.conf.json
    --debug \
    --stack \
    ${extra_args[@]}"
  echo_dot
  echo $cmd > $exec_command_file
  echo_dot
  bash $exec_command_file 1> $test_outfile 2> $test_logfile

  execution_result=$?

  if [ -f $tests_path/$test_name/destroy.sql ] ; then
    mysql --default-character-set=utf8mb4 test < $tests_path/$test_name/destroy.sql
  fi

  if [ -f $tests_path/$test_name/expect_failure ] ; then
    if [ $execution_result -eq 0 ] ; then
      echo
      echo "ERROR $test_name execution was expected to exit on error but did not. cat $test_logfile"
      return 1
    fi
    if [ -s $tests_path/$test_name/expect_failure ] ; then
      # 'expect_failure' file has content. We expect to find this content in the log.
      expected_error_message="$(cat $tests_path/$test_name/expect_failure)"
      if grep -q "$expected_error_message" $test_logfile ; then
          return 0
      fi
      echo
      echo "ERROR $test_name execution was expected to exit with error message '${expected_error_message}' but did not. cat $test_logfile"
      return 1
    fi
    # 'expect_failure' file has no content. We generally agree that the failure is correct
    return 0
  fi

  if [ $execution_result -ne 0 ] ; then
    echo
    echo "ERROR $test_name execution failure. cat $test_logfile"
    return 1
  fi

  if [ -f $tests_path/$test_name/expect_output ] ; then
    diff $tests_path/$test_name/expect_output $test_outfile > $test_diff_file
    diff_result=$?
    if [ $diff_result -ne 0 ] ; then
      echo
      echo "ERROR $test_name diff failure. cat $test_diff_file"
      return 1
    fi
  fi

  # all is well
  return 0
}

build_binary() {
  echo "Building"
  go build -o $orchestrator_binary go/cmd/orchestrator/main.go
}

test_all() {
  build_binary
  if [ $? -ne 0 ] ; then
    echo "ERROR build failed"
    return 1
  fi
  find $tests_path ! -path . -type d -mindepth 1 -maxdepth 1 | xargs ls -td1 | cut -d "/" -f 4 | egrep "$test_pattern" | while read test_name ; do
    test_single "$test_name"
    if [ $? -ne 0 ] ; then
      echo "+ FAIL"
      return 1
    else
      echo
      echo "+ pass"
    fi
  done
}

verify_mysql
test_all
