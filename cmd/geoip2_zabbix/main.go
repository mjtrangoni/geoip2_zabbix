package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	externalip "github.com/glendc/go-external-ip"
	geoip2 "github.com/oschwald/geoip2-golang"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func parseCityRecord(cityStruct *geoip2.City, outLang *string) (resp GeoIPCityJSON) {
	resp.City.Name = cityStruct.City.Names[*outLang]
	resp.Continent.Code = cityStruct.Continent.Code
	resp.Continent.Name = cityStruct.Continent.Names[*outLang]
	resp.Country.IsInEuropeanUnion = cityStruct.Country.IsInEuropeanUnion
	resp.Country.IsoCode = cityStruct.Country.IsoCode
	resp.Country.Name = cityStruct.Country.Names[*outLang]
	resp.Location.AccuracyRadius = uint(cityStruct.Location.AccuracyRadius)
	resp.Location.Latitude = cityStruct.Location.Latitude
	resp.Location.Longitude = cityStruct.Location.Longitude
	resp.Location.MetroCode = cityStruct.Location.MetroCode
	resp.Location.TimeZone = cityStruct.Location.TimeZone
	resp.PostalCode = cityStruct.Postal.Code
	resp.Subdivisions.IsoCode = cityStruct.Subdivisions[0].IsoCode
	resp.Subdivisions.Name = cityStruct.Subdivisions[0].Names[*outLang]
	resp.Traits.IsAnonymousProxy = cityStruct.Traits.IsAnonymousProxy
	resp.Traits.IsSatelliteProvider = cityStruct.Traits.IsSatelliteProvider

	return resp
}

func main() {
	var (
		geoIP2Path = kingpin.Flag("path.geoipdb", "GeoIP DB file path.").Default("GeoLite2-City.mmdb").String()
		hostFQDN   = kingpin.Flag("host.fqdn", "Host FQDN for external IP.").Default("").String()
		outLang    = kingpin.Flag("out.lang", "Output language. [de,en,es,fr,ja,pt-BR,ru,zh-CN]").Default("en").String()
		verbose    = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
		ip         net.IP
	)

	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	db, err := geoip2.Open(*geoIP2Path)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	if *hostFQDN != "" {
		var ips []net.IP

		ips, err = net.LookupIP(*hostFQDN)
		if err != nil {
			log.Panic("Could not get IPs: ", err)
		}

		ip = ips[0]
	} else {
		// Get external IP from an external service
		consensus := externalip.DefaultConsensus(nil, nil)
		ip, err = consensus.ExternalIP()
		if err != nil {
			log.Panic(err)
		}
	}

	record, err := db.City(ip)
	if err != nil {
		log.Panic(err)
	}

	locationJSON := parseCityRecord(record, outLang)

	if *verbose {
		fmt.Printf("%+v", locationJSON)
	}

	locJSON, err := json.Marshal(locationJSON)
	if err != nil {
		log.Panic(err)
	}

	// Convert bytes to string.
	fmt.Println(string(locJSON))
}
