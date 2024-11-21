package main

import "fmt"

func keywords() {
	// Defer
	defer fmt.Println("Hola") // Ejecutar ultima funcion
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

}

func main() {
	keywords()
}
