# properties-tools

[![CI](https://github.com/jason-xie-123/properties-tools/actions/workflows/ci.yml/badge.svg)](https://github.com/jason-xie-123/properties-tools/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/jason-xie-123/properties-tools)](https://github.com/jason-xie-123/properties-tools/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

A small CLI tool (`properties-cli`) to read and write `.properties` config files, with cross-platform line-ending handling.

## Install

Download the archive for your platform from the [latest release](https://github.com/jason-xie-123/properties-tools/releases/latest) — Windows (386/amd64/arm64, `.zip`) and macOS (amd64/arm64, `.tar.gz`) are all published on every release.

Or build from source:

```sh
go install github.com/jason-xie-123/properties-tools/properties-cli/cmd/properties-cli@latest
```

## How to Use

```
properties-cli -h
NAME:
   properties-cli - CLI Tool to read and write properties files

USAGE:
   properties-cli [global options] command [command options]

GLOBAL OPTIONS:
   --read         read flag (default: false)
   --write        write flag (default: false)
   --key value    property key name
   --value value  property key value
   --path value   path to properties file
   --help, -h     show help
   --version, -v  print the version
```

Example:

```sh
properties-cli --read --path=app.properties --key=my.key
properties-cli --write --path=app.properties --key=my.key --value=my-value
```

## Development

This repo is organized as a `go.work` workspace with a single module, `properties-cli`.

```sh
cd properties-cli
go build ./...
go test ./...
gofmt -l .
golangci-lint run ./...
```

There's also a shell-based stress/comparison test at `scripts/test/properties-test.sh` (not part of CI — see `AGENTS.md` for why) and a shell-script formatter/linter at `scripts/update-sh-format.sh`.

Releases are cut by pushing a `vX.Y.Z` tag — see `.github/workflows/release.yml`. Release notes live in `release_notes.md` and are drafted locally before tagging (see `AGENTS.md`).

## License

[MIT](./LICENSE)
