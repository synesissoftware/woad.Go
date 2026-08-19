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
- [Installation \& Usage](#installation--usage)
- [Components](#components)
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

**woad.Go** ships SGR string constants (`RESET`, `FG_*`, `BG_*`, including bright variants) and a version API. The sequences are always emitted; they do not inspect TTY state or Windows console mode. TTY/stream gating and Windows virtual-terminal opt-in are not implemented yet.

```Go
fmt.Println(woad.FG_GREEN + "ok" + woad.RESET)
```


## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/woad.Go "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/woad.Go.


### Dependencies

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
