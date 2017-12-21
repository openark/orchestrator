#!/bin/bash

# Install orchestrator
rpm -i /tmp/orchestrator-release/orchestrator*.rpm

if [[ -e /etc/redhat-release ]]; then

  /sbin/chkconfig orchestrator on
  cp /usr/local/orchestrator/orchestrator-sample.conf.json /etc/orchestrator.conf.json
  /sbin/service orchestrator start

elif [[ -e /etc/debian_version ]]; then

  update-rc.d orchestrator defaults
  cp /usr/local/orchestrator/orchestrator-sample.conf.json /etc/orchestrator.conf.json
  /usr/sbin/service orchestrator start

fi

echo '* * * * * root /usr/bin/orchestrator -c discover -i db1' > /etc/cron.d/orchestrator-discovery

# Discover instances
/usr/bin/orchestrator -c discover -i localhost
