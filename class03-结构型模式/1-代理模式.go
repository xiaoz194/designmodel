package main

import "fmt"

type Goods struct {
	Kind string //商品的种类
	Fact bool   //商品的真伪
}

type Shopping interface {
	Buy(goods *Goods)
}

type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("在韩国购物，购买了: ", goods.Kind)
}

type AmericaShopping struct {
}

func (as *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("在韩国购物，购买了: ", goods.Kind)
}

type OverSeasProxy struct {
	Shopping Shopping
}

func NewOverSeasProxy(shopping Shopping) *OverSeasProxy {
	return &OverSeasProxy{Shopping: shopping}
}

// 代理
func (os *OverSeasProxy) Buy(goods *Goods) {
	// 1 + 先辨别真伪
	if os.distinguish(goods) == true {
		// 2 使用具体的代购的buy方法
		os.Shopping.Buy(goods)
		//3 + 海关安检
		os.check(goods)
	}
}

func (os *OverSeasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对", goods.Kind, "进行了鉴别")
	if goods.Fact == false {
		fmt.Println(goods.Kind, "是假货")
	}
	return goods.Fact
}

func (os *OverSeasProxy) check(goods *Goods) {
	fmt.Println("对", goods.Kind, "进行安检")
}

func main() {
	g1 := &Goods{
		Kind: "beans短袖",
		Fact: true,
	}
	g2 := &Goods{
		Kind: "Nike",
		Fact: false,
	}
	// 不用代理模式
	kshopping := &KoreaShopping{}
	// 验证真伪
	if g1.Fact == true {
		// 购买
		kshopping.Buy(g1)
		//安检
		fmt.Println("对", g1.Kind, "进行安检")
	}
	fmt.Println("----------------------------------")
	// 使用代理模式
	kproxy := NewOverSeasProxy(kshopping)
	kproxy.Buy(g1)
	aproxy := NewOverSeasProxy(&AmericaShopping{})
	aproxy.Buy(g2)

}
