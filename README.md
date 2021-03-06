# Drone Archive Plugin

[![Build Status](https://travis-ci.com/joseluisq/drone-archive.svg?branch=master)](https://travis-ci.com/joseluisq/drone-archive) [![codecov](https://codecov.io/gh/joseluisq/drone-archive/branch/master/graph/badge.svg)](https://codecov.io/gh/joseluisq/drone-archive) [![Go Report Card](https://goreportcard.com/badge/github.com/joseluisq/drone-archive)](https://goreportcard.com/report/github.com/joseluisq/drone-archive) [![PkgGoDev](https://pkg.go.dev/badge/github.com/joseluisq/drone-archive)](https://pkg.go.dev/github.com/joseluisq/drone-archive) [![Docker Image Version (tag latest semver)](https://img.shields.io/docker/v/joseluisq/drone-archive/1)](https://hub.docker.com/r/joseluisq/drone-archive/) [![Docker Image Size (tag)](https://img.shields.io/docker/image-size/joseluisq/drone-archive/1)](https://hub.docker.com/r/joseluisq/drone-archive/) [![Docker Image](https://img.shields.io/docker/pulls/joseluisq/drone-archive.svg)](https://hub.docker.com/r/joseluisq/drone-archive/)

> [Drone](https://drone.io/) plugin that provides Tar/Gzip and Zip archiving with optional checksum computation.

## Usage

```yml
---
kind: pipeline
type: docker
name: production

platform:
  os: linux
  arch: amd64

steps:
- name: archive
  image: joseluisq/drone-archive
  settings:
    format: tar
    src_base_path: ./my-base-path
    src: ./release/myprogram
    dest: ./myprogram.tar.gz
    checksum: true
    checksum_algo: sha256
    checksum_dest: myprogram.CHECKSUM.tar.gz.txt
```

## API

```sh
$ drone-archive --help
# NAME: archive plugin [OPTIONS] COMMAND
#
# Archive a file or directory using Tar/Gzip or Zip with optional checksum computation.
#
# OPTIONS:
#   -b --src-base-path    Source base path directory which will be skipped for each archive file header. [env: PLUGIN_SRC_BASE_PATH]
#   -s --src              File or directory to archive and compress. [env: PLUGIN_SRC]
#   -d --dest             File destination path to save the archived/compressed file. [env: PLUGIN_DEST]
#   -f --format           Define a `tar` and `zip` archiving format with compression. Tar format uses Gzip compression. [default: tar] [env: PLUGIN_FORMAT]
#   -c --checksum         Enable checksum file computation. [default: false] [env: PLUGIN_CHECKSUM]
#   -a --checksum-algo    Define the checksum `md5`, `sha1`, `sha256` or `sha512` algorithm. [default: sha256] [env: PLUGIN_CHECKSUM_ALGO]
#   -e --checksum-dest    File destination path of the checksum. [env: PLUGIN_CHECKSUM_DEST]
#   -h --help             Prints help information
#   -v --version          Prints version information
```

## Development

### Test

```sh
make test
```

### Build

Build the binaries and Docker image.

```sh
make build image.build
```

### Run

Run Docker images examples.

```sh
make image.tar
# or
make image.zip
```

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/joseluisq/drone-archive/pulls) or [issue](https://github.com/joseluisq/drone-archive/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2020-present [Jose Quintana](https://git.io/joseluisq)
