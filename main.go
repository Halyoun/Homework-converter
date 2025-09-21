package main

import "fmt"

const usdToEur = 0.8514
const usdToRub = 83.59
const eurToRub = usdToRub / usdToEur

func main() {
	userInput()
}

func userInput() (float64, string, string) {
	var sum float64
	var curr1 string
	var curr2 string
	fmt.Print("Введите валюту, из которой будет конвертация (usd/eur/rub): ")
	fmt.Scanln(&curr1)
	fmt.Print("Введите валюту, в которую хотите сконвертировать (usd/eur/rub): ")
	fmt.Scanln(&curr2)
	fmt.Print("Введите сумму: ")
	fmt.Scanln(&sum)
	return sum, curr1, curr2
}

func convertOper(sum float64, curr1 string, curr2 string) float64 {
	return 0.0
}
