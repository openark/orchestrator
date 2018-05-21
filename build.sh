#!/bin/bash

# Simple packaging of orchestrator
#
# Requires fpm: https://github.com/jordansissel/fpm
#
set -e

mydir=$(dirname $0)
GIT_COMMIT=$(git rev-parse HEAD)
RELEASE_VERSION=
RELEASE_SUBVERSION=
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
  echo "-v release version (optional; default: content of RELEASE_VERSION file)"
  echo "-s release subversion (optional; default: empty)"
  echo
}

function precheck() {
  local target="$1"
  local build_only="$2"
  local ok=0 # return err. so shell exit code

  if [[ "$target" == "linux" ]]; then
    if [[ $build_only -eq 0 ]] && [[ ! -x "$( which fpm )" ]]; then
      echo "Please install fpm and ensure it is in PATH (typically: 'gem install fpm')"
      ok=1
    fi

    if [[ $build_only -eq 0 ]] && [[ ! -x "$( which rpmbuild )" ]]; then
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

  if ! go version | egrep -q 'go(1[.]8|1[.]9|1[.]10)' ; then
    echo "go version is too low. Must use 1.8 or above"
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
  mkdir -p $b/orchestrator-client/usr/bin
  echo $b
}

function oinstall() {
  local builddir prefix
  builddir="$1"
  prefix="$2"

  cd  $mydir
  gofmt -s -w  go/
  rsync -qa ./resources $builddir/orchestrator${prefix}/orchestrator/
  rsync -qa ./conf/orchestrator-sample*.conf.json $builddir/orchestrator${prefix}/orchestrator/
  cp etc/init.d/orchestrator.bash $builddir/orchestrator/etc/init.d/orchestrator
  chmod +x $builddir/orchestrator/etc/init.d/orchestrator
}

function package() {
  local target builddir prefix packages
  target="$1"
  builddir="$2"
  prefix="$3"

  cd $TOPDIR

  echo "Release version is ${RELEASE_VERSION} ( ${GIT_COMMIT} )"

  case $target in
    'linux')
      echo "Creating Linux Tar package"
      tar -C $builddir/orchestrator -czf $TOPDIR/orchestrator-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./

      echo "Creating Distro full packages"
      fpm -v "${RELEASE_VERSION}" --epoch 1 -f -s dir -n orchestrator -m shlomi-noach --description "MySQL replication topology management and HA" --url "https://github.com/github/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $builddir/orchestrator --prefix=/ --config-files /usr/local/orchestrator/resources/public/css/custom.css --config-files /usr/local/orchestrator/resources/public/js/custom.js -t rpm .
      fpm -v "${RELEASE_VERSION}" --epoch 1 -f -s dir -n orchestrator -m shlomi-noach --description "MySQL replication topology management and HA" --url "https://github.com/github/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $builddir/orchestrator --prefix=/ --config-files /usr/local/orchestrator/resources/public/css/custom.css --config-files /usr/local/orchestrator/resources/public/js/custom.js -t deb .

      cd $TOPDIR
      # orchestrator-cli packaging -- executable only
      echo "Creating Distro cli packages"
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -n orchestrator-cli -m shlomi-noach --description "MySQL replication topology management and HA: binary only" --url "https://github.com/github/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $builddir/orchestrator-cli --prefix=/ -t rpm .
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -n orchestrator-cli -m shlomi-noach --description "MySQL replication topology management and HA: binary only" --url "https://github.com/github/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $builddir/orchestrator-cli --prefix=/ -t deb --deb-no-default-config-files .

      cd $TOPDIR
      # orchestrator-client packaging -- shell script only
      echo "Creating Distro orchestrator-client packages"
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -n orchestrator-client -m shlomi-noach --description "MySQL replication topology management and HA: client script" --url "https://github.com/github/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $builddir/orchestrator-client --prefix=/ -t rpm .
      fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -n orchestrator-client -m shlomi-noach --description "MySQL replication topology management and HA: client script" --url "https://github.com/github/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $builddir/orchestrator-client --prefix=/ -t deb --deb-no-default-config-files .
      ;;
    'darwin')
      echo "Creating Darwin full Package"
      tar -C $builddir/orchestrator -czf $TOPDIR/orchestrator-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
      echo "Creating Darwin cli Package"
      tar -C $builddir/orchestrator-cli -czf $TOPDIR/orchestrator-cli-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
      echo "Creating Darwin orchestrator-client Package"
      tar -C $builddir/orchestrator-client -czf $TOPDIR/orchestrator-client-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
      ;;
  esac

  echo "---"
  if [[ -f /etc/centos-release ]]; then
      if cat /etc/centos-release | grep 'CentOS release 6' ; then
        rm ${TOPDIR:-?}/orchestrator*.deb
        rm ${TOPDIR:-?}/orchestrator*.tar.gz
        # n CentOD 6 box: we only want the rpms for CentOS6
        # Add "-centos6" to the file name.
        ls ${TOPDIR:-?}/*.rpm | while read f; do centos_file=$(echo $f | sed -r -e "s/^(.*)-${RELEASE_VERSION}(.*)/\1-centos6-${RELEASE_VERSION}\2/g") ; mv $f $centos_file ; done
      fi
  fi
  echo "Done. Find releases in $TOPDIR"
}

function build() {
  local target arch builddir gobuild prefix
  os="$1"
  arch="$2"
  builddir="$3"
  prefix="$4"
  ldflags="-X main.AppVersion=${RELEASE_VERSION} -X main.GitCommit=${GIT_COMMIT}"
  echo "Building via $(go version)"
  gobuild="go build -i ${opt_race} -ldflags \"$ldflags\" -o $builddir/orchestrator${prefix}/orchestrator/orchestrator go/cmd/orchestrator/main.go"

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
  cp $builddir/orchestrator${prefix}/orchestrator/resources/bin/orchestrator-client $builddir/orchestrator-client/usr/bin && echo "orchestrator-client copied to orchestrator-client/" || (echo "Failed to copy orchestrator-client to orchestrator-client/" ; exit 1)
}

function main() {
  local target="$1"
  local arch="$2"
  local prefix="$3"
  local build_only=$4
  local builddir

  if [ -z "${RELEASE_VERSION}" ] ; then
    RELEASE_VERSION=$(cat $mydir/RELEASE_VERSION)
  fi
  RELEASE_VERSION="${RELEASE_VERSION}${RELEASE_SUBVERSION}"

  precheck "$target" "$build_only"
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
while getopts a:t:p:s:v:dbhr flag; do
  case $flag in
  a)
    arch="${OPTARG}"
    ;;
  t)
    target="${OPTARG}"
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
    prefix="${OPTARG}"
    ;;
  r)
    opt_race="-race"
    ;;
  v)
    RELEASE_VERSION="${OPTARG}"
    ;;
  s)
    RELEASE_SUBVERSION="_${OPTARG}"
    ;;
  ?)
    usage
    exit 2
    ;;
  esac
done

shift $(( OPTIND - 1 ));

if [ -z "$target" ]; then
	uname=$(uname)
	case $uname in
	Linux)	target=linux;;
	Darwin)	target=darwin;;
	*)	echo "Unexpected OS from uname: $uname. Exiting"
		exit 1
	esac
fi
arch=${arch:-"amd64"} # default for arch is amd64 but should take from environment
prefix=${prefix:-"/usr/local"}

[[ $debug -eq 1 ]] && set -x
main "$target" "$arch" "$prefix" "$build_only"

echo "orchestrator build done; exit status is $?"
