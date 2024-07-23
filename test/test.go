package main

import (
	"fmt"
	"go-zero-dandan/common/resd"
)

func main() {
	_, err := test()
	fmt.Println("最后：", err)
}
func test() (str string, err error) {

	defer end(&err)
	return "123", nil //resd.NewErr(fmt.Sprintf("%v", "1"))
}
func end(err *error) {
	if rec := recover(); rec != nil {
		fmt.Println("panic", rec)
	}
	if *err != nil {
		fmt.Println("进err了")
	}
	*err = resd.NewErr(fmt.Sprintf("%v", "23"))
}
