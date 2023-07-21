package resd

import "fmt"

type FailInfo struct {
	Result bool     `json:"result"`
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Temps  []string `json:"-"`
}

func (t *FailInfo) Error() string {
	return fmt.Sprintf("%s", t.Msg)
}
func Fail(msg string, errCode int, tempStr ...string) *FailInfo {
	res := &FailInfo{
		Result: false,
		Code:   errCode,
		Msg:    msg,
	}
	if len(tempStr) > 0 {
		res.Temps = tempStr
	}
	return res
}
