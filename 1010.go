package main

import (
	"fmt"
)

func main() {

	var (
		qtd   int
		valor float32
		total float32
		x     int
	)

	for i := 0; i < 2; i++ {
		fmt.Scanln(&x, &qtd, &valor)
		total += (valor * float32(qtd))
	}

	fmt.Printf("VALOR A PAGAR: R$ %0.2f\n", total)
}
