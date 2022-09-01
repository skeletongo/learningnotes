package observer

import "fmt"

// 观察者模式
// 被观察者和观察者两种角色，被观察者的行为或状态通知给观察者

// 例如观察股票价格

// Observer 观察者接收消息接口
type Observer interface {
	Receiver(price int) // 接收价格变动
}

// Person 观察者
type Person struct {
	Name string
}

func (o *Person) Receiver(price int) {
	fmt.Println(o.Name, "收到价格变更：", price)
}

// BeObserver 被观察者接口
type BeObserver interface {
	Add(o Observer)
	Remove(o Observer)
	Notify()
}

// Stock 被观察者
type Stock struct {
	Price int
	Ps    []Observer
}

func (s *Stock) Add(o Observer) {
	for _, v := range s.Ps {
		if v == o {
			return
		}
	}
	s.Ps = append(s.Ps, o)
}

func (s *Stock) Remove(o Observer) {
	for k, v := range s.Ps {
		if v == o {
			s.Ps = append(s.Ps[:k], s.Ps[k+1:]...)
			return
		}
	}
}

func (s *Stock) Notify() {
	for _, v := range s.Ps {
		v.Receiver(s.Price)
	}
}
