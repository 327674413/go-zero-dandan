package resd

type SuccInfo struct {
	Result bool `json:"result"`
	Code   int  `json:"code"`
	Data   any  `json:"data"`
}

func Succ(data any) *SuccInfo {
	return &SuccInfo{Result: true, Code: Ok, Data: data}
}
func SuccAsync(data any) *SuccInfo {
	return &SuccInfo{Result: true, Code: OkAsync, Data: data}
}
