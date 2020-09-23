package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) renew() {
	r.width = 10
	r.height = 10
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	r.renew()
	fmt.Println("area: ", r.width)
	fmt.Println("perim:", r.height)
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
}
