package main

import (
	"fmt"
	"structs/producto"
)

func main() {
	spec := producto.Spec{
		Maker: "manzana",
		Price: 10,
	}
	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
}
