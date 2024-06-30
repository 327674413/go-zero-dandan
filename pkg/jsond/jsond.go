package jsond

func MustString(strPt *string) string {
	if strPt == nil {
		return ""
	} else {
		return *strPt
	}
}

// Zeroable 所有可以用0值表示的类型 ，～代表其衍生自定义类型
type Zeroable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func MustInt[T Zeroable](numStr *T) T {
	if numStr == nil {
		return 0
	} else {
		return *numStr
	}
}
