package api

import "fmt"

type CodeMsg struct {
	Result bool   `json:"result"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func (c *CodeMsg) Error() string {
	return fmt.Sprintf("result:%v, code: %d, msg: %s", c.Result, c.Code, c.Msg)
}

func Fail(msg string, code ...int) error {
	apiCode := 400
	if len(code) > 0 {
		apiCode = code[0]
	}
	return &CodeMsg{Result: false, Code: apiCode, Msg: msg}
}
