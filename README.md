# Mac6

[![Build Status](https://api.travis-ci.org/ggiamarchi/mac6.png?branch=master)](https://travis-ci.org/ggiamarchi/mac6)

`mac6` is a simple CLI that compute a link local IPv6 address from a given MAC address.

## Usage

```
Usage: mac6 [ -i=<interface> ] MAC

Compute IPv6 link-local address from MAC address

Arguments:
  MAC               Mac address

Options:
  -i, --interface   Out interface to build IPv6 link local string
```

## Examples

1. Getting the link-local IPv6 address from a MAC address

```bash
$ mac6 8b:7a:1b:06:06:75
fe80::897a:1bff:fe06:0675
```

2. Getting the link-local IPv6 address from a MAC address with the out interface

```bash
$ mac6 -i eth0 8b:7a:1b:06:06:75
fe80::897a:1bff:fe06:0675%eth0
```

## License

Mac6 is licensed under the [MIT license](LICENSE.md).
