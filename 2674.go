package main

import (
	"fmt"
	"strconv"
)

func isPrime(number int64) bool {
	if number <= 1 {
		return false
	} else if number <= 3 {
		return true
	} else if number%2 == 0 || number%3 == 0 {
		return false
	}

	i := 5

	for int64(i*i) <= number {
		if number%int64(i) == 0 || number%int64(i+2) == 0 {
			return false
		}

		i += 6
	}

	return true
}

func classify(number int64) string {

	if !isPrime(number) {
		return "Nada"
	}

	for _, char := range strconv.FormatInt(number, 10) {

		n, _ := strconv.Atoi(string(char))

		if !isPrime(int64(n)) {
			return "Primo"
		}
	}

	return "Super"
}

func main() {

	var number int64

	for {
		n, err := fmt.Scanf("%d\n", &number)

		if n == 0 || err != nil {
			break
		}

		fmt.Println(classify(int64(number)))

	}
}
