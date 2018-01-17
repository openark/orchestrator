#!/bin/bash

# Local integration tests. To be used by CI.
# See https://github.com/github/orchestrator/tree/doc/local-tests.md
#

# Usage: localtests/test/sh [mysql|sqlite] [filter]
# By default, runs all tests. Given filter, will only run tests matching given regep

tests_path=$(dirname $0)
test_logfile=/tmp/orchestrator-test.log
test_outfile=/tmp/orchestrator-test.out
test_diff_file=/tmp/orchestrator-test.diff
test_query_file=/tmp/orchestrator-test.sql
test_config_file=/tmp/orchestrator.conf.json
orchestrator_binary=/tmp/orchestrator-test
exec_command_file=/tmp/orchestrator-test.bash
db_type=""
sqlite_file="/tmp/orchestrator.db"

function run_queries() {
  queries_file="$1"

  if [ "$db_type" == "sqlite" ] ; then
    cat $queries_file |
      sed -e "s/last_checked - interval 1 minute/datetime('last_checked', '-1 minute')/g" |
      sed -e "s/current_timestamp + interval 1 minute/datetime('now', '+1 minute')/g" |
      sqlite3 $sqlite_file
  else
    # Assume mysql
    mysql --default-character-set=utf8mb4 test -ss < $queries_file
  fi
}

check_db() {
  echo "select 1;" > $test_query_file
  query_result="$(run_queries $test_query_file)"
  if [ "$query_result" != "1" ] ; then
    echo "Cannot execute queries"
    exit 1
  fi
  echo "- check_db OK"
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
  run_queries $tests_path/create-per-test.sql

  echo_dot
  if [ -f $tests_path/$test_name/create.sql ] ; then
    run_queries $tests_path/$test_name/create.sql
  fi

  extra_args=""
  if [ -f $tests_path/$test_name/extra_args ] ; then
    extra_args=$(cat $tests_path/$test_name/extra_args)
  fi
  #
  cmd="$orchestrator_binary \
    --config=${test_config_file}
    --debug \
    --stack \
    ${extra_args[@]}"
  echo_dot
  echo $cmd > $exec_command_file
  echo_dot
  bash $exec_command_file 1> $test_outfile 2> $test_logfile

  execution_result=$?

  if [ -f $tests_path/$test_name/destroy.sql ] ; then
    run_queries $tests_path/$test_name/destroy.sql
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

build_binary() {
  echo "Building"
  go build -o $orchestrator_binary go/cmd/orchestrator/main.go
}


deploy_internal_db() {
  echo "Deploying db"
  cmd="$orchestrator_binary \
    --config=${test_config_file}
    --debug \
    --stack \
    -c redeploy-internal-db"
  echo_dot
  echo $cmd > $exec_command_file
  echo_dot
  bash $exec_command_file 1> $test_outfile 2> $test_logfile
  if [ $? -ne 0 ] ; then
    echo "ERROR deploy internal db failed"
    return 1
  fi
  echo "- deploy_internal_db result: $?"
}

generate_config_file() {
  cp ${tests_path}/orchestrator.conf.json ${test_config_file}
  sed -i -e "s/backend-db-placeholder/${db_type}/g" ${test_config_file}
  sed -i -e "s^sqlite-data-file-placeholder^${sqlite_file}^g" ${test_config_file}
  sed -i -e "s^mysql-user-placeholder^${MYSQL_USER:-}^g" ${test_config_file}
  echo "- generate_config_file OK"
}

test_all() {
  deploy_internal_db
  if [ $? -ne 0 ] ; then
    echo "ERROR deploy failed"
    return 1
  fi
  echo "- deploy_internal_db OK"

  test_pattern="${1:-.}"
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

test_db() {
  db_type="$1"
  echo "### testing via $db_type"
  generate_config_file
  check_db
  test_all ${@:2}
  if [ $? -ne 0 ] ; then
    echo "test_db failed"
    return 1
  fi
  echo "- done testing via $db_type"
}

main() {
  build_binary
  if [ $? -ne 0 ] ; then
    echo "ERROR build failed"
    return 1
  fi

  test_dbs=""
  if [ "$1" == "mysql" ] ; then
    test_dbs="mysql"
    shift
  elif [ "$1" == "sqlite" ] ; then
    test_dbs="sqlite"
    shift
  fi
  test_dbs=${test_dbs:-"mysql sqlite"}
  for db in $(echo $test_dbs) ; do
    test_db $db "$@"
    if [ $? -ne 0 ] ; then
      echo "+ FAIL: test_db $db $@"
      return 1
    fi
  done
}

main "$@"
