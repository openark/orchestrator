#!/bin/bash

# Simple packaging of orchestrator
#
# Requires fpm: https://github.com/jordansissel/fpm
#
set -e

RELEASE_VERSION="1.4.288"
TOPDIR=/tmp/orchestrator-release
export RELEASE_VERSION TOPDIR

usage() {
  echo
  echo "Usage: $0 [-t target ] [-a arch ] [-h] [-d]"
  echo "Options:"
  echo "-h Show this screen"
  echo "-t (linux|darwin) Target OS. Default:(linux)"
  echo "-a (amd64|386) Arch Default:(amd64)"
  echo "-d debug output"
  echo
}

function precheck() {
  local ok=0 # return err. so shell exit code

  if [[ ! -x "$( which fpm )" ]]; then
    echo "Please install fpm and ensure it is in PATH"
    ok=1
  fi

  if [[ ! -x "$( which rpmbuild )" ]]; then
    echo "rpmbuild not in PATH, rpm will not be built"
  fi

  if [[ -z "$GOROOT" || -z "$GOPATH" ]]; then
    echo "GOROOT or GOPATH not set"
    ok=1
  fi

  if [[ ! -x "$GOROOT/bin/go" ]]; then
    echo "go binary not found at $GOROOT/bin/go"
    ok=1
  fi

  return $ok
}

function setuptree() {
  local b

  mkdir -p $TOPDIR
rm -rf $TOPDIR/*
  b=$( mktemp -d $TOPDIR/orchestratorXXXXXX ) || return 1
  mkdir -p $b/orchestrator
  mkdir -p $b/orchestrator/usr/local/orchestrator/
  mkdir -p $b/orchestrator/etc/init.d
  mkdir -p $b/orchestrator-cli/usr/bin
  echo $b
}

function oinstall() {
  local builddir
  builddir="$1"

  cd  $(dirname $0)
  go fmt go/./...
  rsync -qa ./resources $builddir/orchestrator/usr/local/orchestrator/
  rsync -qa ./conf/*.sample $builddir/orchestrator/usr/local/orchestrator/
  cp etc/init.d/orchestrator.bash $builddir/orchestrator/etc/init.d/orchestrator
  chmod +x $builddir/orchestrator/etc/init.d/orchestrator
}

function package() {
  local target builddir packages
  target="$1"
  builddir="$2"

  cd $TOPDIR

  case $target in
    'linux')
      echo "Creating Linux Tar package"
      tar -C $builddir/orchestrator -czf $TOPDIR/orchestrator-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./

      echo "Creating Distro full packages"
      fpm -v "${RELEASE_VERSION}" --epoch 1 -f -s dir -t rpm -n orchestrator -C $builddir/orchestrator --prefix=/ .
      fpm -v "${RELEASE_VERSION}" --epoch 1 -f -s dir -t deb -n orchestrator -C $builddir/orchestrator --prefix=/ .

      cd $TOPDIR
      # rpm packaging -- executable only
      echo "Creating Distro cli packages"
      cp $builddir/orchestrator/usr/local/orchestrator/orchestrator $builddir/orchestrator-cli/usr/bin
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -t rpm -n orchestrator-cli -C $builddir/orchestrator-cli --prefix=/ .
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -t deb -n orchestrator-cli -C $builddir/orchestrator-cli --prefix=/ .
      ;;
    'darwin')
      echo "Creating Darwin full Package"
      tar -C $builddir/orchestrator -czf $TOPDIR/orchestrator-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
      echo "Creating Darwin cli Package"
      cp $builddir/orchestrator/usr/local/orchestrator/orchestrator $builddir/orchestrator-cli/usr/bin
      tar -C $builddir/orchestrator-cli -czf $TOPDIR/orchestrator-cli-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
      ;;
  esac

  echo "---"
  echo "Done. Find releases in $TOPDIR"
}

function build() {
  local target arch builddir gobuild

  os="$1"
  arch="$2"
  builddir="$3"
  gobuild="go build -o $builddir/orchestrator/usr/local/orchestrator/orchestrator go/cmd/orchestrator/main.go"

  case $os in
    'linux')
      GOOS=$os GOARCH=$arch $gobuild
    ;;
    'darwin')
      GOOS=darwin GOARCH=amd64 $gobuild
    ;;
  esac
}

function main() {
  local target arch builddir
  target="$1"
  arch="$2"

  precheck
  builddir=$( setuptree )
  oinstall $builddir
  build $target $arch $builddir
  package $target $builddir
  # cleanup
}

while getopts a:t:dh flag; do
  case $flag in
  a)
    arch=$OPTARG
    ;;
  t)
    target=$OPTARG
    ;;
  h)
    usage
    exit 0
    ;;
  d)
    debug=1
    ;;
  ?)
    usage
    exit 2
    ;;
  esac
done

shift $(( OPTIND - 1 ));
target=${target:-"linux"} # default for target is linux
arch=${arch:-"amd64"} # default for arch is amd64

[[ $debug -eq 1 ]] && set -x
main $target $arch

