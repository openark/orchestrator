#!/bin/bash

if [[ -e /etc/redhat-release ]]; then
  # Percona's Yum Repository
  yum -d 0 -y install http://www.percona.com/downloads/percona-release/redhat/0.1-3/percona-release-0.1-3.noarch.rpm epel-release

  # All the project dependencies to build plus some utilities
  # No reason not to install this stuff in all the places :)
  yum -d 0 -y install Percona-Server-server-56 Percona-Server-shared-56 Percona-Server-client-56 Percona-Server-shared-compat percona-toolkit percona-xtrabackup ruby-devel gcc rpm-build git vim-enhanced golang
  # Pin to 1.4 due to 1.5 no longer working on EL6
  gem install fpm --version 1.4

  # Build orchestrator and orchestrator agent
  mkdir -p /home/vagrant/go/{bin,pkg,src} /tmp/orchestrator-release
  mkdir -p /home/vagrant/go/src/github.com/outbrain/orchestrator
  mount --bind /orchestrator /home/vagrant/go/src/github.com/outbrain/orchestrator

  # Build Orchestrator
  export GOPATH=/home/vagrant/go
  export GO15VENDOREXPERIMENT=1
  cd ${GOPATH}/src/github.com/outbrain/orchestrator
  /usr/bin/go get ./...
  ${GOPATH}/src/github.com/outbrain/orchestrator/build.sh
  chown -R vagrant.vagrant /home/vagrant /tmp/orchestrator-release

  # Setup mysql
  /sbin/chkconfig mysql on

  if [[ -e "/orchestrator/vagrant/${HOSTNAME}-my.cnf" ]]; then
    rm -f /etc/my.cnf
    cp /orchestrator/vagrant/${HOSTNAME}-my.cnf /etc/my.cnf
  fi

  /sbin/service mysql start
  cat <<-EOF | mysql -u root
  CREATE DATABASE IF NOT EXISTS orchestrator;
  GRANT ALL PRIVILEGES ON orchestrator.* TO 'orc_client_user'@'%' IDENTIFIED BY 'orc_client_password';
  GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO 'orc_client_user'@'%';
  GRANT ALL PRIVILEGES ON orchestrator.* TO 'orc_client_user'@'localhost' IDENTIFIED BY 'orc_client_password';
  GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO 'orc_client_user'@'localhost';
  GRANT ALL PRIVILEGES ON orchestrator.* TO 'orc_server_user'@'localhost' IDENTIFIED BY 'orc_server_password';
EOF

elif [[ -e /etc/debian_version ]]; then
  sudo echo exit 101 > /usr/sbin/policy-rc.d
  sudo chmod +x /usr/sbin/policy-rc.d
  # Percona's Apt Repository
  sudo apt-key adv --keyserver keys.gnupg.net --recv-keys 1C4CBDCDCD2EFD2A
  echo "deb http://repo.percona.com/apt "$(lsb_release -sc)" main" | sudo tee /etc/apt/sources.list.d/percona.list
  sudo apt-get -y update
  sudo apt-get -y install debconf-utils
  echo "golang-go golang-go/dashboard boolean true" | sudo debconf-set-selections
  echo "percona-server-server-5.6 percona-server-server/root_password password vagrant" | sudo debconf-set-selections
  echo "percona-server-server-5.6 percona-server-server/root_password_again password vagrant" | sudo debconf-set-selections
  DEBIAN_FRONTEND=noninteractive

  # No reason not to install this stuff in all the places :)
  sudo apt-get -y install percona-server-server-5.6 percona-server-common-5.6 percona-server-client-5.6
  sudo apt-get -y install percona-toolkit percona-xtrabackup

  # All the project dependencies to build
  sudo apt-get -y install ruby-dev gcc git rubygems rpm
  sudo gem install fpm

  # Go
  sudo apt-get -y install golang-go
fi

cat <<-EOF >> /etc/hosts
  192.168.57.200   admin
  192.168.57.201   db1
  192.168.57.202   db2
  192.168.57.203   db3
  192.168.57.204   db4
EOF

# Generated a random SSH keypair to be used by the vagrant user for convenience
mkdir -p /home/vagrant/.ssh

cp /orchestrator/vagrant/vagrant-ssh-key /home/vagrant/.ssh/id_rsa
cp /orchestrator/vagrant/vagrant-ssh-key.pub /home/vagrant/.ssh/id_rsa.pub

cat <<EOF > /home/vagrant/.ssh/config
Host admin
  User vagrant
  IdentityFile /home/vagrant/.ssh/id_rsa
Host db1
  User vagrant
  IdentityFile /home/vagrant/.ssh/id_rsa
Host db2
  User vagrant
  IdentityFile /home/vagrant/.ssh/id_rsa
Host db3
  User vagrant
  IdentityFile /home/vagrant/.ssh/id_rsa
Host db4
  User vagrant
  IdentityFile /home/vagrant/.ssh/id_rsa
EOF

chmod go-rwx /home/vagrant/.ssh/*
chown -R vagrant:vagrant /home/vagrant/.ssh

if [[ -e /etc/redhat-release ]]; then
  sudo service iptables stop
fi

if [[ $HOSTNAME == 'admin' ]]; then
  bash /orchestrator/vagrant/admin-build.sh
  if [[ -e /orchestrator/vagrant/admin-post-install.sh ]]; then
    bash /orchestrator/vagrant/admin-post-install.sh
  fi
else
  bash /orchestrator/vagrant/$HOSTNAME-build.sh

  if [[ -e /orchestrator/vagrant/db-post-install.sh ]]; then
    bash /orchestrator/vagrant/db-post-install.sh
  fi

  if [[ -e /orchestrator/vagrant/$HOSTNAME-post-install.sh ]]; then
    bash /orchestrator/vagrant/$HOSTNAME-post-install.sh
  fi
fi
