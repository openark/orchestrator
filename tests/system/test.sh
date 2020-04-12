#!/bin/bash

# Local integration tests. To be used by CI.
# See https://github.com/github/orchestrator/tree/doc/local-tests.md
#

# Usage: localtests/test/sh [mysql|sqlite] [filter]
# By default, runs all tests. Given filter, will only run tests matching given regep

tests_path=$(dirname $0)/tests
test_logfile=/tmp/orchestrator-test.log
test_outfile=/tmp/orchestrator-test.out
test_diff_file=/tmp/orchestrator-test.diff
test_query_file=/tmp/orchestrator-test.sql
test_restore_outfile=/tmp/orchestrator-test-restore.out
test_restore_diff_file=/tmp/orchestrator-test-restore.diff

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
  if [ ! -f $tests_path/$test_name/run ] ; then
    echo
    echo "missing 'run' script"
    return 1
  fi

  echo_dot
  bash $tests_path/$test_name/run 1> $test_outfile 2> $test_logfile
  execution_result=$?

  if [ -f $tests_path/$test_name/restore ] ; then
    bash $tests_path/$test_name/restore
  else
    script/deploy-replication
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
    diff -b $tests_path/$test_name/expect_output $test_outfile > $test_diff_file
    diff_result=$?
    if [ $diff_result -ne 0 ] ; then
      echo
      echo "ERROR $test_name diff failure. cat $test_diff_file"
      echo "---"
      cat $test_diff_file
      echo "---"
      return 1
    fi
  fi

  # all is well
  return 0
}

test_all() {
  test_pattern="${1:-.}"
  find $tests_path ! -path . -type d -mindepth 1 -maxdepth 1 | xargs ls -td1 | cut -d "/" -f 5 | egrep "$test_pattern" | while read test_name ; do
    test_single "$test_name"
    if [ $? -ne 0 ] ; then
      echo "+ FAIL"
      return 1
    else
      echo
      echo "+ pass"
    fi

    bash $tests_path/check_restore > $test_restore_outfile
    diff -b $tests_path/expect_restore $test_restore_outfile > $test_restore_diff_file
    diff_result=$?
    if [ $diff_result -ne 0 ] ; then
      echo
      echo "ERROR $test_name restore failure. cat $test_restore_diff_file"
      echo "---"
      cat $test_restore_diff_file
      echo "---"
      return 1
    fi
  done
}

main() {
  test_all "$@"
}

main "$@"
