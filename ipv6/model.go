package ipv6

type Dat struct {
	Data     []byte
	Version  uint16
	FilePath string
	Index    Index
	Offset   uint32
}

type Index struct {
	Start  uint64
	End    uint64
	Offlen uint64
	Count  uint64
	Data   [][]uint64
}

type Result struct {
	IP      string `json:"ip"`
	Number  uint64 `json:"number"`
	Country string `json:"country"`
	Area    string `json:"area"`
	Offset  uint64 `json:"offset"`
}
