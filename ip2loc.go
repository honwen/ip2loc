package ip2loc

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/honwen/ip2loc/ipdb"
	"github.com/honwen/ip2loc/ipv6"
)

const regxIPv4 = `(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`

const regxIPv6 = `([0-9A-Fa-f]{0,4}:){2,7}([0-9A-Fa-f]{1,4}$|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`

//go:embed assets/qqwry.ipdb
var assets_qqwry []byte

//go:embed assets/ipv6wry.db
var assets_ipv6wry []byte

var (
	city *ipdb.City
	data *ipv6.Dat

	v4 *regexp.Regexp
	v6 *regexp.Regexp
)

func IP2loc(ip string) (loc *ipdb.CityInfo, err error) {
	// log.Println(ip)
	if nil == v4 {
		v4 = regexp.MustCompile(regxIPv4)
	}
	if v4.MatchString(ip) {
		if nil == city {
			city, err = ipdb.NewCityFromBytes(assets_qqwry) // For City Level IP Database
			if err != nil {
				return
			}
		}
		return city.FindInfo(ip, "CN")
	}
	if nil == v6 {
		v6 = regexp.MustCompile(regxIPv6)
	}
	if v6.MatchString(ip) {
		if nil == data {
			data = &ipv6.Dat{
				Data: assets_ipv6wry,
			}
			data.InitData()
		}
		info_v6 := data.Find(ip)
		info := ipdb.CityInfo{
			CountryName: info_v6.Country,
			RegionName:  info_v6.Area,
		}
		return &info, nil
	}
	return nil, fmt.Errorf("NOT IPv4 or IPv6, %+v", ip)
}

func IP2locCHS(ip string) (str string) {
	if loc, err := IP2loc(ip); err != nil {
		log.Printf("%+v", err)
	} else {
		str = fmt.Sprintf("[%s %s %s %s]", loc.CountryName, loc.RegionName, loc.CityName, loc.IspDomain)
		for strings.Contains(str, "\t") {
			str = strings.ReplaceAll(str, "\t", " ")
		}
		for strings.Contains(str, "  ") {
			str = strings.ReplaceAll(str, "  ", " ")
		}
		for strings.Contains(str, " ]") {
			str = strings.ReplaceAll(str, " ]", "]")
		}
		for strings.Contains(str, "  ") {
			str = strings.ReplaceAll(str, "  ", " ")
		}
	}
	return
}
