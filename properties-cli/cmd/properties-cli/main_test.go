package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeTempProps(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "app.properties")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write temp properties file: %v", err)
	}
	return path
}

func TestPropReadExistingKey(t *testing.T) {
	path := writeTempProps(t, "foo=bar\nbaz=qux\n")

	got := propRead("baz", path)
	if got != "qux" {
		t.Fatalf("propRead() = %q, want %q", got, "qux")
	}
}

func TestPropReadFirstOccurrenceOnly(t *testing.T) {
	path := writeTempProps(t, "foo=first\nfoo=second\n")

	got := propRead("foo", path)
	if got != "first" {
		t.Fatalf("propRead() = %q, want %q", got, "first")
	}
}

func TestPropReadMissingKey(t *testing.T) {
	path := writeTempProps(t, "foo=bar\n")

	got := propRead("missing", path)
	if got != "" {
		t.Fatalf("propRead() = %q, want empty string", got)
	}
}

func TestPropReadSkipsCommentsAndBlankLines(t *testing.T) {
	path := writeTempProps(t, "# comment\n\n; another comment\nfoo=bar\n")

	got := propRead("foo", path)
	if got != "bar" {
		t.Fatalf("propRead() = %q, want %q", got, "bar")
	}
}

func TestPropWriteAppendsNewKey(t *testing.T) {
	path := writeTempProps(t, "foo=bar\n")

	propWrite("baz", "qux", path)

	if got := propRead("baz", path); got != "qux" {
		t.Fatalf("propRead() after propWrite() = %q, want %q", got, "qux")
	}
	if got := propRead("foo", path); got != "bar" {
		t.Fatalf("propRead() for untouched key = %q, want %q", got, "bar")
	}
}

func TestPropWriteReplacesExistingKey(t *testing.T) {
	path := writeTempProps(t, "foo=bar\nbaz=old\n")

	propWrite("baz", "new", path)

	if got := propRead("baz", path); got != "new" {
		t.Fatalf("propRead() after propWrite() = %q, want %q", got, "new")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read back file: %v", err)
	}
	if strings.Count(string(data), "baz=") != 1 {
		t.Fatalf("expected exactly one baz= line, got content: %q", string(data))
	}
}

func TestPropWriteEndsWithTrailingNewline(t *testing.T) {
	path := writeTempProps(t, "foo=bar\n")

	propWrite("baz", "qux", path)

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read back file: %v", err)
	}
	eol := detectPlatformEOL()
	if !strings.HasSuffix(string(data), eol) {
		t.Fatalf("expected file to end with %q, got: %q", eol, string(data))
	}
}

func TestDetectPlatformEOL(t *testing.T) {
	eol := detectPlatformEOL()
	if eol != "\n" && eol != "\r\n" {
		t.Fatalf("detectPlatformEOL() = %q, want \\n or \\r\\n", eol)
	}
}
