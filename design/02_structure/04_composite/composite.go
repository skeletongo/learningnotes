package composite

import "fmt"

// 组合模式

type ItemDesc struct {
	name        string
	description string
}

func (d *ItemDesc) Name() string {
	return d.name
}

func (d *ItemDesc) Description() string {
	return d.description
}

type Item struct {
	ItemDesc
	price float32
}

func NewItem(name, desc string, price float32) *Item {
	return &Item{
		ItemDesc: ItemDesc{
			name:        name,
			description: desc,
		},
		price: price,
	}
}

func (i *Item) Price() float32 {
	return i.price
}

func (i *Item) Print() {
	fmt.Printf("%s, ¥%0.2f, %s\n", i.name, i.price, i.description)
}

type ItemGroup struct {
	ItemDesc
	items []*Item
}

func NewItemGroup(name, desc string, item ...*Item) *ItemGroup {
	ret := new(ItemGroup)
	ret.name = name
	ret.description = desc
	for _, v := range item {
		ret.items = append(ret.items, v)
	}
	return ret
}

func (i *ItemGroup) Add(item *Item) {
	i.items = append(i.items, item)
}

func (i *ItemGroup) Remove(item *Item) {
	for k, v := range i.items {
		if v == item {
			i.items = append(i.items[:k], i.items[k+1:]...)
			break
		}
	}
}

func (i *ItemGroup) RemoveByName(name string) {
	for k, v := range i.items {
		if v.name == name {
			i.items = append(i.items[:k], i.items[k+1:]...)
			break
		}
	}
}

func (i *ItemGroup) Price() float32 {
	var n float32
	for _, v := range i.items {
		n += v.price
	}
	return n
}

func (i *ItemGroup) Print() {
	fmt.Printf("%s, %s, ¥%.2f\n", i.name, i.description, i.Price())
	fmt.Println("------------------------")
	for _, v := range i.items {
		v.Print()
	}
}
