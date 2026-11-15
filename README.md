# woad.Go <!-- omit in toc -->

Minimal ANSI terminal colour codes, for Go

![Language](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/woad.Go.svg)](https://github.com/synesissoftware/woad.Go/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/woad.Go)](https://github.com/synesissoftware/woad.Go/commits/master)
[![Go](https://github.com/synesissoftware/woad.Go/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/woad.Go/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/woad.Go.svg)](https://pkg.go.dev/github.com/synesissoftware/woad.Go)


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Developer Experience](#developer-experience)
- [Installation \& Usage](#installation--usage)
- [Components](#components)
- [Environment](#environment)
- [Examples](#examples)
- [Project Information](#project-information)
  - [Where to get help](#where-to-get-help)
  - [Contribution guidelines](#contribution-guidelines)
  - [Dependencies](#dependencies)
    - [Development/Example/Testing Dependencies](#developmentexampletesting-dependencies)
  - [Related projects](#related-projects)
  - [License](#license)


## Introduction

**woad** provides the smallest useful set of fixed ANSI SGR colour sequences for library authors. It is not a console or TUI framework.

**woad.Go** is the **Go** implementation.


## Developer Experience

**Go 1.21** is the module floor. From a checkout:

```bash
go test ./...
go run ./examples/colour
```

The smallest useful call prints a sequence only when the writer supports colour:

```Go
import (
	woad "github.com/synesissoftware/woad.Go"

	"fmt"
	"os"
)

writer := os.Stdout
green := woad.FG_GREEN
reset := woad.RESET

fmt.Println(woad.SeqFor(writer, &green) + "ok" + woad.SeqFor(writer, &reset))
```

That is [examples/colour](./examples/colour/main.go). Fixed sequences (`BOLD`, `BLINKING`, `FG_*`, `BG_*`) always contain the SGR bytes. `ColourPolicy`, `SupportsColour`, and `SeqFor` apply the cached environment policy. `Refresh` drops that cache. A nil sequence yields `""`.


## Installation & Usage

Install via `go get`, as in:

```bash
go get "github.com/synesissoftware/woad.Go"
```

and then import as:

```Go
import woad "github.com/synesissoftware/woad.Go"
```

or, simply, as:

```Go
import "github.com/synesissoftware/woad.Go"
```


## Components

**woad.Go** ships fixed SGR sequences (`BLINKING`, `BOLD`, `RESET`, `FG_*`, `BG_*`, including bright variants), a version API (**`Version()`**, **`VersionString()`**), and the portable colour policy: **`ColourPolicy`**, **`Refresh`**, **`SupportsColour`**, and **`SeqFor`**. On Unix a terminal supports colour when the policy is automatic and `TERM` is not `dumb`. Windows virtual-terminal mode and mintty PTY recognition are not implemented.

```Go
fmt.Println(woad.FG_GREEN + "ok" + woad.RESET)
```

```Go
const (
	VersionMajor uint16 = /* ... */
	VersionMinor uint16 = /* ... */
	VersionPatch uint16 = /* ... */
	VersionAB    uint16 = /* ... */
)

func Version() uint64
func VersionString() string
```


## Environment

`ColourPolicy` returns `PolicyAuto` (0), `PolicyForce` (1), or `PolicyInhibit` (2). The result is cached until `Refresh`.

A variable that is unset or empty is absent. Values are not trimmed. The first present variable in this order wins:

1. Non-empty `WOAD_COLOUR`, otherwise non-empty `WOAD_COLOR`. Compared case-insensitively. Truey forces: `1`, `always`, `force`, `on`, `true`, `yes`. Falsey inhibits: `0`, `false`, `never`, `no`, `none`, `off`. Anything else, including `auto`, selects auto and does not fall through;
2. Non-empty `NO_COLOR` inhibits, whatever its text is;
3. Non-empty `CLICOLOR_FORCE` other than the exact bytes `0` forces. Exact `0` does not force and does not inhibit;
4. Non-empty `FORCE_COLOR` other than exact `0` or `false` (any capitalisation) forces. Exact `0` and `false` do not force and do not inhibit;
5. `CLICOLOR` equal to the exact bytes `0` inhibits. Any other value is ignored;
6. Otherwise auto.

`TERM` equal to `dumb`, in any capitalisation, is cached beside the policy. It is not a policy value. It suppresses colour only when the policy is auto. `PolicyForce` ignores the terminal and `TERM`. `PolicyInhibit` never emits colour.

Attachment is true only for an `*os.File` whose descriptor is a terminal (`golang.org/x/term`). Every other `io.Writer` is not a terminal. On Unix, auto plus a terminal plus a `TERM` that is not `dumb` means colour. There is no console probe. `WOAD_SET_CONSOLE_MODE` is not read in this alpha.


## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/woad.Go "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/woad.Go.


### Dependencies

* [**golang.org/x/term**](https://pkg.go.dev/golang.org/x/term) — descriptor terminal check for `SupportsColour` and `SeqFor`;
* [**ver2go**](https://github.com/synesissoftware/ver2go/);


#### Development/Example/Testing Dependencies

* [**testify**](https://github.com/stretchr/testify);


### Related projects

* [**woad**](https://github.com/synesissoftware/woad/);
* [**woad.Python**](https://github.com/synesissoftware/woad.Python/);
* [**woad.Ruby**](https://github.com/synesissoftware/woad.Ruby/);
* [**woad.Rust**](https://github.com/synesissoftware/woad.Rust/);


### License

**woad.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.


<!-- ########################### end of file ########################### -->
