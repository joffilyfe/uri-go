package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(input string) string {
	novaString := ""

	for i := len(input) - 1; i >= 0; i-- {
		novaString += string(input[i])
	}

	return novaString
}

func cripto(input string) string {
	novaString := ""
	for _, char := range input {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			novaString += string(char + 3)
		} else {
			novaString += string(char)
		}

	}

	novaString = reverse(novaString)

	metade_inferior := novaString[:len(novaString)/2]
	metade_superior := novaString[len(novaString)/2:]

	for _, char := range metade_superior {
		metade_inferior += string(char - 1)
	}

	return metade_inferior
}

func main() {
	var qtd int

	fmt.Scan(&qtd)

	scanner := bufio.NewReader(os.Stdin)

	for i := 0; i < qtd; i++ {
		r, _, _ := scanner.ReadLine()
		fmt.Println(cripto(string(r)))
	}
}
