package downplatforms

const (
	LIANLUOYUN = "lianluoyun"
	VSIPER     = "vsiper"
)

var platformDesc = map[string]string{
	LIANLUOYUN: "联络云",
	VSIPER:     "平台标准",
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
