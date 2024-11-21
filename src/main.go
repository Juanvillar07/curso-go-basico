package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("51")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

}
