package main

import "fmt"

// 开闭原则: 类的改动是通过增加代码执行的，而不是修改源代码. 在不破坏系统的情况下增加内容

// 抽象的银行业务员
type AbstractBanker interface {
	DoBusi() // 抽象的处理业务接口
}

type SaveBanker struct {
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款业务")
}

// +
type TransferBanker struct {
}

// 实现一个架构层
func (sb *TransferBanker) DoBusi() {
	fmt.Println("进行了转账业务")
}

func BankBussiness(banker AbstractBanker) {
	banker.DoBusi()
}

func main() {
	BankBussiness(&SaveBanker{})
	BankBussiness(&TransferBanker{})
}
