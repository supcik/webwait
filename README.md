# webwait

A simple CLI tool to wait for a web server to become available.

![GitHub Release](https://img.shields.io/github/v/release/supcik/webwait)
[![build-and-release](https://github.com/supcik/webwait/actions/workflows/build-release.yml/badge.svg)](https://github.com/supcik/webwait/actions/workflows/build-release.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/supcik/webwait)
[![Go Reference](https://pkg.go.dev/badge/github.com/supcik/webwait.svg)](https://pkg.go.dev/github.com/supcik/webwait)
![Static Badge](https://img.shields.io/badge/Made_in-Switzerland_%2B-DA291C)

## Overview

`webwait` repeatedly checks a given URL until it responds with HTTP 200 OK or a timeout is reached. This is useful for scripts and automation where you need to wait for a service to be up before proceeding.

## Usage

```sh
webwait <url> [flags]
```

### Arguments

- `<url>`: The URL to check (e.g., `http://localhost:8080/health`)

### Flags

- `-t, --timeout <duration>`: Timeout duration for waiting (default: `1m`)
- `-i, --interval <duration>`: Minimum interval between checks (default: `10s`)

Example:

```sh
webwait http://localhost:8080/health -t 2m -i 5s
```

## Installation

### From Binaries

Install using the binaries from the [releases page](https://github.com/supcik/webwait/releases)

### Homebrew

To install `webwait` using Homebrew, run the following command:

```sh
brew install supcik/tap/webwait
```

You can also add Jacques Supcik's _tap_ and then simply install `webwait`:

```sh
brew tap supcik/tap
brew install webwait
```

### Using `go install`

To install `webwait` using `go install`, run the following command:

```sh
go install github.com/supcik/webwait@latest
```

### Build from Source

If you prefer to build from source, ensure you have Go installed and run the following commands:

```sh
git clone https://github.com/supcik/webwait.git
cd webwait
go build -o webwait
```

## License

This project is licensed under the MIT license. See the `LICENSES/` directory for details.

## Copyright

2025 Jacques Supcik <jacques.supcik@hefr.ch>
