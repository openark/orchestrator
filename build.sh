#!/bin/bash

# Simple packaging of orchestrator
#
# Requires fpm: https://github.com/jordansissel/fpm
#
set -e

basedir=$(dirname $0)
GIT_COMMIT=$(git rev-parse HEAD)
RELEASE_VERSION=
RELEASE_SUBVERSION=
release_base_path=/tmp/orchestrator-release
export RELEASE_VERSION release_base_path

binary_build_path="build"
binary_artifact="$binary_build_path/bin/orchestrator"
skip_build=""
build_only=""
build_paths=""
retain_build_paths=""
opt_race=

usage() {
  echo
  echo "Usage: $0 [-t target ] [-a arch ] [-i init-system] [ -p prefix ] [-h] [-d] [-r]"
  echo "Options:"
  echo "-h Show this screen"
  echo "-t (linux|darwin) Target OS Default:(linux)"
  echo "-i (systemd|sysv) Target init system Default:(systemd)"
  echo "-a (amd64|386) Arch Default:(amd64)"
  echo "-d debug output"
  echo "-b build only, do not generate packages"
  echo "-N do not build; use existing ./build/bin/orchestrator binary"
  echo "-P create build/deployment paths"
  echo "-R retain existing build/deployment paths"
  echo "-p build prefix Default:(/usr/local)"
  echo "-r build with race detector"
  echo "-v release version (optional; default: content of RELEASE_VERSION file)"
  echo "-s release subversion (optional; default: empty)"
  echo
}


function fail() {
  local message="${1}"

  export message
  (>&2 echo "$message")
  exit 1
}

function debug() {
  local message="${1}"

  export message
  (>&2 echo "[DEBUG] $message")
}


function precheck() {
  local target="$1"
  local build_only="$2"
  local ok=0 # return err. so shell exit code

  if [[ "$target" == "linux" ]]; then
    if [[ -z "$build_only" ]] && [[ ! -x "$( which fpm )" ]]; then
      echo "Please install fpm and ensure it is in PATH (typically: 'gem install fpm')"
      ok=1
    fi

    if [[ -z "$build_only" ]] && [[ ! -x "$( which rpmbuild )" ]]; then
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

  if ! go version | egrep -q 'go(1\.1[6789])' ; then
    echo "go version must be 1.16 or above"
    ok=1
  fi

  return $ok
}

