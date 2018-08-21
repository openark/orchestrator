#!/bin/sh
if [ ! -e /etc/orchestrator.conf.json ] ; then
cat <<EOF > /etc/orchestrator.conf.json
{
  "Debug": true,
  "ListenAddress": ":3000",
  "BackendDB": "sqlite3",
  "SQLite3DataFile": "/usr/local/orchestrator/orchestrator.db",
  "MySQLTopologyUser": "${ORC_TOPOLOGY_USER:-orchestrator}",
  "MySQLTopologyPassword": "${ORC_TOPOLOGY_PASSWORD:-orchestrator}",
}
EOF
fi

exec /usr/local/orchestrator/orchestrator http
