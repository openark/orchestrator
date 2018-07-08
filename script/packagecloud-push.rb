# Pushes all deb and rpm files from ./repos to PackageCloud.

packagecloud_user = ENV["PACKAGECLOUD_USER"] || "github"
packagecloud_token = ENV["PACKAGECLOUD_TOKEN"] || begin
  puts "PACKAGECLOUD_TOKEN env required"
  exit 1
end

require "json"

packagecloud_ruby_minimum_version = "1.0.4"
begin
  gem "packagecloud-ruby", ">=#{packagecloud_ruby_minimum_version}"
  require "packagecloud"
  puts "Using packagecloud-ruby:#{Gem.loaded_specs["packagecloud-ruby"].version}"
rescue LoadError
  puts "Requires packagecloud-ruby >=#{packagecloud_ruby_minimum_version}"
  puts %(gem install packagecloud-ruby)
  exit 1
end

credentials = Packagecloud::Credentials.new(packagecloud_user, packagecloud_token)
$client = Packagecloud::Client.new(credentials)

# matches package directories built by docker to one or more packagecloud distros
# https://packagecloud.io/docs#os_distro_version
$distro_name_map = {
  "centos/6" => %w(
    el/6
  ),
  "centos/7" => %w(
    el/7
    fedora/22
    fedora/23
    fedora/24
  ),
  "debian/7" => %w(
    debian/wheezy
    ubuntu/precise
  ),
  "debian/8" => %w(
    debian/jessie
    linuxmint/sarah
    linuxmint/rebecca
    linuxmint/rafaela
    linuxmint/rosa
    linuxmint/sarah
    linuxmint/serena
    linuxmint/sonya
    linuxmint/sylvia
    ubuntu/trusty
    ubuntu/vivid
    ubuntu/wily
    ubuntu/xenial
    ubuntu/yakkety
    ubuntu/zesty
    ubuntu/artful
    ubuntu/bionic
  ),
  "debian/9" => %w(
    debian/stretch
  ),
  "debian/10" => %w(
    debian/buster
  ),
  "debian/11" => %w(
    debian/bullseye
  ),
  "debian/12" => %w(
    debian/bookworm
  ),
}

# caches distro id lookups
$distro_id_map = {}

def distro_names_for(filename)
  result = []
  $distro_name_map.each do |pattern, distros|
    if filename.include?(".deb") and pattern .include?("debian")
      result.concat distros
    end

    if filename.include?(".rpm") and pattern .include?("centos")
      if filename.include?("centos6") and pattern.include?("centos/6")
        result.concat distros
      end
      if !filename.include?("centos6") and !pattern.include?("centos/6")
        result.concat distros
      end
    end

    if filename.include?(pattern)
      result.concat distros
    end
  end

  if result.empty?
    raise "no distro for #{filename.inspect}"
  end
  return result
end

package_files = Dir.glob("/tmp/orchestrator-release/**/*.rpm") + Dir.glob("/tmp/orchestrator-release/**/*.deb")
package_files.each do |full_path|
  next if full_path =~ /repo-release/
  pkg = Packagecloud::Package.new(:file => full_path)
  distro_names = distro_names_for(full_path)
  distro_names.map do |distro_name|
    distro_id = $distro_id_map[distro_name] ||= $client.find_distribution_id(distro_name)
    if !distro_id
      raise "no distro id for #{distro_name.inspect}"
    end

    puts "pushing #{full_path} to #{$distro_id_map.key(distro_id).inspect}"
    result = $client.put_package("orchestrator", pkg, distro_id)
    result.succeeded || begin
      raise "packagecloud put_package failed, error: #{result.response}"
    end
  end
end

package_files.each do |full_path|
  next if full_path.include?("SRPM") || full_path.include?("i386") || full_path.include?("i686")
  next unless (full_path =~ /\/orchestrator[-|_]\d/) or (full_path =~ /\/orchestrator-cli[-|_]\d/) or (full_path =~ /\/orchestrator-client[-|_]\d/)
  os, distro = case full_path
  when /debian\/7/ then ["Debian 7", "debian/wheezy"]
  when /debian\/8/ then ["Debian 8", "debian/jessie"]
  when /centos\/6/ then ["RPM RHEL 6/CentOS 6", "el/6"]
  when /centos\/7/ then ["RPM RHEL 7/CentOS 7", "el/7"]
  end

  next unless os

  puts "[#{os}](https://packagecloud.io/#{packagecloud_user}/orchestrator/packages/#{distro}/#{File.basename(full_path)}/download)"
end
