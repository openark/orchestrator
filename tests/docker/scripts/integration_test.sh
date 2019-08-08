#!/bin/sh
#
# Ugly wrapper around tests/integration/test.sh but using a docker mysql image
# for testing.
#
# Call this from the git root directory tests/integration/test.sh

# Horrible hack to figure out directories etc.
MYBASEDIR=$(dirname $0)
MYBASEDIR=$MYBASEDIR/../root
if ! echo "$MYBASEDIR" | grep -q ^/; then
	MYBASEDIR=$PWD/$MYBASEDIR
fi
BASEDIR=${BASEDIR:-$MYBASEDIR}

echo $BASEDIR

echo "Starting up docker mysql image needed for testing..."
BASEDIR=$BASEDIR tests/docker/scripts/mysql start

echo "Running integration tests using docker mysql image..."
tests/integration/test.sh -d tests/docker/scripts/dot_my_cnf

echo "Stopping docker mysql image needed for testing..."
BASEDIR=$BASEDIR tests/docker/scripts/mysql stop
