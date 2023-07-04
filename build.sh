#!/bin/bash -exu

function main() {
  local version
  version="${1}"

  for os in darwin linux; do
    for arch in amd64 arm64; do
      GOOS="${os}" GOARCH="${arch}" go build -o "faux-${os}-${arch}" -ldflags "-X main.version=${version}" .
      shasum -a 256 "faux-${os}-${arch}"
    done
  done

}

main "${@:-}"
