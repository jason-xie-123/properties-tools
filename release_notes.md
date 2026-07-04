## Changelog for v0.2.1

Bug-fix release, no CLI behavior changes.

- **Fixed `go install` not working**: `properties-cli/go.mod` declared its module path as the bare name `properties-cli`, which conflicted with the `go install github.com/jason-xie-123/properties-tools/properties-cli/cmd/properties-cli@latest` command documented in the README. The module path is now `github.com/jason-xie-123/properties-tools/properties-cli`, matching the actual import path.
