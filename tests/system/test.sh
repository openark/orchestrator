#!/bin/bash

# Local integration tests. To be used by CI.
# See https://github.com/openark/orchestrator/tree/doc/local-tests.md
#

# Usage: localtests/test/sh [mysql|sqlite] [filter]
# By default, runs all tests. Given filter, will only run tests matching given regep

export tests_path=$(dirname $0)
setup_teardown_logfile=/tmp/orchestrator-setup-teardown.log
test_logfile=/tmp/orchestrator-test.log
test_outfile=/tmp/orchestrator-test.out
test_diff_file=/tmp/orchestrator-test.diff
test_query_file=/tmp/orchestrator-test.sql
test_restore_outfile=/tmp/orchestrator-test-restore.out
test_restore_diff_file=/tmp/orchestrator-test-restore.diff
export deploy_replication_file=/tmp/deploy_replication.log
tests_todo_file=/tmp/system-tests-todo.txt
tests_successful_file=/tmp/system-tests-successful.txt
tests_failed_file=/tmp/system-tests-failed.txt

exec_cmd() {
  echo "$@"
  command "$@" 1> $test_outfile 2> $test_logfile
  return $?
}

echo_dot() {
  echo -n "."
}

check_environment() {
  if ! echo -n "" > $tests_todo_file ; then
    echo "ERROR: unable to write to $tests_todo_file"
    exit 1
  fi
  if ! echo -n "" > $tests_successful_file ; then
    echo "ERROR: unable to write to $tests_successful_file"
    exit 1
  fi
  if ! echo -n "" > $tests_failed_file ; then
    echo "ERROR: unable to write to $tests_failed_file"
    exit 1
  fi
  echo "checking orchestrator-client"
  if ! which orchestrator-client ; then
    echo "+ not found in PATH"
    if [ -f resources/bin/orchestrator-client ] ; then
      echo "found in resources/bin. Updating PATH"
      export PATH="$PATH:$(pwd)/resources/bin"
    else
      echo "orchestrator-client not found"
      exit 1
    fi
  fi
  echo "checking mysqladmin"
  if ! which mysqladmin ; then
    echo "+ not found in PATH"
    mysqladmin_file=$(find ~/opt/mysql -type f -name "mysqladmin" | head -n 1)
    if [ -n "$mysqladmin_file" ] ; then
      mysqladmin_path="$(dirname "$mysqladmin_file")"
      echo "found in $mysqladmin_path. Updating PATH"
      export PATH="$PATH:$mysqladmin_path"
    else
      echo "mysqladmin not found"
      exit 1
    fi
  fi
  export PATH="$PATH:$tests_path"
  if ! which test-retry ; then
    echo "test-retry not found"
    exit 1
  fi
  if ! which test-refresh-connections ; then
    echo "test-refresh-connections not found"
    exit 1
  fi
}

test_step() {
  local test_path
  test_path="$1"

  local test_name
  test_name="$2"

  local test_step_name
  test_step_name="$3"

  echo "Testing: $test_name/$test_step_name"

  if [ -f $test_path/config.json ] ; then
    echo "- applying configuration: $test_path/config.json"
    real_config_path="$(realpath $test_path/config.json)"
    orchestrator-client -c api -path "reload-configuration?config=$real_config_path" | jq -r '.Code'
  fi

  if [ -f $test_path/setup ] ; then
    bash $test_path/setup 1> $setup_teardown_logfile 2>&1
    if [ $? -ne 0 ] ; then
      echo "ERROR $test_name/$test_step_name setup failed"
      cat $setup_teardown_logfile
      return 1
    fi
  fi

  if [ -f $test_path/run ] ; then
    bash $test_path/run 1> $test_outfile 2> $test_logfile
    execution_result=$?
  elif [ -f $test_path/skip_run ] ; then
    echo "+ skipping run"
  else
    echo "ERROR: neither 'run' nor 'skip_run' found in $test_path"
    return 1
  fi

  if [ -f $test_path/teardown ] ; then
    bash $test_path/teardown 1> $setup_teardown_logfile 2>&1
    if [ $? -ne 0 ] ; then
      echo "ERROR $test_name/$test_step_name teardown failed"
      cat $setup_teardown_logfile
      return 1
    fi
  fi
  if [ -f $test_path/teardown_redeploy ] ; then
    bash $tests_path/deploy-replication || return 1
  fi

  if [ -f $test_path/run ] ; then
    if [ -f $test_path/expect_failure ] ; then
      if [ $execution_result -eq 0 ] ; then
        echo "ERROR $test_name/$test_step_name execution was expected to exit on error but did not. cat $test_logfile"
        echo "---"
        cat $test_logfile
        echo "---"
        return 1
      fi
      if [ -s $test_path/expect_failure ] ; then
        # 'expect_failure' file has content. We expect to find this content in the log.
        expected_error_message="$(cat $test_path/expect_failure)"
        if grep -q "$expected_error_message" $test_logfile ; then
          return 0
        fi
        echo "ERROR $test_name/$test_step_name execution was expected to exit with error message '${expected_error_message}' but did not. cat $test_logfile"
        echo "---"
        cat $test_logfile
        echo "---"
        return 1
      fi
      # 'expect_failure' file has no content. We generally agree that the failure is correct
      return 0
    fi

    if [ $execution_result -ne 0 ] ; then
      echo "ERROR $test_name/$test_step_name execution failure. cat $test_logfile"
      echo "---"
      cat $test_logfile
      echo "---"
      return 1
    fi

    if [ -f $test_path/expect_output ] ; then
      diff -b $test_path/expect_output $test_outfile > $test_diff_file
      diff_result=$?
      if [ $diff_result -ne 0 ] ; then
        echo "ERROR $test_name/$test_step_name output does not match expect_output"
        echo "---"
        cat $test_diff_file
        echo "---"
        return 1
      fi
    fi
  fi # [ -f $test_path/run ]

  return 0
}

