package platforms

const (
	HUAWEI          = "huawei"
	WEILIAO         = "weiliao"
	XIANGYUN        = "xiangyun"
	XINGFANG        = "xingfang"
	JUNYUE          = "junyue"
	HZUNICOM        = "hzunicom"
	RONGLIAN        = "ronglian"
	XIAORONG        = "xiaorong"
	XIAORONGSOPEN   = "xiaorongsopen"
	DIXINGTONG      = "dixingtong"
	ZHIBOVOCIE      = "zhibovoice"
	SXUNICOM        = "shanxiunicom"
	CHIYAN          = "chiyan"
	SHOUHUI         = "shouhui"
	ALIBABA         = "alibaba"
	HBUNICOM        = "hbunicom"
	DIXINGTONGNEW   = "dixingtongnew"
	GXUNICOM        = "gxunicom"
	ZJUNICOM        = "zjunicom"
	ZJUNICOMSOPEN   = "zjunicomsopen"
	ZJUNICOMBD      = "zjunicombd"
	ZJUNICOMWL      = "zjunicomwl"
	WUXIMOBILE      = "wuximobile"
	WUXIMOBILESOPEN = "wuximobilesopn"
	ZJMOBILE        = "zjmobile"
	DXKCANH         = "dxkcanh"
	ZGDONGXING      = "zgdongxing"
	HAOBAI          = "HaoBai"
	JIANGSHUMOBILE  = "jiangshumobile"
	QUANZHOUUNICOM  = "quanzhouunicom"
	TENCENTCLOUD    = "tencentcloud"
	ZJUNICOMXIAOYI  = "zjunicomxiaoyi"
	ZHONGLIAN       = "zhonglian"
)

var platformDesc = map[string]string{
	HUAWEI:          "华为",
	WEILIAO:         "微聊",
	XIANGYUN:        "祥云",
	XINGFANG:        "新方",
	JUNYUE:          "君悦",
	HZUNICOM:        "杭州联通",
	RONGLIAN:        "容联",
	XIAORONG:        "小荣",
	XIAORONGSOPEN:   "小荣(sopen)",
	DIXINGTONG:      "迪信通",
	ZHIBOVOCIE:      "智博双呼",
	SXUNICOM:        "陕西联通(小西)",
	CHIYAN:          "赤焰",
	SHOUHUI:         "首汇(电信)",
	ALIBABA:         "阿里小号",
	HBUNICOM:        "河北联通",
	DIXINGTONGNEW:   "迪信通(新)",
	GXUNICOM:        "广西联通",
	ZJUNICOM:        "浙江联通",
	ZJUNICOMSOPEN:   "浙江联通(SOPEN)",
	ZJUNICOMWL:      "浙江联通(微聊)",
	ZJUNICOMBD:      "浙江联通(微聊本地)",
	WUXIMOBILE:      "无锡移动(暂停)",
	WUXIMOBILESOPEN: "无锡移动(SOPEN)",
	ZJMOBILE:        "浙江移动",
	DXKCANH:         "东信昆辰ANH",
	ZGDONGXING:      "中国东信",
	HAOBAI:          "号百(电信)",
	JIANGSHUMOBILE:  "江苏(移动)",
	QUANZHOUUNICOM:  "福建泉州(联通)",
	TENCENTCLOUD:    "腾讯云",
	ZJUNICOMXIAOYI:  "浙江联通(小易)",
	ZHONGLIAN:       "中联",
}

func Desc(code string) string {

	s, ok := platformDesc[code]
	if !ok {
		return code
	}

	return s
}

func All() map[string]string {
	return platformDesc
}
