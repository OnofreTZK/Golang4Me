package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// Method with receiver for the custom type
func (prod product) priceWithDiscount() float64 {
	return prod.price * (1 - prod.discount)
}

func main() {
	var product1 product
	product1 = product{
		name:     "iPad Pro 11' Space Grey 2021",
		price:    7200.00,
		discount: 0.15,
	}

	fmt.Println(product1, product1.priceWithDiscount())

	product2 := product{"PlayStation 5 Pink", 5000.0, 0.1}
	fmt.Println(product2, product2.priceWithDiscount())
}
