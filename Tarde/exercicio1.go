package main

import "fmt"

func main() {
	palavra := "O Senhor dos Aneis"

	fmt.Println("A palavra tem", len(palavra), "letras")

	fmt.Println("Letras da palavra:")
	for _, letra := range palavra {
		fmt.Println(string(letra))
	}

}
