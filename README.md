# GeoIP2 Zabbix ![Build Status](https://github.com/mjtrangoni/geoip2_zabbix/actions/workflows/build.yml/badge.svg?branch=master)

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=for-the-badge)](https://raw.githubusercontent.com/mjtrangoni/geoip2_zabbix/master/LICENSE)

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

This program works with the commercial and the lite version of the **GeoIP2-City Database**. You can choose which one to use.

# Templates

For Zabbix 4.2 [here](https://github.com/mjtrangoni/geoip2_zabbix/blob/master/templates/zbx_template_geoip2zabbix_zabbix42.xml)

## Running

### Manually
```
$ ./geoip2_zabbix <flags>
```

### Zabbix UserParameters

Using zabbix [UserParameters](https://www.zabbix.com/documentation/current/manual/config/items/userparameters),
you can call **geoip2_zabbix** from linux or windows,

#### Linux

You can build **geoip2_zabbix** for linux running simply `make`, and then add
this UserParameter to your `zabbix-agent`.

```
UserParameter=geoip2.zabbix,/pathtobinary/geoip2_zabbix --path.geoipdb=/pathtodb/GeoIPCityV2.mmdb
```

#### Windows

You can crossbuild **geoip2_zabbix** for windows on linux running simply `make crossbuild`, and then add
this UserParameter to your `zabbix-agent`.

```
UserParameter=geoip2.zabbix,PowerShell -NoProfile -ExecutionPolicy Bypass -Command " & 'C:\PathToBinary\geoip2_zabbix.exe' --path.geoipdb=C:\PathToDB\GeoIPCityV2.mmdb"
```

## Grafana Dashboard

WIP

## Contributing

Refer to [CONTRIBUTING.md](https://github.com/mjtrangoni/geoip2_zabbix/blob/master/CONTRIBUTING.md)

## License

Apache License 2.0, see [LICENSE](https://github.com/mjtrangoni/geoip2_zabbix/blob/master/LICENSE).

[travis]: https://travis-ci.org/mjtrangoni/geoip2_zabbix
