/*
   Copyright 2017 GitHub Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package remote

var GetRelayLogContentsScript = `#!/bin/bash
#
# We will magically set these variables:
FIRST_RELAYLOG_FILE="$MAGIC_FIRST_RELAYLOG_FILE"
START_POSITION="$MAGIC_START_POSITION"
STOP_POSITION="$MAGIC_STOP_POSITION"
[ "$STOP_POSITION" == "0" ] && STOP_POSITION=""

set -e

header_size=""
datadir=""
relaylog_index_file=""
relay_logs=""

bootstrap() {
  datadir=$(find /etc -name "my.cnf" | xargs grep datadir | head -1 | cut -d"=" -f2 | xargs echo)
  [ -z "$datadir" ] && datadir=$(ps aux | grep mysqld | egrep -o '[-][-]datadir=[^ ]+' | cut -d"=" -f2)
  if [ -z "$datadir" ] ; then
    echo "Cannot find @@datadir"
    exit 1
  fi

  relaylog_index_file=$(find $datadir -name "*relay*.index" | head -1)
  if [ -z "$relaylog_index_file" ] ; then
    echo "Cannot find relaylog index file"
    exit 1
  fi
}

relaylogs_starting_at() {
  starting_relay_log=$(basename "$1")

  relay_logs_file=$(mktemp)
  basedir=$(dirname $relaylog_index_file)
  cat $relaylog_index_file | awk "/$starting_relay_log/,0" | while read f ; do
    echo "${basedir}/${f}" >> $relay_logs_file
  done
  relay_logs=$(cat $relay_logs_file)
}

binlog_header_size() {
  binlog_file="$1"
	# magic header
	# There are the first 4 bytes, and then there's also the first entry (the format-description).
	# We need both from the first log file.
	# Typically, the format description ends at pos 120, but let's verify...

  header_size=""
	header_size=$(mysqlbinlog $binlog_file --start-position=4 | head | egrep -o 'end_log_pos [^ ]+' | head -1 | awk '{print $2}')
}


relaylog_tail() {
  starting_relay_log=$(basename "$FIRST_RELAYLOG_FILE")
  start_position="$START_POSITION"
  contents_file=$(mktemp)

  [ "$start_position" == "0" ] && start_position=""

  relaylogs_starting_at $starting_relay_log

  is_first_relaylog=1
  last_relaylog=$(echo "$relay_logs" | tail -1)

  for relay_log in $relay_logs ; do
    binlog_header_size $relay_log
    stop_position=$(wc -c $relay_log)
    if [ "$relay_log" == "$last_relaylog" ] ; then
      if [ ! -z "$STOP_POSITION" ] ; then
        stop_position="$STOP_POSITION"
      fi
    fi
    # header_size variable is now populated
    if [ $is_first_relaylog -eq 1 ] ; then
      # First relaylog file. We only get the header from the first file.
      # Then, we also filter by --start-position
      cat $relay_log | head -c$header_size >> $contents_file
      [ -z "$start_position" ] && start_position="$header_size"
      # _tail_ command is unfortunately 1-based, which is why we ++
      ((start_position++))
      cat $relay_log | head -c $stop_position | tail -c+$start_position >> $contents_file
    else
      # Skip header
      # _tail_ command is unfortunately 1-based, which is why we ++
      ((header_size++))
      cat $relay_log | head -c $stop_position | tail -c+$header_size >> $contents_file
    fi
    is_first_relaylog=0
  done

  cat $contents_file | gzip | base64
}

main() {
  bootstrap
  relaylog_tail
}

main
`

var ApplyRelayLogContentsScript = `#!/bin/bash
#
# Usage:
#
# apply-relaylog-contents [path-to-mysql]
# - [path-to-mysql] would be a mysql command that has full privileges (including GRANT privileges).
#   for example: "mysql --default-file=/root/.my.cnf"
#   Notice this should be quoted as one argument to the program

mysql_command="$MAGIC_MYSQL_COMMAND"
mysql_command="${mysql_command:-mysql}"
contents_file="$MAGIC_CONTENTS_FILE"
contents_file="${contents_file:-"-"}"


set -e

apply_relaylog() {
  binlog_file=$(mktemp)

  cat "$contents_file" | base64 --decode | gunzip > $binlog_file
  mysqlbinlog $binlog_file | $mysql_command
}

main() {
  apply_relaylog
}

main
`
