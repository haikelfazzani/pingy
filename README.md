<div align="center">

# Pingy

A simple ICMP Echo implementation, based on golang.org/x/net/icmp.

![Language:Go](https://img.shields.io/static/v1?label=Language&message=Go&color=blue&style=flat-square)
![License:MIT](https://img.shields.io/static/v1?label=License&message=MIT&color=blue&style=flat-square)
[![Latest Release](https://img.shields.io/github/v/release/haikelfazzani/pingy?style=flat-square)](https://github.com/haikelfazzani/pingy/releases/latest)

</div>

## Usage

Simply specify the target host name or IP address in the first argument e.g.
`pingy github.com` or `pingy 8.8.8.8`.

```
Usage:
  pingy [OPTIONS] HOST

Application Options:
  -u        Url to ping
  -a        IP address to ping
  -t        Wait for the response from the remote host timeout (default: 10)
  -v        Show version

Help Options:
  -h, --help        Show this help message
```

## Installation

### Download executable binaries

You can download executable binaries from the latest release page.

> [![Latest Release](https://img.shields.io/github/v/release/haikelfazzani/pingy?style=flat-square)](https://github.com/haikelfazzani/pingy/releases/latest)

### Build from source

To build from source, clone this repository then run `make build` or
`go install`. Develo_ping_ on `go1.2.0 linux/amd64`.

Another way to install it if you have go in your machine just:

```sh
go install github.com/haikelfazzani/pingy@latest
```

## LICENSE

[MIT](./LICENSE)

## Author

[haikelfazzani](https://github.com/haikelfazzani)
