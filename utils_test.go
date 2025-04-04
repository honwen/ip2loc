package ip2loc

import (
	"testing"

	"github.com/ysmood/got"
)

func TestIp2locCN_Loop(t *testing.T) {
	for i := 1; i < 99999; i++ {
		// got.T(t).Eq(IP2locCHS("202.96.128.86"), "[中国 广东 广州 电信]")
		// got.T(t).Eq(IP2locCHS("2408:8888::8"), "[中国 中国联通DNS服务器]")
	}
}

func TestIp2locCN(t *testing.T) {
	got.T(t).Eq(IP2locCHS(""), "")

	got.T(t).Eq(IP2locCHS("2408:8888::8"), "[中国 中国联通DNS服务器]")
	got.T(t).Eq(IP2locCHS("240e:5b::6666"), "[中国 江苏省 中国电信业务平台]")
	got.T(t).Eq(IP2locCHS("2409:8028:2000::1111"), "[中国 浙江省 中国移动DNS服务器]")

	got.T(t).Eq(IP2locCHS("2400:3200::1"), "[中国 浙江省 杭州市 阿里云计算有限公司]")
	got.T(t).Eq(IP2locCHS("2400:da00::6666"), "[中国 北京市 北京百度网讯科技有限公司]")

	got.T(t).Eq(IP2locCHS("2001:cc0:2fff:1::6666"), "[中国 北京市 中国科技网网络中心]")
	got.T(t).Eq(IP2locCHS("2001:da8::666"), "[中国 北京市 教育网(CERNET2)清华大学公共DNS服务器]")

	got.T(t).Eq(IP2locCHS("2001:4860:4860::8888"), "[全球 Google Inc 服务器网段 (Anycast)]")
	got.T(t).Eq(IP2locCHS("2606:4700:4700::1111"), "[全球 Cloudflare Inc Anycast网段]")
	got.T(t).Eq(IP2locCHS("2620:74:1b::1:1"), "[美国 Virginia州 Sterling NeuStar Inc]")
	got.T(t).Eq(IP2locCHS("2620:fe::fe"), "[全球 Packet Clearing House 公共DNS服务器]")

	got.T(t).Eq(IP2locCHS("202.96.128.86"), "[中国 广东 广州 电信]")
	got.T(t).Eq(IP2locCHS("202.96.209.133"), "[中国 上海 上海 电信]")
	got.T(t).Eq(IP2locCHS("219.141.136.10"), "[中国 北京 北京 电信]")

	got.T(t).Eq(IP2locCHS("210.22.70.3"), "[中国 上海 上海 联通]")
	got.T(t).Eq(IP2locCHS("123.123.123.123"), "[中国 北京 北京 联通]")

	got.T(t).Eq(IP2locCHS("223.87.238.22"), "[中国 四川 成都 移动]")

	got.T(t).Eq(IP2locCHS("101.6.6.6"), "[中国 北京 北京 教育网]")

	got.T(t).Eq(IP2locCHS("168.95.1.1"), "[中国 台湾 中华电信(HiNet)数据中心DNS服务器]")
	got.T(t).Eq(IP2locCHS("202.67.240.222"), "[中国 香港 HKnet公司]")

	got.T(t).Eq(IP2locCHS("203.189.136.148"), "[柬埔寨]")
	got.T(t).Eq(IP2locCHS("203.112.2.4"), "[日本 UCOM公司MIINET网络DNS服务器]")
	got.T(t).Eq(IP2locCHS("80.80.80.80"), "[荷兰 Freenom公司公共DNS服务器]")
	got.T(t).Eq(IP2locCHS("74.82.42.42"), "[美国 加利福尼亚州 阿拉梅达 Hurricane_Electric公司]")
}
