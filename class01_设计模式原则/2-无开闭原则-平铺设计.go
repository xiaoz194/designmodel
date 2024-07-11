package main

import "fmt"

type Banker struct{}

// 存款业务
func (b *Banker) Save() {
	fmt.Println("进行了 存款业务")
}

// 转账业务
func (b *Banker) Transfer() {
	fmt.Println("进行了 存款业务")
}

// 支付业务
func (b *Banker) Pay() {
	fmt.Println("进行了 存款业务")
}

// 股票业务（新增+）
// 直接在类中增加方法，违背开闭原则
func (b *Banker) Stock() {
	fmt.Println("进行了 股票业务")
}

func main() {
	b := &Banker{}

	b.Save()
	b.Transfer()
	b.Pay()
}
