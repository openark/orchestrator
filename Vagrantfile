DEVELOPMENT_DIR = File.join(ENV['HOME'], 'Development')

ENV['VAGRANT_SERVER_URL'] = "https://atlas.hashicorp.com" if ENV['VAGRANT_SERVER_URL'].nil? || ENV['VAGRANT_SERVER_URL'].empty?
ENV['VAGRANT_DEFAULT_PROVIDER'] = 'virtualbox' if ENV['VAGRANT_DEFAULT_PROVIDER'].nil? || ENV['VAGRANT_DEFAULT_PROVIDER'].empty?
BOX = ENV['VAGRANT_BOX'].nil? || ENV['VAGRANT_BOX'].empty? ? 'nrel/CentOS-6.6-x86_64' : ENV['VAGRANT_BOX']

VAGRANTFILE_API_VERSION = "2"

system("
    if [[ #{ARGV[0]} = 'up' ]] && [[ ! -e 'vagrant/vagrant-ssh-key' ]]; then
      ssh-keygen -t rsa -b 768 -N '' -q -f vagrant/vagrant-ssh-key
    fi
")

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = BOX
  config.vm.box_download_insecure = true
  config.vm.box_check_update = false
  config.vm.provider 'virtualbox' do |vb|
    vb.customize ['modifyvm', :id, '--nictype1', 'virtio']
    vb.customize ['modifyvm', :id, '--nictype2', 'virtio']
  end

  config.vm.synced_folder '.', '/orchestrator', type: 'rsync',
    rsync__auto: true,
    rsync__exclude: ".git/"

  (0..4).each do |n|
    name = (n > 0 ? ("db" + n.to_s) : "admin")
    config.vm.define name do |db|
      db.vm.hostname = name
      db.vm.network "private_network", ip: "192.168.57.20" + n.to_s
      db.vm.provision "shell", path: "vagrant/base-build.sh"
      if name == "admin"
        db.vm.network "forwarded_port", guest:3000, host:3000
      end
    end
  end
end
