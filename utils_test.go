package ip2loc

import (
	"testing"

	"github.com/ysmood/got"
)

func TestIp2locCN_Loop(t *testing.T) {
	for i := 1; i < 99999; i++ {
		got.T(t).Eq(IP2locCHS("202.96.128.86"), "[中国 广东 广州 电信]")
	}
}

func TestIp2locCN(t *testing.T) {
	got.T(t).Eq(IP2locCHS(""), "")
	got.T(t).Eq(IP2locCHS("202.96.128.86"), "[中国 广东 广州 电信]")
	got.T(t).Eq(IP2locCHS("202.96.209.133"), "[中国 上海 上海 电信]")
	got.T(t).Eq(IP2locCHS("219.141.136.10"), "[中国 北京 北京 电信]")

	got.T(t).Eq(IP2locCHS("210.22.70.3"), "[中国 上海 上海 联通]")
	got.T(t).Eq(IP2locCHS("123.123.123.123"), "[中国 北京 北京 联通]")

	got.T(t).Eq(IP2locCHS("223.87.238.22"), "[中国 四川 成都 移动]")

	got.T(t).Eq(IP2locCHS("101.6.6.6"), "[中国 北京 北京 教育网]")

	got.T(t).Eq(IP2locCHS("168.95.1.1"), "[中国 台湾 中华电信]")
	got.T(t).Eq(IP2locCHS("202.67.240.222"), "[中国 香港]")

	got.T(t).Eq(IP2locCHS("203.189.136.148"), "[柬埔寨 柬埔寨]")
	got.T(t).Eq(IP2locCHS("203.112.2.4"), "[日本 日本]")
	got.T(t).Eq(IP2locCHS("80.80.80.80"), "[荷兰 荷兰]")
	got.T(t).Eq(IP2locCHS("74.82.42.42"), "[美国加利福尼亚州弗里蒙特市 Hurricane Electric公司]")

}
