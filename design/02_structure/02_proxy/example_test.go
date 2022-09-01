package proxy

func ExampleRouteProxy_Handler() {
	r := new(route)
	r.Handler("info")

	p := &RouteProxy{
		r,
	}
	p.Handler("info")
	// Output:
	// info
	// user_info
}
