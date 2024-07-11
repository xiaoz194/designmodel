package main

import "fmt"

/*
*
简单工厂模式的优缺点
优点：实现了对象创建和使用的分离
缺点：新增类会修改简单工厂的创建方法，违背开闭原则

*
*/
type Fruit interface {
	Show()
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

type Factory struct{}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var f Fruit
	if kind == "apple" {
		f = new(Apple)
	} else if kind == "banana" {
		f = new(Banana)
	}
	return f
}

func main() {
	factory := &Factory{}
	apple := factory.CreateFruit("apple")
	apple.Show()
	banana := factory.CreateFruit("banana")
	banana.Show()
}
