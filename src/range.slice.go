package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func rangeSlice() {
	slice := []string{"Hola", "que", "hace"}

	for i := range slice {
		fmt.Println(i)
	}
}

func main() {
	rangeSlice()
	isPalindromo("Ama")
}
