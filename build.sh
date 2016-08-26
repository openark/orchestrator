#!/bin/bash

# Simple packaging of orchestrator
#
# Requires fpm: https://github.com/jordansissel/fpm
#
set -e

RELEASE_VERSION=$(cat RELEASE_VERSION)
TOPDIR=/tmp/orchestrator-release
export RELEASE_VERSION TOPDIR
export GO15VENDOREXPERIMENT=1

usage() {
  echo
  echo "Usage: $0 [-t target ] [-a arch ] [ -p prefix ] [-h] [-d] [-r]"
  echo "Options:"
  echo "-h Show this screen"
  echo "-t (linux|darwin) Target OS Default:(linux)"
  echo "-a (amd64|386) Arch Default:(amd64)"
  echo "-d debug output"
  echo "-b build only, do not generate packages"
  echo "-p build prefix Default:(/usr/local)"
  echo "-r build with race detector"
  echo "-s release subversion"
  echo
}

function precheck() {
  local target
  local ok=0 # return err. so shell exit code

  if [[ "$target" == "linux" ]]; then
    if [[ ! -x "$( which fpm )" ]]; then
      echo "Please install fpm and ensure it is in PATH (typically: 'gem install fpm')"
      ok=1
    fi

    if [[ ! -x "$( which rpmbuild )" ]]; then
      echo "rpmbuild not in PATH, rpm will not be built (OS/X: 'brew install rpm')"
    fi
  fi

  if [[ -z "$GOPATH" ]]; then
    echo "GOPATH not set"
    ok=1
  fi

  if [[ ! -x "$( which go )" ]]; then
    echo "go binary not found in PATH"
    ok=1
  fi

  if [[ $(go version | egrep "go1[.][01234]") ]]; then
    echo "go version is too low. Must use 1.5 or above"
    ok=1
  fi

  return $ok
}

function setuptree() {
  local b prefix
  prefix="$1"

  mkdir -p $TOPDIR
  rm -rf ${TOPDIR:?}/*
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
  rsync -qa ./conf/orchestrator-sample.* $builddir/orchestrator${prefix}/orchestrator/
  cp etc/init.d/orchestrator.bash $builddir/orchestrator/etc/init.d/orchestrator
  chmod +x $builddir/orchestrator/etc/init.d/orchestrator
}

function package() {
  local target builddir prefix packages
  target="$1"
  builddir="$2"
  prefix="$3"

  cd $TOPDIR

  echo "Release version is ${RELEASE_VERSION}"

  case $target in
    'linux')
      echo "Creating Linux Tar package"
      tar -C $builddir/orchestrator -czf $TOPDIR/orchestrator-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./

      echo "Creating Distro full packages"
      fpm -v "${RELEASE_VERSION}" --epoch 1 -f -s dir -t rpm -n orchestrator -C $builddir/orchestrator --prefix=/ .
      fpm -v "${RELEASE_VERSION}" --epoch 1 -f -s dir -t deb -n orchestrator -C $builddir/orchestrator --prefix=/ --deb-no-default-config-files .

      cd $TOPDIR
      # rpm packaging -- executable only
      echo "Creating Distro cli packages"
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -t rpm -n orchestrator-cli -C $builddir/orchestrator-cli --prefix=/ .
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -t deb -n orchestrator-cli -C $builddir/orchestrator-cli --prefix=/ --deb-no-default-config-files .
      ;;
    'darwin')
      echo "Creating Darwin full Package"
      tar -C $builddir/orchestrator -czf $TOPDIR/orchestrator-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
      echo "Creating Darwin cli Package"
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
  ldflags="-X main.AppVersion=${RELEASE_VERSION}"
  echo "Building via $(go version)"
  gobuild="go build ${opt_race} -ldflags \"$ldflags\" -o $builddir/orchestrator${prefix}/orchestrator/orchestrator go/cmd/orchestrator/main.go"

  case $os in
    'linux')
      echo "GOOS=$os GOARCH=$arch $gobuild" | bash
    ;;
    'darwin')
      echo "GOOS=darwin GOARCH=amd64 $gobuild" | bash
    ;;
  esac
  [[ $(find $builddir/orchestrator${prefix}/orchestrator/ -type f -name orchestrator) ]] &&  echo "orchestrator binary created" || (echo "Failed to generate orchestrator binary" ; exit 1)
  cp $builddir/orchestrator${prefix}/orchestrator/orchestrator $builddir/orchestrator-cli/usr/bin && echo "binary copied to orchestrator-cli" || (echo "Failed to copy orchestrator binary to orchestrator-cli" ; exit 1)
}

function main() {
  local target arch builddir prefix build_only
  target="$1"
  arch="$2"
  prefix="$3"
  build_only=$4

  precheck "$target"
  builddir=$( setuptree "$prefix" )
  oinstall "$builddir" "$prefix"
  build "$target" "$arch" "$builddir" "$prefix"
  [[ $? == 0 ]] || return 1
  if [[ $build_only -eq 0 ]]; then
    package "$target" "$builddir" "$prefix"
  fi
}

build_only=0
opt_race=
while getopts a:t:p:s:dbhr flag; do
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
  b)
    echo "Build only; no packaging"
    build_only=1
    ;;
  p)
    prefix="$OPTARG"
    ;;
  r)
    opt_race="-race"
    ;;
  s)
    RELEASE_VERSION="${RELEASE_VERSION}_${OPTARG}"
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
main "$target" "$arch" "$prefix" "$build_only"

echo "orchestrator build done; exit status is $?"
