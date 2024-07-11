package main

import "fmt"

// 抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 实现层
type Benz struct {
}

type Bmw struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running")
}

func (b *Bmw) Run() {
	fmt.Println("Bmw is running")
}

type Zs struct {
}

func (zs *Zs) Drive(car Car) {
	fmt.Println("zs drive car")
	car.Run()
}

// 业务层
func main() {
	var zs Driver
	zs = &Zs{}
	var benz Car
	benz = &Benz{}
	zs.Drive(benz)
}
