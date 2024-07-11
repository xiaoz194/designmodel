package main

import "fmt"

/*
*
简单工厂模式 + 开闭原则 = 工厂模式

*
*/
type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit()
}

type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("apple")
}

type Banana struct {
}

func (banana *Banana) Show() {
	fmt.Println("banana")
}

type AppleFactory struct {
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var f Fruit
	f = new(Apple)
	return f
}

type BananaFactory struct{}

func (fac *BananaFactory) CreateFruit() Fruit {
	var f Fruit
	f = new(Banana)
	return f
}
func main() {
	appleFac := &AppleFactory{}

	apple := appleFac.CreateFruit()
	apple.Show()
	bananaFac := &BananaFactory{}

	banana := bananaFac.CreateFruit()
	banana.Show()
}