test_single() {
  local test_name="$1"

  bash $tests_path/setup 1> $setup_teardown_logfile 2>&1
  if [ $? -ne 0 ] ; then
    echo "ERROR global setup failed"
    cat $setup_teardown_logfile
    return 1
  fi

  # test steps:
  find "$tests_path/$test_name" ! -path . -type d -mindepth 1 -maxdepth 1 | sort | cut -d "/" -f 5 | while read test_step_name ; do
    [ "$test_step_name" == "." ] && continue
    test_step "$tests_path/$test_name/$test_step_name" "$test_name" "$test_step_name"
    if [ $? -ne 0 ] ; then
      echo "+ FAIL"
      bash $tests_path/debug_dump
      return 1
    fi
    echo "+ pass"
  done || return 1

  # test main step:
  test_step "$tests_path/$test_name" "$test_name" "main"
  if [ $? -ne 0 ] ; then
    echo "+ FAIL"
    bash $tests_path/debug_dump
    return 1
  fi
  echo "+ pass"

  bash $tests_path/teardown 1> $setup_teardown_logfile 2>&1
  if [ $? -ne 0 ] ; then
    echo "ERROR global teardown failed"
    cat $setup_teardown_logfile
    return 1
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
  echo "# marking test as successful"
}

test_listed_as_successful() {
  local check_test_name="$1"
  cat $tests_successful_file | egrep -q "^$check_test_name\$"
}

test_listed_as_failed() {
  local check_test_name="$1"
  cat $tests_failed_file | egrep -q "^$check_test_name\$"
}

test_listed_as_attempted() {
  local check_test_name="$1"
  if test_listed_as_successful $check_test_name ; then
    return 0
  fi
  if test_listed_as_failed $check_test_name ; then
    return 0
  fi
  return 1
}

test_listed_as_failed() {
  local check_test_name="$1"
  cat $tests_failed_file | egrep -q "^$check_test_name\$"
}

should_attempt_test() {
  local check_test_name="$1"
  local force_test_pattern="${2:-}"

  if test_listed_as_successful $check_test_name ; then
    return 1
  fi
  if test_listed_as_failed $check_test_name ; then
    return 1
  fi
  if [ "$force_test_pattern" != "." ] ; then
    return 0
  fi
  # iterate dependencies
  (cat ${tests_path}/${check_test_name}/depends-on 2> /dev/null || echo -n "") | while read dependency ; do
    if test_listed_as_failed $dependency ; then
      # dependency has failed; fail this test, too
      echo "$check_test_name" >> $tests_failed_file
      return 1
    fi
  done
  (cat ${tests_path}/${check_test_name}/depends-on 2> /dev/null || echo -n "") | while read dependency ; do
    if ! test_listed_as_successful $dependency ; then
      # dependency hasn't been successful (yet?)
      return 1
    fi
  done
  return 0
}

test_all() {
  local test_pattern="${1:-.}"

  echo "." > $tests_todo_file
  while [ -s $tests_todo_file ] ; do
    echo -n > $tests_todo_file

    find $tests_path ! -path . -type d -mindepth 1 -maxdepth 1 | xargs ls -td1 | cut -d "/" -f 4 | egrep "$test_pattern" | while read test_name ; do
      if ! test_listed_as_attempted "$test_name" ; then
        echo "# test $test_name hasn't been executed before"
        echo "$test_name" >> $tests_todo_file
      fi
      if should_attempt_test "$test_name" "$test_pattern" ; then
        test_single "$test_name"
        if [ $? -eq 0 ] ; then
          echo "$test_name" >> $tests_successful_file
        else
          echo "$test_name" >> $tests_failed_file
          exit 1
        fi
      else
        echo "# should not attempt $test_name"
      fi
    done || return 1
    echo "# Iterated: $(cat $tests_todo_file)"
  done
}

main() {
  check_environment
  test_all "$@"
}

main "$@"
