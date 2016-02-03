#!/bin/bash
# orchestrator daemon
# chkconfig: 345 20 80
# description: orchestrator daemon
# processname: orchestrator


# Script credit: http://werxltd.com/wp/2012/01/05/simple-init-d-script-template/

DAEMON_PATH="/usr/local/orchestrator"

DAEMON=orchestrator
DAEMONOPTS="--verbose http"

NAME=orchestrator
DESC="orchestrator: MySQL replication management and visualization"
PIDFILE=/var/run/$NAME.pid
SCRIPTNAME=/etc/init.d/$NAME

ulimit -n 16384

# The file /etc/orchestrator_profile can be used to inject pre-service execution
# scripts, such as exporting variables or whatever. It's yours!
[ -f /etc/orchestrator_profile ] && . /etc/orchestrator_profile

case "$1" in
  start)
    printf "%-50s" "Starting $NAME..."
    cd $DAEMON_PATH
    PID=$(./$DAEMON $DAEMONOPTS >> /var/log/${NAME}.log 2>&1 & echo $!)
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
