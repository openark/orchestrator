#!/bin/bash
set -xeuo pipefail
if [[ -e /etc/redhat-release ]]; then
  # Percona's Yum Repository
  yum -d 0 -y install http://www.percona.com/downloads/percona-release/redhat/0.1-4/percona-release-0.1-4.noarch.rpm epel-release

  # All the project dependencies to build plus some utilities
  # No reason not to install this stuff in all the places :)
  yum -d 0 -y install Percona-Server-server-56 Percona-Server-shared-56 Percona-Server-client-56 Percona-Server-shared-compat percona-toolkit percona-xtrabackup ruby-devel gcc rpm-build git vim-enhanced golang jq
  # newest versions of java aren't compatible with the installed version of ruby (1.8.7)
  gem install json --version 1.8.6
  # Pin to 1.4 due to 1.5 no longer working on EL6
  gem install fpm --version 1.4

  echo "PATH=$PATH:/usr/local/bin" | sudo tee -a /etc/environment
  export PATH="PATH=$PATH:/usr/local/bin"


  # Build orchestrator and orchestrator agent
  mkdir -p /home/vagrant/go/{bin,pkg,src} /tmp/orchestrator-release
  mkdir -p /home/vagrant/go/src/github.com/openark/orchestrator
  mount --bind /orchestrator /home/vagrant/go/src/github.com/openark/orchestrator

  # Build Orchestrator
  export GOPATH=/home/vagrant/go
  export GO15VENDOREXPERIMENT=1
  cd ${GOPATH}/src/github.com/openark/orchestrator
  /usr/bin/go get ./...
  ${GOPATH}/src/github.com/openark/orchestrator/build.sh
  chown -R vagrant.vagrant /home/vagrant /tmp/orchestrator-release

  # Setup mysql
  set +e
  /sbin/chkconfig mysql on
  /sbin/chkconfig mysqld on
  set -e

  if [[ -e "/orchestrator/vagrant/${HOSTNAME}-my.cnf" ]]; then
    rm -f /etc/my.cnf
    cp /orchestrator/vagrant/${HOSTNAME}-my.cnf /etc/my.cnf
  fi

  /sbin/service mysql start

elif [[ -e /etc/debian_version ]]; then
  sudo echo exit 101 > /usr/sbin/policy-rc.d
  sudo chmod +x /usr/sbin/policy-rc.d


  # Percona's Apt Repository
  sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1C4CBDCDCD2EFD2A 9334A25F8507EFA5
  echo "deb http://repo.percona.com/apt "$(lsb_release -sc)" main" | sudo tee /etc/apt/sources.list.d/percona.list
  sudo apt-get -y update
  sudo apt-get -y install debconf-utils
  echo "golang-go golang-go/dashboard boolean true" | sudo debconf-set-selections
  echo percona-server-server-5.6 percona-server-server/root_password password "" | sudo debconf-set-selections
  echo percona-server-server-5.6 percona-server-server/root_password_again password "" | sudo debconf-set-selections
  export DEBIAN_FRONTEND=noninteractive

  # No reason not to install this stuff in all the places :)
  #sudo apt-get -y install percona-server-server-5.6 percona-server-common-5.6 percona-server-client-5.6
  #sudo apt-get -y install percona-toolkit percona-xtrabackup

  # add the mysql community packages
  # from https://dev.mysql.com/doc/mysql-apt-repo-quick-guide/en/
  sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 5072E1F5 8C718D3B5072E1F5
  export CODENAME=$(/usr/bin/lsb_release -c | cut -f2)
  echo "deb http://repo.mysql.com/apt/ubuntu/ ${CODENAME} mysql-5.7" | sudo tee /etc/apt/sources.list.d/mysql.list
  apt-get -y update
  echo mysql-community-server mysql-community-server/root-pass password "" | sudo debconf-set-selections
  echo mysql-community-server mysql-community-server/re-root-pass password "" | sudo debconf-set-selections
  apt-get -y --force-yes install mysql-server
  chmod a+w /var/log

  # All the project dependencies to build
  sudo apt-get -y install ruby-dev gcc git rubygems rpm jq make
  # Jump though some hoops to get a non-decrepit version of golang
  sudo apt-get remove -y golang-go
  cd /tmp
  wget -c --quiet "https://redirector.gvt1.com/edgedl/go/go1.9.4.linux-amd64.tar.gz"
  sudo tar -C /usr/local -xzf go1.9.4.linux-amd64.tar.gz
  echo "PATH=$PATH:/usr/local/go/bin:/usr/local/bin" | sudo tee -a /etc/environment
  export PATH="PATH=$PATH:/usr/local/go/bin:/usr/local/bin"

  # newest versions of java aren't compatible with the installed version of ruby (1.8.7)
  gem install json --version 1.8.6
  gem install fpm --version 1.4

  # Build orchestrator and orchestrator agent
  mkdir -p /home/vagrant/go/{bin,pkg,src} /tmp/orchestrator-release
  mkdir -p /home/vagrant/go/src/github.com/openark/orchestrator
  mount --bind /orchestrator /home/vagrant/go/src/github.com/openark/orchestrator

  # Build Orchestrator
  export GOPATH=/home/vagrant/go
  export GO15VENDOREXPERIMENT=1
  cd ${GOPATH}/src/github.com/openark/orchestrator
  /usr/local/go/bin/go get ./...
  ${GOPATH}/src/github.com/openark/orchestrator/build.sh
  chown -R vagrant.vagrant /home/vagrant /tmp/orchestrator-release


  # Go
  sudo apt-get -y install golang-go

  update-rc.d mysql defaults
  /usr/sbin/service mysql start
fi

sudo mysql -e "grant all on *.* to 'root'@'localhost' identified by ''"
cat <<-EOF | mysql -u root
CREATE DATABASE IF NOT EXISTS orchestrator;
GRANT ALL PRIVILEGES ON orchestrator.* TO 'orc_client_user'@'%' IDENTIFIED BY 'orc_client_password';
GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO 'orc_client_user'@'%';
GRANT ALL PRIVILEGES ON orchestrator.* TO 'orc_client_user'@'localhost' IDENTIFIED BY 'orc_client_password';
GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO 'orc_client_user'@'localhost';
GRANT ALL PRIVILEGES ON orchestrator.* TO 'orc_server_user'@'localhost' IDENTIFIED BY 'orc_server_password';
EOF

cat <<-EOF >> /etc/hosts
  192.168.57.200   admin
  192.168.57.201   db1
  192.168.57.202   db2
  192.168.57.203   db3
  192.168.57.204   db4
EOF

if [[ -e /etc/redhat-release ]]; then
  set +e
  sudo service iptables stop
  set -e
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
