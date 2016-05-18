if [[ -e /etc/debian_version ]]; then
  sudo cp /orchestrator/vagrant/db1-my.cnf /etc/mysql/my.cnf
  sudo /etc/init.d/mysql restart
fi

/usr/bin/mysql -uroot -ss -e 'GRANT REPLICATION SLAVE ON *.* TO "repl"@"192.168.57.%" IDENTIFIED BY "vagrant_repl"'
/usr/bin/mysql -uroot -ss -e 'CHANGE MASTER TO MASTER_HOST="192.168.57.202", MASTER_USER="repl", MASTER_PASSWORD="vagrant_repl", MASTER_CONNECT_RETRY=10, MASTER_RETRY_COUNT=36'
/usr/bin/mysql -uroot -ss -e 'START SLAVE'
