package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/honwen/ip2loc"
)

func stdin() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ip := scanner.Text()
		if _, err := ip2loc.IP2loc(ip); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", ip, err)
		} else {
			fmt.Println(ip2loc.IP2locCHS(ip))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed:", err)
		os.Exit(1)
	}
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
