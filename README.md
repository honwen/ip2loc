### Source

- https://github.com/chenhw2/ip2loc

### Thanks

- https://github.com/metowolf/qqwry.ipdb
- https://github.com/ipipdotnet/datx-go
- https://github.com/rakyll/statik

### Usage

```go
package main

import (
        "fmt"
        "github.com/chenhw2/ip2loc"
)

func main() {
        if loc, err := ip2loc.IP2loc("8.8.8.8"); err != nil {
                fmt.Printf("%+v", err)
        } else {
                fmt.Printf("%+v", loc)
        }
}

```

### Assets

```bash
go get github.com/rakyll/statik
statik -src=assets -f
```
