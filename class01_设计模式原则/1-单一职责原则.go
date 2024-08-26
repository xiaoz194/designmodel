package main

import "fmt"

// // 穿衣服的方式
// type Clothes struct {
// 	//逛街穿衣服
// }

// func (c *Clothes) OnWork() {
// 	fmt.Println("工作的装扮")
// }

// func (c *Clothes) OnShop() {
// 	fmt.Println("逛街的装扮")
// }
// func main() {
// 	c := Clothes{}
// 	c.OnWork()
// 	c.OnShop()
// }

// 单一职责原则：类的职责尽可能单一，对外只提供一种功能。 当然这是理想化的，实际开发中很难做到一个类之对应一种功能，这样的话会产生很多的类。
// 但是这个原则的思想是告诉我们设计类时尽可能的让其职责单一,而不是一个类冗余很多方法

type ClothesShop struct{}

func (cs *ClothesShop) Style() {
	fmt.Println("逛街的装扮")
}

type ClothesWork struct{}

func (cw *ClothesWork) Style() {
	fmt.Println("工作的装扮")
}

func main() {
	//工作的业务
	cw := ClothesWork{}
	cw.Style()
	//逛街的业务
	cs := ClothesShop{}
	cs.Style()
}
