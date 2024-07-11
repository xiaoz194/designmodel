package main

import "fmt"

// 抽象的观察者
type Listener interface {
	Active() //观察者得到通知后要触发的具体动作
}

// 抽象的通知者（被观察者）
type Notifier interface {
	AddListener()    // 添加观察者
	RemoveListener() // 删除观察者
	Notify()         // 通知所有观察者
}

// 实现层
// 具体的观察者
type ListenerA struct {
	DoSth string
}

func (a *ListenerA) Do() {
	fmt.Println("listener A do: ", a.DoSth)
}

func (a *ListenerA) Active() {
	fmt.Println("stop do: ,", a.DoSth)
}

type ListenerB struct {
	DoSth string
}

func (b *ListenerB) Do() {
	fmt.Println("listener B do: ", b.DoSth)
}

func (b *ListenerB) Active() {
	fmt.Println("continue do: ,", b.DoSth)
}

// 具体的通知者
type NotifierA struct {
	listenerList []Listener //需要通知的全部观察者集合
}

func (na *NotifierA) AddListener(listener Listener) {
	na.listenerList = append(na.listenerList, listener)
}

func (na *NotifierA) RemoveListener(listener Listener) {
	for idx, l := range na.listenerList {
		// 找到要删除的位置
		if listener == l {
			// 将删除的元素的前后位置连接起来
			na.listenerList = append(na.listenerList[:idx], na.listenerList[idx+1:]...)
			break
		}
	}
}

func (na *NotifierA) Notify() {
	for _, l := range na.listenerList {
		l.Active()
	}
}

func main() {
	l1 := &ListenerA{
		DoSth: "doA",
	}
	l2 := &ListenerB{
		DoSth: "doB",
	}

	n := new(NotifierA)
	n.AddListener(l1)
	n.AddListener(l2)
	fmt.Println("before notify...")
	l1.Do()
	l2.Do()
	fmt.Println("after notify...")
	n.Notify()

}
