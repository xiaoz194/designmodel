package main

import "fmt"

// 理解一对概念： 产品族 和 产品结构
// 如： 产品族 苹果，华为，小米
// 产品结构  手机 平板 电脑
//{苹果手机，苹果平板，苹果电脑} , {华为手机，华为平板，华为电脑} ， {小米手机，小米平板，小米电脑} 三个产品族
// 以产品族为一个工厂 生产对象

type Phone interface {
	ShowPhone()
}

type IPhone struct{}

func (iphone *IPhone) ShowPhone() {
	fmt.Println("product iphone")
}

type HuaweiMete struct{}

func (huaweiMete *HuaweiMete) ShowPhone() {
	fmt.Println("product huaweiMete")
}

type RedMi struct{}

func (redMi *RedMi) ShowPhone() {
	fmt.Println("product redMi")
}

type Pad interface {
	ShowIpad()
}

type IPad struct{}

func (ipad *IPad) ShowIpad() {
	fmt.Println("product ipad")
}

type HuaweiPad struct{}

func (huaweiPad *HuaweiPad) ShowIpad() {
	fmt.Println("product huaweiPad")
}

type MiPad struct{}

func (miPad *MiPad) ShowIpad() {
	fmt.Println("product miPad")
}

type AbstractFactory interface {
	CreatePhone()
	CreatePad()
	// computer ...
}

type AppleFactory struct{}

func (af *AppleFactory) CreatePhone() Phone {
	return &IPhone{}
}

func (af *AppleFactory) CreatePad() Pad {
	return &IPad{}
}

type HuaweiFactory struct{}

func (ah *HuaweiFactory) CreatePhone() Phone {
	return &HuaweiMete{}
}

func (af *HuaweiFactory) CreatePad() Pad {
	return &HuaweiPad{}
}

type MiFactory struct{}

func (ah *MiFactory) CreatePhone() Phone {
	return &RedMi{}
}

func (af *MiFactory) CreatePad() Pad {
	return &MiPad{}
}

func main() {
	appleFactory := new(AppleFactory)
	ipad := appleFactory.CreatePad()
	iphone := appleFactory.CreatePhone()
	ipad.ShowIpad()
	iphone.ShowPhone()

	xiaomiFactory := new(MiFactory)
	miPad := xiaomiFactory.CreatePad()
	miPad.ShowIpad()
}
