# AGENTS.md

Guidance for AI coding assistants (Claude Code, Codex, etc.) working in this repository.

## Repo structure: go.work workspace

This repo uses a `go.work` file at the root that references a single module, `./properties-cli`. All Go commands need to be run from inside `properties-cli/` (or with `go.work` resolving it — most `go` subcommands do this automatically from the repo root too, but when in doubt `cd properties-cli` first).

- `properties-cli/cmd/properties-cli/main.go` — CLI entrypoint and core read/write logic
- `properties-cli/cmd/properties-cli/main_test.go` — unit tests (use `t.TempDir()`, no external dependencies)
- `properties-cli/internal/version/version.go` — single `Version` constant, bumped manually before each release

## Release artifacts are compressed archives, not raw binaries

Unlike some of the other CLI tools in this account, this one has always published **compressed archives**: `.zip` for the three Windows targets, `.tar.gz` for the two macOS targets. Keep it that way — don't switch to raw binaries "for consistency" with other repos; that would break the existing download convention for this specific tool.

## `scripts/test/properties-test.sh` depends on another repo's release

This shell script cross-checks a pure-shell (`awk`/`grep`) properties implementation against `properties-cli` itself, running 500 iterations of read/write and comparing results. On Windows, it needs to detect the CPU architecture, and for that it downloads `windows-os-info.exe` from `jason-xie-123/windows-os-info`'s GitHub Releases (see `check_windows_os_info_exist` in `scripts/base/env.sh`). This is a real, intentional cross-repo dependency — don't "fix" it by removing the download or assuming it's a mistake.

This script is **not** part of CI (`.github/workflows/ci.yml`) — it needs a fully built and compressed release first, which doesn't fit a fast per-PR check. Run it manually via `scripts/local-build.sh` when you want to stress-test a change.

## Build, test, lint

```sh
cd properties-cli
go build ./...
go test ./...
gofmt -l .              # must produce no output
golangci-lint run ./...  # must report 0 issues
```

## Commit messages

Write commit messages in English. Keep them short and describe the actual change — avoid placeholder messages like `init` or `update`.

## Release process

Releases are tag-triggered, not push-triggered:

1. Draft `release_notes.md` locally by reading the diff since the last tag (`git diff <last-tag>..HEAD`) — an AI assistant can draft this, but a human must review it before tagging.
2. Bump the `Version` constant in `properties-cli/internal/version/version.go` to match the new tag.
3. `git tag vX.Y.Z && git push origin vX.Y.Z` — this triggers `.github/workflows/release.yml`, which cross-compiles all 5 targets, packages them into the same archive formats as before, and creates the GitHub Release using the committed `release_notes.md`.

Do not call any LLM API from within CI to generate release notes — that step happens locally, before tagging, to avoid paying per-run API costs in the pipeline.
