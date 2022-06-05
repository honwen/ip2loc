package ip2loc

import (
	"embed"

	"github.com/honwen/ip2loc/ipdb"
)

//go:embed assets/qqwry.ipdb
var assets embed.FS

var city *ipdb.City

func IP2loc(ip string) (loc *ipdb.CityInfo, err error) {
	if nil == city {
		fs, _ := assets.Open("assets/qqwry.ipdb")
		city, err = ipdb.NewCity(fs) // For City Level IP Database
	}
	return city.FindInfo(ip, "CN")
}
