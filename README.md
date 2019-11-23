# GeoIP2 Zabbix [![Build Status](https://travis-ci.org/mjtrangoni/geoip2_zabbix.svg)][travis]

[![GoDoc](https://godoc.org/github.com/mjtrangoni/geoip2_zabbix?status.svg)](https://godoc.org/github.com/mjtrangoni/geoip2_zabbix)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/mjtrangoni/geoip2_zabbix/master/LICENSE)
[![StyleCI](https://github.styleci.io/repos/107779392/shield?branch=master)](https://github.styleci.io/repos/107779392)

Monitoring script for getting the actual host location from 
[GeoIP2-City](https://www.maxmind.com/en/geoip2-city) and expose it to [Zabbix](https://www.zabbix.com/).

## Getting

```
$ go get github.com/mjtrangoni/geoip2_zabbix
```

## Building

```
$ cd $GOPATH/src/github.com/mjtrangoni/geoip2_zabbix
$ make
```

## Configuration


## GeoIP2 DB

This program works with the commercial and the lite version of the 
**GeoIP2-City Database**. You can choose which one to use.

# Templates

WIP

## Running

### Manually
```
$ ./geoip2_zabbix <flags>
```

### Zabbix UserParameters

A [UserParameters](https://www.zabbix.com/documentation/current/manual/config/items/userparameters) example will be available soon.

### Zabbix loadable module

Once the golang version of `zabbix_agent`, called `zabbix_agent2` will be released,
I will be adding this capability.

## Dashboard

WIP

## Contributing

Refer to [CONTRIBUTING.md](https://github.com/mjtrangoni/geoip2_zabbix/blob/master/CONTRIBUTING.md)

## License

Apache License 2.0, see [LICENSE](https://github.com/mjtrangoni/geoip2_zabbix/blob/master/LICENSE).

[travis]: https://travis-ci.org/mjtrangoni/geoip2_zabbix
