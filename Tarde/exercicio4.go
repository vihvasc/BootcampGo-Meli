package main

import "fmt"

func main() {
	employees := map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	idadeBenjamin, ok := employees["Benjamin"]
	if ok {
		fmt.Println("Idade de Benjamin:", idadeBenjamin)
	} else {
		fmt.Println("Benjamin não encontrado")
	}

	count := 0
	for _, idade := range employees {
		if idade > 21 {
			count++
		}
	}
	fmt.Println("Número de funcionários com mais de 21 anos:", count)

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println("Funcionários:")
	for nome, idade := range employees {
		fmt.Println(nome, "-", idade, "anos")
	}
}
