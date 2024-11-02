package main

import (
	"fmt"
	"go-zero-dandan/common/utild"
)

func main() {
	_, err := utild.SendEmail("shengdanwanju@126.com", "dan5127642", "327674413@qq.com", "有新的订单了", "订单编号：111")
	if err != nil {
		fmt.Println(err)
	}
}
