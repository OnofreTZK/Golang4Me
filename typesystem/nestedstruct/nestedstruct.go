package main

import "fmt"

type item struct {
	productId int
	amount    int
	price     float64
}

type order struct {
	userId int
	itens  []item // slice
}

func (ord order) totalValue() float64 {
	total := 0.0
	for _, item := range ord.itens {
		total += item.price * float64(item.amount)
	}

	return total
}

func main() {

	ord := order{
		userId: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 23, 34.10},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("Total: %.2f\n", ord.totalValue())
}
