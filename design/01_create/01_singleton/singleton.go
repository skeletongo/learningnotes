package singleton

import "sync"

// 单例模式

type Singleton struct {
}

var instance *Singleton
var once sync.Once

func GetSingleton() *Singleton {
	once.Do(func() { // 并发安全，只执行一次
		instance = new(Singleton)
	})
	return instance
}
