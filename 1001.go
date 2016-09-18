package main

import (
	"fmt"
	"os"
)

func main() {
	var n, m int

	_, err := fmt.Scanf("%d\n%d", &n, &m)

	if err != nil {
		fmt.Println("Please write just integer numbers")
		os.Exit(1)
	}

	fmt.Printf("X = %d\n", n+m)
}
