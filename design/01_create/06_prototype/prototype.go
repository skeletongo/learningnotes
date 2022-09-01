package prototype

// 原型模式
// 复制已存在的实例

type Phone struct {
	CPU    string
	Memory string
	IMEI   string
}

func Clone(p *Phone, imei string) *Phone {
	ret := *p
	ret.IMEI = imei
	return &ret
}
