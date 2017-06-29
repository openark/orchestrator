#!/bin/bash

PACKAGECLOUD_USER="${PACKAGECLOUD_USER:-github}"
REPOSITORY=orchestrator
packagecloud_uri="https://${PACKAGECLOUD_TOKEN}:@packagecloud.io"
latest_tag=$(git describe --abbrev=0 --tags | tr -d "v")

function fail() {
  message="$1"
  >&2 echo "$message"
}

function verify() {
  if [ -z "$latest_tag" ] ; then
    fail "Cannot find latest tag"
  fi

  if [ -z "$PACKAGECLOUD_TOKEN" ] ; then
    fail "Cannot find PACKAGECLOUD_TOKEN"
  fi
}

function list_old_packages() {
  url="${packagecloud_uri}/api/v1/repos/${PACKAGECLOUD_USER}/${REPOSITORY}/packages.json?per_page=1000"
  curl -s "$url" | jq -r ".[] | select(.version != \"$latest_tag\") | .destroy_url"
}

function purge_package() {
  destroy_url="$packagecloud_uri/$1"
  curl -s -X DELETE "$destroy_url"
  echo "destroyed $destroy_url"
}

main() {
  verify
  list_old_packages | while read destroy_url ; do
    purge_package "$destroy_url"
  done
}

main
