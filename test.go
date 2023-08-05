package main

import (
	"errors"
	"fmt"
	"go-zero-dandan/common/resd"
)

func main() {
	fmt.Println(resd.Error(errors.New("123")))
}
