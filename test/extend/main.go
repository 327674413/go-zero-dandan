package main

import "fmt"

// 定义一个接口，包含 Pay 方法
type Payer interface {
	Pay()
}

type Parent struct {
	Code string
	name string
	Payer
}

func (t *Parent) CreateOrder() {
	fmt.Println("Parent的CreateOrder，Code：", t.Code, "，name：", t.name)
	t.Payer.Pay()
}

func (t *Parent) Pay() {
	fmt.Println("Parent的Pay，Code：", t.Code, "，name：", t.name)
}

type Child1 struct {
	*Parent
}

func (t *Child1) Pay() {
	fmt.Println("Child1的Pay，Code：", t.Code, "，name：", t.name)
}

func main() {
	// 创建父类实例
	parent := &Parent{Code: "parent", name: "parent"}

	// 创建子类实例，并将 Payer 接口设置为子类的实例
	child := &Child1{parent}
	child.Payer = child

	// 调用父类的方法
	child.CreateOrder()
}
