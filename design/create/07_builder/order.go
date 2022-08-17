package builder

import (
	"bytes"
	"fmt"
)

// PriceMap 饮料单价
var PriceMap = map[string]float32{
	"橙汁": 2,
	"牛奶": 5,
}

// Drinks 饮料
type Drinks struct {
	Name string
}

// Order 订单
type Order struct {
	Array []Drinks // 购买的饮料
	Total float32  // 总金额
}

func (o Order) String() string {
	buf := bytes.NewBufferString("")
	for _, v := range o.Array {
		buf.WriteString(v.Name)
		buf.WriteString(",")
	}
	buf.WriteString(fmt.Sprintf("%.2f", o.Total))
	return buf.String()
}

// OrderBuilder 订单构造接口
type OrderBuilder interface {
	Create()
	GetOrder() *Order
}

// OrderDirector 订单建造者
type OrderDirector struct {
	build OrderBuilder
}

func (d *OrderDirector) Construct() {
	// 创建订单
	d.build.Create()
	// 计算订单总价和活动减免
	// 满5减1
	// 满10减3
	total := float32(0)
	order := d.build.GetOrder()
	for _, v := range order.Array {
		total += PriceMap[v.Name]
	}
	if total >= 10 {
		total -= 3
	} else if total >= 5 {
		total--
	}
	order.Total = total
}

func (d *OrderDirector) SetOrderBuilder(b OrderBuilder) {
	d.build = b
}

// 不同的客户订单

type Order1 struct {
	Order
}

func (o *Order1) Create() {
	o.Array = []Drinks{
		{Name: "橙汁"},
	}
}

func (o *Order1) GetOrder() *Order {
	return &o.Order
}

type Order2 struct {
	Order
}

func (o *Order2) Create() {
	o.Array = []Drinks{
		{Name: "橙汁"},
		{Name: "牛奶"},
	}
}

func (o *Order2) GetOrder() *Order {
	return &o.Order
}

type Order3 struct {
	Order
}

func (o *Order3) Create() {
	o.Array = []Drinks{
		{Name: "橙汁"},
		{Name: "牛奶"},
		{Name: "牛奶"},
	}
}

func (o *Order3) GetOrder() *Order {
	return &o.Order
}
