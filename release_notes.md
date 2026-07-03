## Changelog for v0.2.0

This release is a repository relaunch — the CLI's flags, output, and read/write behavior are unchanged. It focuses on making the project properly usable, testable, and maintainable by others:

- **Licensing**: added an MIT `LICENSE` (previously the repo had none).
- **Testability**: added unit tests for `propRead`, `propWrite`, and `detectPlatformEOL` (previously zero Go test coverage — the only testing was a shell-based stress test, which is kept as-is for local/manual use).
- **CI/CD migrated from Azure DevOps to GitHub Actions**: runs on `ubuntu-latest`. Releases are now triggered by pushing a `vX.Y.Z` tag instead of every push to `main`; the published archives are still the same 5 targets (`windows-386`/`windows-amd64`/`windows-arm64` as `.zip`, `darwin-amd64`/`darwin-arm64` as `.tar.gz`) with the same file naming as before.
- **Project layout**: moved `properties-cli` to the standard `cmd/properties-cli/` + `internal/version/` Go layout (the `go.work` multi-module structure itself is unchanged).
- **Docs**: expanded `README.md` with install/usage/dev instructions, added `AGENTS.md` for AI coding assistants.
- Translated the remaining Chinese comments in `main.go`, `scripts/base/env.sh`, and `scripts/update-sh-format.sh` to English.
- Removed the now-unused Azure-specific scripts (`scripts/config/gh-config.sh`, `scripts/upload/`); everything else under `scripts/` (build, compress, the shell-script formatter, the stress test) is unchanged.
