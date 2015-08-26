#!/bin/bash

if [[ -e /etc/redhat-release ]]; then
  # Percona's Yum Repository
  sudo yum --nogpgcheck -y install http://www.percona.com/downloads/percona-release/redhat/0.1-3/percona-release-0.1-3.noarch.rpm

  # No reason not to install this stuff in all the places :)
  sudo yum --nogpgcheck -y install Percona-Server-server-56 Percona-Server-shared-56 Percona-Server-client-56 Percona-Server-shared-compat
  sudo yum --nogpgcheck -y install percona-toolkit percona-xtrabackup

  # All the project dependencies to build
  sudo yum --nogpgcheck -y install ruby-devel gcc rpm-build git
  gem install fpm

  # Go (via EPEL)
  sudo yum --nogpgcheck -y install epel-release
  sudo yum --nogpgcheck -y install golang
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

cat <<-EOF > /home/vagrant/.ssh/id_rsa
  -----BEGIN RSA PRIVATE KEY-----
  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
  -----END RSA PRIVATE KEY-----
EOF

cat <<-EOF > /home/vagrant/.ssh/id_rsa.pub
  ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDqIhZSWQZq1Tw/caTZUtZRrpRv9zEhNVSJReLnlCseRYaMCCdf+ZLQylPgEf8khgCj9Z0Td8xwus0pYyME7GmHf4/Bpl7i6RFXto84ZyEwgkwC+OMqZh4g8JsZJkd/Fhf0JS7UaZqiL/9PbyCnZ0b2dEji6sXcLN8cZA9qUn1j1MccRozDsUDwQRWypbszgarrTJSwRzjmh2q+L9iAA91qqt0fDJS3OCZnFtF07ZfSv2vAyrZuhnDUg4YRzt8ol4uX0oVGE+Ihb/uchVa/EJFhrDa95wH4NUUIwc6j4D7p4gBiZbjBZsokYUoqhJpnSTIQBZa8FTo3VgR2KFaCz3iDj3DsIW8IPdwTtcHoGQSgZYah1fzZWb25VOlONqfX+ugCbeJYx/oOV2J4wqtQK0QK5Gl8nF293hs96ZCethxvf2AfbZCGPJmp2n5jTTXP4zw1xPKaeQCjmYt49xWY/r78Few86x4W6edw/On+oLVCh9bfBgIwK6Pzy6fvz08IT5G2cI37JbKCLpgt5qUeuHAOxhpvf73Cdi0iUWZZpBwA/GZdMPUaZx9803yqpEO0L0h2N3yQhdk5Ax1/OFtgu7c+0K6g3N33CW05gv4nzwTuuMRtFT7PLeyaHKr+G7gmkRl1mqJDrCRTRXxNuUqJNXOBdLH1ijhACPmIhMcGRKUblQ== vagrant@orchestrator
EOF

cat <<EOF > /home/vagrant/.ssh/config
Host admin
  User vagrant
  IdentifyFile /home/vagrant/.ssh/id_rsa
Host db1
  User vagrant
  IdentifyFile /home/vagrant/.ssh/id_rsa
Host db2
  User vagrant
  IdentifyFile /home/vagrant/.ssh/id_rsa
Host db3
  User vagrant
  IdentifyFile /home/vagrant/.ssh/id_rsa
Host db4
  User vagrant
  IdentifyFile /home/vagrant/.ssh/id_rsa
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
