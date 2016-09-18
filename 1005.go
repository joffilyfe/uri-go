package main

import "fmt"

func main() {
	var a, b float64

	fmt.Scan(&a, &b)
	a = a * 3.5
	b = b * 7.5
	fmt.Printf("MEDIA = %0.5f\n", (a+b)/11)
}
