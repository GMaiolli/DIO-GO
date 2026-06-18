package main

import "fmt"

func main() {
	var op int
	var temp float64
	nTemp := 0.0
	uni := ""
	fmt.Println("=== CONVERSOR DE ESCALAS TERMOMÉTRICAS ===")
	fmt.Println("1 - Celsius para Fahrenheit")
	fmt.Println("2 - Celsius para Kelvin")
	fmt.Println("3 - Fahrenheit para Celsius")
	fmt.Println("4 - Fahrenheit para Kelvin")
	fmt.Println("5 - Kelvin para Celsius")
	fmt.Println("6 - Kelvin para Fahrenheit")
	fmt.Print("Escolha uma opção: ")
	fmt.Scanln(&op)
	fmt.Print("Digite a temperatura: ")
	fmt.Scanln(&temp)
	switch op {
	case 1:
		uni = "Fahrenheit"
		nTemp = temp * (9.0/2.0) + 32.0
	case 2:
		uni = "Kelvin"
		nTemp = temp + 273.15
	case 3:
		uni = "Celsius"
		nTemp = (temp - 32.0) * (5.0/9.0)
	case 4:
		uni = "Kelvin"
		nTemp = (temp - 32.0) * (5.0/9.0) + 273.15
	case 5:
		uni = "Celsius"
		nTemp = temp - 273.15
	case 6:
		uni = "Fahrenheit"
		nTemp = (temp - 273.15) * (9.0/5.0) + 32.0
	default:
		fmt.Println("Opção inválida! Execute o programa novamente e escolha de 1 a 6.")
		return
	}
	fmt.Printf("Sua nova temperatura convertida para %s é de %.2f \n", uni, nTemp)
}