### Source

- https://github.com/honwen/ip2loc

### Thanks

- https://github.com/metowolf/qqwry.ipdb
- https://github.com/ipipdotnet/ipdb-go

### Usage

```go
package main

import (
        "fmt"
        "github.com/honwen/ip2loc"
)

func main() {
        if loc, err := ip2loc.IP2loc("8.8.8.8"); err != nil {
                fmt.Printf("%+v", err)
        } else {
                fmt.Printf("%+v", loc)
        }
}

```
