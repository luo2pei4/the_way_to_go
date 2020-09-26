package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (sp stockPosition) getValue() float32 {
	return sp.sharePrice * sp.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuabe interface {
	getValue() float32
}

func showValue(intf valuabe) {
	fmt.Printf("Value of the asset is %f\n", intf.getValue())
}

func main() {

	var o valuabe = stockPosition{"GOOD", 577.20, 4}
	showValue(o)

	o = car{"BMW", "M3", 66500}
	showValue(o)
}
