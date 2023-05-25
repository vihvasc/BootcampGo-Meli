package main

import "fmt"

func main() {
	var temperatura float64 = 25.6
	var umidade float64 = 62.8
	var pressao float64 = 1012.3

	fmt.Println("Temperatura:", temperatura, "°C")
	fmt.Println("Umidade:", umidade, "%")
	fmt.Println("Pressão atmosférica:", pressao, "hPa")
}