setup_build_path() {
  local prefix build_path
  prefix="$1"

  mkdir -p $release_base_path
  if [ -z "$retain_build_paths" ] ; then
    rm -rf ${release_base_path:?}/*
  fi
  build_path=$(mktemp -d $release_base_path/orchestratorXXXXXX) || fail "Unable to 'mktemp -d $release_base_path/orchestratorXXXXXX'"

  echo $build_path
}

build_binary() {
  local target arch build_path gobuild prefix
  os="$1"
  arch="$2"
  build_path="$3"
  prefix="$4"
  ldflags="-X main.AppVersion=${RELEASE_VERSION} -X main.GitCommit=${GIT_COMMIT}"
  debug "Building via $(go version)"
  mkdir -p "$binary_build_path/bin"
  rm -f $binary_artifact
  gobuild="go build -i ${opt_race} -ldflags \"$ldflags\" -o $binary_artifact go/cmd/orchestrator/main.go"

  case $os in
    'linux')
      echo "GOOS=$os GOARCH=$arch $gobuild" | bash
    ;;
    'darwin')
      echo "GOOS=darwin GOARCH=amd64 $gobuild" | bash
    ;;
  esac
  find $binary_artifact -type f || fail "Failed to generate orchestrator binary"
}

copy_binary_artifacts() {
  build_path="$1"
  prefix="$2"

  for dest in "orchestrator-cli/usr/bin" "orchestrator$prefix/orchestrator" ; do
    cp $binary_artifact $build_path/$dest && debug "binary copied to $build_path/$dest" || fail "Failed to copy orchestrator binary to $build_path/$dest"
  done
  cp $build_path/orchestrator${prefix}/orchestrator/resources/bin/orchestrator-client $build_path/orchestrator-client/usr/bin && debug "orchestrator-client copied to orchestrator-client/" || fail "Failed to copy orchestrator-client to orchestrator-client/"
}

setup_artifact_paths() {
  local prefix build_path
  build_path="$1"
  prefix="$2"

  mkdir -p $build_path/orchestrator
  mkdir -p $build_path/orchestrator${prefix}/orchestrator/
  mkdir -p $build_path/orchestrator-cli/usr/bin
  mkdir -p $build_path/orchestrator-client/usr/bin
  [ "$init_system" == "sysv" ] && mkdir -p $build_path/orchestrator/etc/init.d
  [ "$init_system" == "systemd" ] && mkdir -p $build_path/orchestrator/etc/systemd/system
  ln -s $build_path $release_base_path/build
}

function copy_resource_artifacts() {
  local build_path prefix
  build_path="$1"
  init_system="$2"
  prefix="$3"

  cd  $basedir
  gofmt -s -w  go/
  rsync -qa ./resources $build_path/orchestrator${prefix}/orchestrator/
  rsync -qa ./conf/orchestrator-sample*.conf.json $build_path/orchestrator${prefix}/orchestrator/

  case $init_system in
    "sysv")
      cp etc/init.d/orchestrator.bash $build_path/orchestrator/etc/init.d/orchestrator
      chmod +x $build_path/orchestrator/etc/init.d/orchestrator
      ;;
    "systemd")
      cp etc/systemd/orchestrator.service $build_path/orchestrator/etc/systemd/system/orchestrator.service
      ;;
    *)
      fail "Unsupported init system '$init_system'."
      ;;
  esac
}

package_linux() {
  local arch build_path
  arch="$1"
  build_path="$2"

  local do_tar=1
  local do_rpm=1
  local do_deb=1

  cd $basedir

  tmp_build_path="$release_base_path/build/tmp"
  mkdir -p $tmp_build_path
  rm -f "${tmp_build_path:-?}/*.*"

  package_name_extra=""

  if [ "$init_system" == "sysv" ] ; then
    package_name_extra="${package_name_extra}-sysv"
  fi

  if [[ -f /etc/centos-release ]]; then
      if cat /etc/centos-release | grep 'CentOS release 6' ; then
        do_tar=0
        do_deb=0
        # n CentOD 6 box: we only want the rpms for CentOS6
        # Add "-centos6" to the file name.
        package_name_extra="${package_name_extra}-centos6"
      fi
  fi

  cd $tmp_build_path

  debug "Creating Linux Tar package"
  [ $do_tar -eq 1 ] && tar -C $build_path/orchestrator -czf $release_base_path/orchestrator-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./

  debug "Creating Distro full packages"
  [ $do_rpm -eq 1 ] && fpm -v "${RELEASE_VERSION}" --epoch 1 -f -s dir -n orchestrator -m shlomi-noach --description "MySQL replication topology management and HA" --url "https://github.com/openark/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $build_path/orchestrator --prefix=/ --config-files /usr/local/orchestrator/resources/public/css/custom.css --config-files /usr/local/orchestrator/resources/public/js/custom.js --depends 'jq >= 1.5' -t rpm .
  [ $do_deb -eq 1 ] && fpm -v "${RELEASE_VERSION}" --epoch 1 -f -s dir -n orchestrator -m shlomi-noach --description "MySQL replication topology management and HA" --url "https://github.com/openark/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $build_path/orchestrator --prefix=/ --config-files /usr/local/orchestrator/resources/public/css/custom.css --config-files /usr/local/orchestrator/resources/public/js/custom.js --depends 'jq >= 1.5' -t deb --deb-no-default-config-files .

  debug "Creating Distro cli packages"
  # orchestrator-cli packaging -- executable only
  [ $do_rpm -eq 1 ] && fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -n orchestrator-cli -m shlomi-noach --description "MySQL replication topology management and HA: binary only" --url "https://github.com/openark/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $build_path/orchestrator-cli --prefix=/ --depends 'jq >= 1.5' -t rpm .
  [ $do_deb -eq 1 ] && fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -n orchestrator-cli -m shlomi-noach --description "MySQL replication topology management and HA: binary only" --url "https://github.com/openark/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $build_path/orchestrator-cli --prefix=/ --depends 'jq >= 1.5' -t deb --deb-no-default-config-files .

  debug "Creating Distro orchestrator-client packages"
  # orchestrator-client packaging -- shell script only
  [ $do_rpm -eq 1 ] && fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -n orchestrator-client -m shlomi-noach --description "MySQL replication topology management and HA: client script" --url "https://github.com/openark/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $build_path/orchestrator-client --prefix=/ --depends 'jq >= 1.5' -t rpm .
  [ $do_deb -eq 1 ] && fpm -v "${RELEASE_VERSION}" --epoch 1  -f -s dir -n orchestrator-client -m shlomi-noach --description "MySQL replication topology management and HA: client script" --url "https://github.com/openark/orchestrator" --vendor "GitHub" --license "Apache 2.0" -C $build_path/orchestrator-client --prefix=/ --depends 'jq >= 1.5' -t deb --deb-no-default-config-files .

  if [ ! -z "$package_name_extra" ] ; then
    ls *.rpm | while read f; do package_file=$(echo $f | sed -r -e "s/^(.*)-${RELEASE_VERSION}(.*)/\1${package_name_extra}-${RELEASE_VERSION}\2/g") ; mv $f $package_file ; done
    ls *.deb | while read f; do package_file=$(echo $f | sed -r -e "s/^(.*)_${RELEASE_VERSION}(.*)/\1${package_name_extra}-${RELEASE_VERSION}\2/g") ; mv $f $package_file ; done
  fi

  debug "packeges:"
  for f in * ; do debug "- $f" ; done

  mv ./*.* $release_base_path/
  cd $basedir
}

package_darwin() {
  local arch build_path
  arch="$1"
  build_path="$2"

  cd $release_base_path
  debug "Creating Darwin full Package"
  tar -C $build_path/orchestrator -czf $release_base_path/orchestrator-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
  debug "Creating Darwin cli Package"
  tar -C $build_path/orchestrator-cli -czf $release_base_path/orchestrator-cli-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
  debug "Creating Darwin orchestrator-client Package"
  tar -C $build_path/orchestrator-client -czf $release_base_path/orchestrator-client-"${RELEASE_VERSION}"-$target-$arch.tar.gz ./
}

package() {
  local target arch build_path
  target="$1"
  arch="$2"
  build_path="$3"

  debug "Release version is ${RELEASE_VERSION} (${GIT_COMMIT})"

  case $target in
    'linux') package_linux "$arch" "$build_path" ;;
    'darwin') package_darwin "$arch" "$build_path" ;;
  esac

  debug "Done. Find releases in $release_base_path"
}

main() {
  local target="$1"
  local init_system="$2"
  local arch="$3"
  local prefix="$4"
  local build_only=$5
  local build_path

  if [ -z "${RELEASE_VERSION}" ] ; then
    RELEASE_VERSION=$(cat $basedir/RELEASE_VERSION)
  fi
  RELEASE_VERSION="${RELEASE_VERSION}${RELEASE_SUBVERSION}"

  precheck "$target" "$build_only"
  if [ -z "$skip_build" ] ; then
    build_binary "$target" "$arch" "$build_path" "$prefix" || fail "Failed building binary"
  fi
  if [ "$build_paths" == "true" ] ; then
    build_path=$(setup_build_path "$prefix")
    setup_artifact_paths "$build_path" "$prefix"
    copy_resource_artifacts "$build_path" "$init_system" "$prefix"
    copy_binary_artifacts "$build_path" "$prefix"
  fi
  if [[ -z "$build_only" ]]; then
    package "$target" "$arch" "$build_path"
  fi
}

while getopts "a:t:i:p:s:v:dbNPRhr" flag; do
  case $flag in
  a)
    arch="${OPTARG}"
    ;;
  t)
    target="${OPTARG}"
    ;;
  i)
    init_system="${OPTARG}"
    ;;
  h)
    usage
    exit 0
    ;;
  d)
    debug=1
    ;;
  b)
    debug "Build only; no packaging"
    build_only="true"
    ;;
  N)
    debug "skipping build"
    [ -f "$binary_artifact" ] || fail "cannot find $binary_artifact"
    skip_build="true"
    ;;
  P)
    debug "Creating build paths"
    build_paths="true"
    ;;
  R)
    debug "Retaining existing build paths"
    retain_build_paths="true"
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

if [ -z "$build_only" ] ; then
  # To build packages means we also need to build the paths
  build_paths="true"
fi

if [ -z "$target" ]; then
	uname=$(uname)
	case $uname in
    Linux)	target=linux ;;
    Darwin)	target=darwin ;;
    *)      fail "Unexpected OS from uname: $uname. Exiting" ;;
	esac
fi

init_system=${init_system:-"systemd"}
arch=${arch:-"amd64"} # default for arch is amd64 but should take from environment
prefix=${prefix:-"/usr/local"}

[[ $debug -eq 1 ]] && set -x
main "$target" "$init_system" "$arch" "$prefix" "$build_only"

debug "orchestrator build done; exit status is $?"
