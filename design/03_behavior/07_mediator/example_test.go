package mediator

func Example() {
	m := new(Mediator)

	t := new(Technical)
	k := new(Market)

	m.Technical = t
	m.Market = k

	technical := &Technical{mediator: m}
	technical.Send("technical")
	market := Market{mediator: m}
	market.Send("market")
	// Output:
	// 技术部收到消息: technical
	// 市场部部收到消息: market
}
