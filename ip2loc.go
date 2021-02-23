//go:generate statik -src=./assets
//go:generate go fmt statik/statik.go

package ip2loc

import (
	"github.com/honwen/ip2loc/ipdb"
	_ "github.com/honwen/ip2loc/statik"
	"github.com/rakyll/statik/fs"
)

var city *ipdb.City

func IP2loc(ip string) (loc *ipdb.CityInfo, err error) {
	if nil == city {
		stkfs, err := fs.New()
		if err != nil {
			return loc, err
		}

		file, err := stkfs.Open("/qqwry.ipdb")
		if err != nil {
			return loc, err
		}
		city, err = ipdb.NewCity(file) // For City Level IP Database
	}
	return city.FindInfo(ip, "CN")
}
