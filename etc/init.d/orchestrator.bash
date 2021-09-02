#!/bin/bash
# orchestrator daemon
# chkconfig: 345 20 80
# description: orchestrator daemon
# processname: orchestrator
#
### BEGIN INIT INFO
# Provides:          orchestrator
# Required-Start:    $local_fs $syslog
# Required-Stop:     $local_fs $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Start orchestrator daemon
# Description:       Start orchestrator daemon
### END INIT INFO


# Script credit: http://werxltd.com/wp/2012/01/05/simple-init-d-script-template/

DAEMON_PATH="/usr/local/orchestrator"

DAEMON=orchestrator
DAEMONOPTS="--verbose http"

NAME=orchestrator
DESC="orchestrator: MySQL replication management and visualization"
PIDFILE=/var/run/$NAME.pid
SCRIPTNAME=/etc/init.d/$NAME

# Limit the number of file descriptors (and sockets) used by
# orchestrator.  This setting should be fine in most cases but a
# large busy environment may # reach this limit. If exceeded expect
# to see errors of the form:
#   2017-06-12 02:33:09 ERROR dial tcp 10.1.2.3:3306: connect: cannot assign requested address
# To avoid touching this script you can use /etc/orchestrator_profile
# to increase this limit.
ulimit -n 16384

# initially noop but can adjust according by modifying orchestrator_profile
# - see https://github.com/openark/orchestrator/issues/227 for more details.
post_start_daemon_hook () {
	# by default do nothing
	:
}

# Start the orchestrator daemon in the background
start_daemon () {
	# start up daemon in the background
	$DAEMON_PATH/$DAEMON $DAEMONOPTS >> /var/log/${NAME}.log 2>&1 &
	# collect and print PID of started process
	echo $!
	# space for optional processing after starting orchestrator
	# - redirect stdout to stderro to prevent this corrupting the pid info
	post_start_daemon_hook 1>&2
}

# This files can be used to inject pre-service execution
# scripts, such as exporting variables or whatever. It's yours!
[ -f /etc/default/orchestrator ] && . /etc/default/orchestrator
[ -f /etc/orchestrator_profile ] && . /etc/orchestrator_profile
[ -f /etc/profile.d/orchestrator ] && . /etc/profile.d/orchestrator

case "$1" in
  start)
    printf "%-50s" "Starting $NAME..."
    cd $DAEMON_PATH
    PID=$(start_daemon)
    #echo "Saving PID" $PID " to " $PIDFILE
    if [ -z $PID ]; then
      printf "%s\n" "Fail"
      exit 1
    elif [ -z "$(ps axf | awk '{print $1}' | grep ${PID})" ]; then
      printf "%s\n" "Fail"
      exit 1
    else
      echo $PID > $PIDFILE
      printf "%s\n" "Ok"
    fi
  ;;
  status)
    printf "%-50s" "Checking $NAME..."
    if [ -f $PIDFILE ]; then
      PID=$(cat $PIDFILE)
      if [ -z "$(ps axf | awk '{print $1}' | grep ${PID})" ]; then
        printf "%s\n" "Process dead but pidfile exists"
        exit 1
      else
        echo "Running"
      fi
    else
      printf "%s\n" "Service not running"
      exit 1
    fi
  ;;
  stop)
    printf "%-50s" "Stopping $NAME"
    PID=$(cat $PIDFILE)
    cd $DAEMON_PATH
    if [ -f $PIDFILE ]; then
      kill -TERM $PID
      rm -f $PIDFILE
      # Wait for orchestrator to stop otherwise restart may fail.
      # (The newly restarted process may be unable to bind to the
      # currently bound socket.)
      while ps -p $PID >/dev/null 2>&1; do
        printf "."
        sleep 1
      done
      printf "\n"
      printf "Ok\n"
    else
      printf "%s\n" "pidfile not found"
      exit 1
    fi
  ;;
  restart)
    $0 stop
    $0 start
  ;;
  reload)
    PID=$(cat $PIDFILE)
    cd $DAEMON_PATH
    if [ -f $PIDFILE ]; then
      kill -HUP $PID
      printf "%s\n" "Ok"
    else
      printf "%s\n" "pidfile not found"
      exit 1
    fi
	;;
  *)
    echo "Usage: $0 {status|start|stop|restart|reload}"
    exit 1
esac
