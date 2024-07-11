package main

import "fmt"

/**
优点：
(1) 它对客户端屏蔽了子系统组件，减少了客户端所需处理的对象数目，并使得子系统使用起来更加容易。通过引入外观模式，客户端代码将变得很简单，与之关联的对象也很少。
(2) 它实现了子系统与客户端之间的松耦合关系，这使得子系统的变化不会影响到调用它的客户端，只需要调整外观类即可。
(3) 一个子系统的修改对其他子系统没有任何影响。
缺点：
(1) 不能很好地限制客户端直接使用子系统类，如果对客户端访问子系统类做太多的限制则减少了可变性和灵活 性。
(2) 如果设计不当，增加新的子系统可能需要修改外观类的源代码，违背了开闭原则。
*/

type SubSystemA struct {
}

func (sa *SubSystemA) MethodA() {
	fmt.Println("子系统A的方法A")
}

type SubSystemB struct {
}

func (sb *SubSystemB) MethodB() {
	fmt.Println("子系统B的方法B")
}

type SubSystemC struct {
}

func (sc *SubSystemC) MethodC() {
	fmt.Println("子系统C的方法C")
}

// 外观模式，提供了一个外观类，简化成一个简单的接口供使用
type Facade struct {
	A *SubSystemA
	B *SubSystemB
	C *SubSystemC
}

func (f *Facade) MethodABC() {
	f.A.MethodA()
	f.B.MethodB()
	f.C.MethodC()
}

func main() {
	f := &Facade{
		A: new(SubSystemA),
		B: new(SubSystemB),
		C: new(SubSystemC),
	}
	f.MethodABC()
}
