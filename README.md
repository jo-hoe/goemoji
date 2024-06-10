# Go Emoji

[![Test Status](https://github.com/jo-hoe/goemoji/workflows/test/badge.svg)](https://github.com/jo-hoe/goemoji/actions?workflow=test)
[![Lint Status](https://github.com/jo-hoe/goemoji/workflows/lint/badge.svg)](https://github.com/jo-hoe/goemoji/actions?workflow=lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/goemoji)](https://goreportcard.com/report/github.com/jo-hoe/goemoji)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/goemoji/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/goemoji?branch=main)

Adds emojis to strings.

## Development Setup

### Pre-Requisites

- [Golang](https://go.dev/doc/install)

#### Optional

You can use `make` to enhance development on this project. `make` is typically installed out of the box on Linux and Mac.

- [make](https://www.gnu.org/software/make/)

If you do not have it and run on Windows, you can directly install it from [gnuwin32](https://gnuwin32.sourceforge.net/packages/make.htm) or via `winget`

```PowerShell
winget install GnuWin32.Make
```

Run `make help` to discover the commands you can use.

## Linting

The project used `golangci-lint` for linting.

### Installation

<https://golangci-lint.run/welcome/install/>

### Run Linting

Run the linting locally by executing.

```cli
golangci-lint run ./...
```
