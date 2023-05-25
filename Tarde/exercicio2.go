package main

import "fmt"

func main() {
	idade := 25
	empregado := true
	anosAtividade := 2
	salario := 150000

	if idade > 22 && empregado && anosAtividade > 1 {
		if salario > 100000 {
			fmt.Println("Parabéns! Você é elegível para um empréstimo sem juros.")
		} else {
			fmt.Println("Parabéns! Você é elegível para um empréstimo.")
		}
	} else {
		fmt.Println("Desculpe, você não é elegível para um empréstimo.")
	}
}
