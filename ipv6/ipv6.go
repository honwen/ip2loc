package ipv6

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

func (d *Dat) InitData() (err error) {
	// 0~3 字符串 "IPDB"
	if tag := d.Data[:4]; string(tag) != "IPDB" {
		return errors.New("文件格式错误")
	}

	// 6 byte 偏移地址长度(2~8)  3
	d.Index.Offlen = uint64(d.Data[6])
	// fmt.Printf("偏移地址长度: %d\n", d.Index.Offlen)

	// 7 byte IP地址长度(4或8或12或16, 现在只支持4(ipv4)和8(ipv6))  8
	// iplen := d.Data[7]
	// fmt.Printf("IP地址长度: %d\n", iplen)

	// 8~15 int64 记录数, 因为不可能为负数, Uint64应该也放的下
	d.Index.Count = binary.LittleEndian.Uint64(d.Data[8:16])
	// fmt.Printf("记录数: %d\n", d.Index.Count)

	// 16~23 int64 索引区第一条记录的偏移
	d.Index.Start = binary.LittleEndian.Uint64(d.Data[16:24])
	// fmt.Printf("索引区第一条记录的偏移: %d\n", d.Index.Start)

	d.Index.End = d.Index.Start + (d.Index.Count-1)*(d.Index.Offlen+8)
	// fmt.Printf("索引区最后一条记录的偏移: %d\n", d.Index.End)

	return nil
}

func (d *Dat) Dump() {
	for i := uint64(0); i < d.Index.Count; i++ {
		ip := d.GetUint64(d.Index.Start+i*(8+d.Index.Offlen), 8)
		r := d.GetUint64(d.Index.Start+i*(8+d.Index.Offlen)+8, d.Index.Offlen)
		cArea, aArea := d.getAddr(r)
		fmt.Printf("ip: %X cArea:%s aArea:%s\n", ip, cArea, aArea)
	}
}

func (d *Dat) GetLast() {
	// ip := d.GetUint64(d.Index.End, 8)
	r := d.GetUint64(d.Index.End+8, d.Index.Offlen)
	_, aArea := d.getAddr(r)
	fmt.Printf("信息:%s\n", aArea)
	// return res
}

func (d *Dat) Find(ip string) (res Result) {
	if ip == "" {
		return
	}

	ipv6 := net.ParseIP(ip)
	if ipv6 == nil {
		return
	}

	res.IP = fmt.Sprintf("%X", ipv6)
	r := binary.BigEndian.Uint64(ipv6[:8])
	res.Number = r
	o := d.SearchIndex(r, 0, d.Index.Count)
	res.Offset = d.GetUint64(o+8, d.Index.Offlen)
	res.Country, res.Area = d.getAddr(res.Offset)
	return res
}

func (d *Dat) SearchIndex(ip, l, r uint64) (offset uint64) {
	// 使用二分法查找网络字节编码的IP地址的索引记录
	// 返回索引位置
	if r-l <= 1 {
		return d.Index.Start + l*(8+d.Index.Offlen)
	}

	m := (l + r) / 2
	o := d.Index.Start + m*(8+d.Index.Offlen)
	new_ip := d.GetUint64(o, 8)
	if ip < new_ip {
		return d.SearchIndex(ip, l, m)
	} else {
		return d.SearchIndex(ip, m, r)
	}
}

func (d *Dat) getAddr(offset uint64) (string, string) {
	flag := d.GetUint64(offset, 1)
	if flag == 1 {
		// # 重定向模式1
		// # [IP][0x01][国家和地区信息的绝对偏移地址]
		// # 使用接下来的3字节作为偏移量调用字节取得信息
		return d.getAddr(d.GetUint64(offset+1, d.Index.Offlen))
	} else {
		// # 重定向模式2 + 正常模式
		// # [IP][0x02][信息的绝对偏移][...]
		cArea := d.getAreaAddr(offset)
		// fmt.Println(cArea)
		// return "", ""
		if flag == 2 {
			offset += 1 + d.Index.Offlen
		} else {
			offset = d.getOffset(offset) + 1
		}
		aArea := d.getAreaAddr(offset)
		return cArea, aArea
	}
}

func (d *Dat) getAreaAddr(offset uint64) string {
	flag := d.GetUint64(offset, 1)
	if flag == 1 || flag == 2 {
		p := d.GetUint64(offset+1, d.Index.Offlen)
		return d.getAreaAddr(p)
	} else {
		return d.getString(offset)
	}
}

func (d *Dat) getString(offset uint64) string {
	res := []byte{}
	for {
		if buf := d.Data[offset]; buf == 0 {
			break
		} else {
			res = append(res, buf)
			offset += 1
		}
	}
	return string(res)
}

func (d *Dat) getOffset(offset uint64) uint64 {
	for {
		if buf := d.Data[offset]; buf == 0 {
			return offset
		} else {
			offset += 1
		}
	}
}

func (d *Dat) GetUint64(offset, size uint64) uint64 {
	return byte2UInt64(d.Data[offset : offset+size])
}

func byte2UInt64(data []byte) uint64 {
	i := make([]byte, 8)
	//lint:ignore S1001 ignore this!
	for j := 0; j < len(data); j++ {
		i[j] = data[j]
	}

	return binary.LittleEndian.Uint64(i)
}
