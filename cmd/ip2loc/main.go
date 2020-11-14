package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chenhw2/ip2loc"
)

func stdin() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ip := scanner.Text()
		if loc, err := ip2loc.IP2loc(ip); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", ip, err)
		} else {
			fmt.Println(ip, loc.CountryName, loc.RegionName, loc.CityName, loc.IspDomain)
			// fmt.Printf("%+v\n", loc)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed:", err)
		os.Exit(1)
	}

	return
}

func main() {
	if len(os.Args) > 1 {
		if loc, err := ip2loc.IP2loc(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[1], err)
			os.Exit(1)
		} else {
			fmt.Println(loc.CountryName, loc.RegionName, loc.CityName, loc.IspDomain)
		}
	} else {
		stdin()
	}
}
