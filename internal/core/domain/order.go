package domain

import (
	"fmt"
	"math/rand"
)

type Order struct {
	Id       int
	Products []Product
}

func (order Order) Init() Order {
	order.Id = rand.Intn(3000)
	return order
}

func (order *Order) AddProduct(product Product) {
	order.Products = append(order.Products, product)
}

func (order Order) GetTotalOrder() float64 {
	var total float64
	for i := 0; i < len(order.Products); i++ {
		total += order.Products[i].Price
	}
	return total
}

func (order Order) DetailOrder() {
	fmt.Printf("Número do Pedido: %v\n", order.Id)
	fmt.Println("Produtos:")
	for i := 0; i < len(order.Products); i++ {
		fmt.Printf("\t Nome: %v, Preço: R$ %v\n", order.Products[i].Name, order.Products[i].Price)
	}
	fmt.Printf("\nTotal do Pedido: R$ %v\n\n", order.GetTotalOrder())
}
