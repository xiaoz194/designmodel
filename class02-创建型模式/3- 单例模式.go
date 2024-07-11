package main

import "sync"

// singleton 是一个单例结构体
type singleton struct {
	// 单例的状态或配置
}

var instance *singleton
var once sync.Once

// GetInstance 返回单例对象的引用
func GetInstance() *singleton {
	if instance == nil {
		once.Do(func() {
			instance = &singleton{}
			// 在这里初始化单例的其他资源
		})
	}
	return instance
}

func main() {
	GetInstance()
}
