package main

import (
	"go-zero-dandan/common/resd"
)

const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorReset   = "\033[0m"
)

var res *resd.Resp

func main() {
	res = resd.NewResp(nil, "dev")
	err := f1()
	a, _ := resd.AssertErr(err)
	if a != nil {

	}
}
func f1() error {
	return res.Error(f2())
}
func f2() error {
	err := res.Error(f3())
	return err
}
func f3() error {
	err := res.Error(f4())
	return err
}
func f4() error {
	return res.NewError("aaa")
}
