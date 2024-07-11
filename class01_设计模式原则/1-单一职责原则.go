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
