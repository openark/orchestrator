#!/bin/bash

# Install orchestrator
rpm -i /tmp/orchestrator-release/orchestrator*.rpm
/sbin/chkconfig orchestrator on
cp /usr/local/orchestrator/orchestrator-sample.conf.json /etc/orchestrator.conf.json
/sbin/service orchestrator start

echo '* * * * * root /usr/bin/orchestrator -c discover -i db1' > /etc/cron.d/orchestrator-discovery

# Discover instances
/usr/bin/orchestrator -c discover -i localhost