package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	a = a * 2
	b = b * 3
	c = c * 5

	fmt.Printf("MEDIA = %0.1f\n", (a+b+c)/10)
}
