package main

import "fmt"

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
