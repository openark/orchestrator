#!/bin/bash
set -xeuo pipefail
# Install orchestrator
if [[ -e /etc/redhat-release ]]; then
  rpm -i /tmp/orchestrator-release/orchestrator*.rpm
fi

if [[ -e /etc/debian_version ]]; then
  dpkg -i /tmp/orchestrator-release/orchestrator*.deb
fi

if [[ -e /orchestrator/vagrant/.sqlite ]]; then
  cp -fv /usr/local/orchestrator/orchestrator-sample-sqlite.conf.json /etc/orchestrator.conf.json
else
  cp -fv /usr/local/orchestrator/orchestrator-sample.conf.json /etc/orchestrator.conf.json
fi

if [[ -e /etc/redhat-release ]]; then

  /sbin/chkconfig orchestrator on
  /sbin/service orchestrator start

elif [[ -e /etc/debian_version ]]; then

  update-rc.d orchestrator defaults
  /usr/sbin/service orchestrator start

fi

echo '* * * * * root /usr/bin/orchestrator -c discover -i db1' > /etc/cron.d/orchestrator-discovery

# Discover instances
/usr/bin/orchestrator --verbose --debug --stack -c discover -i localhost
