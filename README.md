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

## Main use case

The initial idea for this project is to ease connecting to a machine having an IPv6 stack. Today most of the
Linux distribution comes with a dual IP stack. So, even when working on a legacy IPv4 network, it' possible to
connect to any other machine using it's IPv6 link-local address. It can be usefull for instance to connected
to a machine that failed its DHCP request.

When you know the machine's MAC address, you can use `mac6` to contact a machine through it's IPv6 link-local
address.

__Examples__

```
$ ping6 $(mac6 -i eth0 8b:7a:1b:06:06:75)
```

is equivalent to

```
$ ping6 fe80::897a:1bff:fe06:0675%eth0
```

And

```
$ ssh user@$(mac6 -i eth0 8b:7a:1b:06:06:75)
```

is equivalent to

```
$ ssh user@fe80::897a:1bff:fe06:0675%eth0
```

## License

Mac6 is licensed under the [MIT license](LICENSE.md).
