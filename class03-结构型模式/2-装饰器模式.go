package main

import "fmt"

type Phone interface {
	Show()
}

type Decorator struct {
	Phone Phone
}

func (d *Decorator) Show() {
}

type Huawei struct {
}

func (hw *Huawei) Show() {
	fmt.Println("huawei")
}

type Xiaomi struct {
}

func (hw *Xiaomi) Show() {
	fmt.Println("xiaomi")
}

type MoDecorator struct {
	Decorator
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{Phone: phone}}
}

func (md *MoDecorator) Show() {
	md.Phone.Show()
	fmt.Println("贴膜的手机") // 装饰器额外装饰的功能
}

type KeDecorator struct {
	Decorator
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{Phone: phone}}
}

func (kd *KeDecorator) Show() {
	kd.Phone.Show()
	fmt.Println("装壳的手机") // 装饰器额外装饰的功能
}

func main() {
	huawei := &Huawei{}
	huawei.Show()
	fmt.Println("====================")
	moHuawei := NewMoDecorator(huawei)
	moHuawei.Show()

	fmt.Println("====================")
	keMoHuawei := NewKeDecorator(moHuawei)
	keMoHuawei.Show()
}
