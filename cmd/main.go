package main

import "github.com/evertonvieira/estudos-golang/internal/core/domain"

func main() {
	order := new(domain.Order).Init()
	order2 := new(domain.Order).Init()

	product1 := domain.Product{
		Name:  "cerveja",
		Price: 14.4,
	}

	product2 := domain.Product{
		Name:  "água",
		Price: 5,
	}
	product3 := domain.Product{
		Name:  "suco de cajú",
		Price: 3.5,
	}
	order.AddProduct(product1)
	order.AddProduct(product2)
	order.AddProduct(product3)

	order.DetailOrder()
	order2.DetailOrder()

}
