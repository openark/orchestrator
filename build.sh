#!/bin/bash

# Simple packaging of orchestrator
#
# Requires fpm: https://github.com/jordansissel/fpm
#

release_version="1.4.243"
release_dir=/tmp/orchestrator-release
release_files_dir=$release_dir/orchestrator
release_files_cli_dir=$release_dir/orchestrator-cli
rm -rf $release_dir/*
mkdir -p $release_files_dir
mkdir -p $release_files_dir/usr/local
mkdir -p $release_files_dir/etc/init.d
mkdir -p $release_files_cli_dir
mkdir -p $release_files_cli_dir/usr/bin

cd  $(dirname $0)
for f in $(find . -name "*.go"); do go fmt $f; done

rsync -av ./resources $release_files_dir/usr/local/orchestrator/
rsync -av ./conf $release_files_dir/usr/local/orchestrator/
for f in $release_files_dir/usr/local/orchestrator/conf/*; do mv $f $f.sample; done
cp etc/init.d/orchestrator.bash $release_files_dir/etc/init.d/orchestrator
chmod +x $release_files_dir/etc/init.d/orchestrator

go build -o $release_files_dir/usr/local/orchestrator/orchestrator go/cmd/orchestrator/main.go

if [[ $? -ne 0 ]] ; then
	exit 1
fi

cd $release_dir
# tar packaging
tar cfz orchestrator-"${release_version}".tar.gz ./orchestrator
# rpm packaging -- full package
fpm -v "${release_version}" -f -s dir -t rpm -n orchestrator -C $release_dir/orchestrator --prefix=/ .
fpm -v "${release_version}" -f -s dir -t deb -n orchestrator -C $release_dir/orchestrator --prefix=/ .

cp $release_files_dir/usr/local/orchestrator/orchestrator $release_files_cli_dir/usr/bin
cd $release_dir
# rpm packaging -- executable only
fpm -v "${release_version}" -f -s dir -t rpm -n orchestrator-cli -C $release_dir/orchestrator-cli --prefix=/ .
fpm -v "${release_version}" -f -s dir -t deb -n orchestrator-cli -C $release_dir/orchestrator-cli --prefix=/ .

echo "---"
echo "Done. Find releases in $release_dir"
