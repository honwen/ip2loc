package ip2loc

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/honwen/ip2loc/ipdb"
)

//go:embed assets/qqwry.ipdb
var assets []byte

var city *ipdb.City

func IP2loc(ip string) (loc *ipdb.CityInfo, err error) {
	if nil == city {
		city, err = ipdb.NewCityFromBytes(assets) // For City Level IP Database
		if err != nil {
			return
		}
	}
	return city.FindInfo(ip, "CN")
}

func IP2locCHS(ip string) (str string) {
	if strings.Count(ip, `.`) < 3 {
		return
	}
	if loc, err := IP2loc(ip); err != nil {
		log.Printf("%+v", err)
	} else {
		str = fmt.Sprintf("[%s %s %s %s]", loc.CountryName, loc.RegionName, loc.CityName, loc.IspDomain)
		for strings.Contains(str, " ]") {
			str = strings.ReplaceAll(str, " ]", "]")
		}
		for strings.Contains(str, "  ") {
			str = strings.ReplaceAll(str, "  ", " ")
		}
	}
	return
}
