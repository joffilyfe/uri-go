package main

import "fmt"

func main() {
	var (
		funcNumber int
		horas      float64
		valorHora  float64
	)

	fmt.Scan(&funcNumber, &horas, &valorHora)
	fmt.Println("NUMBER =", funcNumber)
	fmt.Printf("SALARY = U$ %0.2f\n", horas*valorHora)
}
