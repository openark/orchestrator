#!/bin/bash

# Simple packaging of orchestrator
#
# Requires fpm: https://github.com/jordansissel/fpm
#
set -e

RELEASE_VERSION="1.4.337"
TOPDIR=/tmp/orchestrator-release
export RELEASE_VERSION TOPDIR

usage() {
  echo
  echo "Usage: $0 [-t target ] [-a arch ] [ -p prefix ] [-h] [-d]"
  echo "Options:"
  echo "-h Show this screen"
  echo "-t (linux|darwin) Target OS Default:(linux)"
  echo "-a (amd64|386) Arch Default:(amd64)"
  echo "-d debug output"
  echo "-p build prefix Default:(/usr/local)"
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
  local b prefix
  prefix="$1"

  mkdir -p $TOPDIR
  rm -rf $TOPDIR/*
  b=$( mktemp -d $TOPDIR/orchestratorXXXXXX ) || return 1
  mkdir -p $b/orchestrator
  mkdir -p $b/orchestrator${prefix}/orchestrator/
  mkdir -p $b/orchestrator/etc/init.d
  mkdir -p $b/orchestrator-cli/usr/bin
  echo $b
}

function oinstall() {
  local builddir prefix
  builddir="$1"
  prefix="$2"

  cd  $(dirname $0)
  gofmt -s -w  go/
  rsync -qa ./resources $builddir/orchestrator${prefix}/orchestrator/
  rsync -qa ./conf/*.sample $builddir/orchestrator${prefix}/orchestrator/
  cp etc/init.d/orchestrator.bash $builddir/orchestrator/etc/init.d/orchestrator
  chmod +x $builddir/orchestrator/etc/init.d/orchestrator
}

function package() {
  local target builddir prefix packages
  target="$1"
  builddir="$2"
  prefix="$3"

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
      cp $builddir/orchestrator${prefix}/orchestrator/orchestrator $builddir/orchestrator-cli/usr/bin
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -t rpm -n orchestrator-cli -C $builddir/orchestrator-cli --prefix=/ .
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -t deb -n orchestrator-cli -C $builddir/orchestrator-cli --prefix=/ .
      ;;
    'darwin')
      echo "Creating Darwin full Package"
      tar -C $builddir/orchestrator -czf $TOPDIR/orchestrator-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
      echo "Creating Darwin cli Package"
      cp $builddir/orchestrator${prefix}/orchestrator/orchestrator $builddir/orchestrator-cli/usr/bin
      tar -C $builddir/orchestrator-cli -czf $TOPDIR/orchestrator-cli-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
      ;;
  esac

  echo "---"
  echo "Done. Find releases in $TOPDIR"
}

function build() {
  local target arch builddir gobuild prefix
  os="$1"
  arch="$2"
  builddir="$3"
  prefix="$4"
  gobuild="go build -o $builddir/orchestrator${prefix}/orchestrator/orchestrator go/cmd/orchestrator/main.go"

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
  local target arch builddir prefix
  target="$1"
  arch="$2"
  prefix="$3"

  precheck
  builddir=$( setuptree "$prefix" )
  oinstall "$builddir" "$prefix"
  build "$target" "$arch" "$builddir" "$prefix"
  package "$target" "$builddir" "$prefix"
  # cleanup
}

while getopts a:t:p:dh flag; do
  case $flag in
  a)
    arch="$OPTARG"
    ;;
  t)
    target="$OPTARG"
    ;;
  h)
    usage
    exit 0
    ;;
  d)
    debug=1
    ;;
  p)
    prefix="$OPTARG"
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
prefix=${prefix:-"/usr/local"}

[[ $debug -eq 1 ]] && set -x
main "$target" "$arch" "$prefix"

