#!/bin/bash -exu

function main() {
  local version
  version="${1}"

  for os in darwin linux; do
    GOOS="${os}" GOARCH=amd64 go build -o "faux-${os}-amd64" -ldflags "-X main.version=${version}" .
    shasum -a 256 "faux-${os}-amd64"
  done

}

main "${@:-}"
