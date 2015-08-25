sudo rm -rf /etc/my.cnf
sudo cp /orchestrator/vagrant/db3-my.cnf /etc/my.cnf

sudo service mysql start

/usr/bin/mysql -uroot -ss -e 'GRANT REPLICATION SLAVE ON *.* TO "repl"@"192.168.57.%" IDENTIFIED BY "vagrant_repl"'
/usr/bin/mysql -uroot -ss -e 'CHANGE MASTER TO MASTER_HOST="192.168.57.201", MASTER_USER="repl", MASTER_PASSWORD="vagrant_repl"'
/usr/bin/mysql -uroot -ss -e 'START SLAVE'
