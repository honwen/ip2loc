//go:generate statik -src=./assets
//go:generate go fmt statik/statik.go

package ip2loc

import (
	"github.com/chenhw2/ip2loc/datx"
	_ "github.com/chenhw2/ip2loc/statik"
	"github.com/rakyll/statik/fs"
)

var city *datx.City

func IP2loc(ip string) (loc datx.Location, err error) {
	if nil == city {
		stkfs, err := fs.New()
		if err != nil {
			return loc, err
		}

		file, err := stkfs.Open("/17monipdb.datx")
		if err != nil {
			return loc, err
		}
		city, err = datx.NewCity(file) // For City Level IP Database
	}
	return city.FindLocation(ip)
}
