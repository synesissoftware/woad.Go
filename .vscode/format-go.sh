#! /usr/bin/env bash

# Formats Go source on stdin for the VS Code Go extension.
# gofmt, then gci in the Synesis group order. The extension's
# organize-imports action uses goimports order; this formatter
# runs as the document formatter so a save ends on the gci order.
# The working directory is the source file's directory.

set -eu

source=$(mktemp)
mv "$source" "$source.go"
source="$source.go"
trap 'rm -f "$source"' EXIT
cat > "$source"
gofmt -w "$source"

gci_bin=$(command -v gci || true)
if [ -z "$gci_bin" ]; then
	gopath=$(go env GOPATH 2> /dev/null || true)
	if [ -n "$gopath" ] && [ -x "$gopath/bin/gci" ]; then
		gci_bin="$gopath/bin/gci"
	fi
fi

if [ -z "$gci_bin" ]; then
	cat "$source"
	exit 0
fi

here=$(pwd)
case "$here" in
	*/examples | */examples/*)
		"$gci_bin" print --custom-order -s default -s standard "$source" < /dev/null
		;;
	*)
		"$gci_bin" print --custom-order \
			-s 'prefix(github.com/synesissoftware/woad.Go)' \
			-s default \
			-s standard \
			"$source" < /dev/null
		;;
esac
