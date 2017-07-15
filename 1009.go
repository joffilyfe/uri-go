package main

import "fmt"

func main() {
	var name string
	var salary float64
	var sells float64

	fmt.Scan(&name, &salary, &sells)
	fmt.Printf("TOTAL = R$ %0.2f\n", salary+(sells*0.150))
}
