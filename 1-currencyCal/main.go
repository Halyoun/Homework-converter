package main

import (
	"errors"
	"fmt"
)

const usdToEur = 0.8514
const usdToRub = 83.59
const eurToRub = usdToRub / usdToEur

func main() {
	fmt.Println("\033[32m---Добро пожаловать в конвертор валют!---\033[0m")
	for {
		var mg string
		res, err := convertOper(userInput())
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		fmt.Printf("\033[32mРезультат: %.2f \033[0m\n", res)
		fmt.Print("Начать заново? (да/нет): ")
		fmt.Scanln(&mg)
		if mg == "да" {
			continue
		}
		fmt.Println("\033[31mВыход из программы\033[0m")
		break
	}
}

func userInput() (float64, string, string) {
	var sum float64
	var curr1 string
	var curr2 string
	for {
		fmt.Print("Введите валюту, из которой будет конвертация (usd/eur/rub): ")
		fmt.Scanln(&curr1)
		if curr1 != "usd" && curr1 != "eur" && curr1 != "rub" {
			fmt.Println("Ошибка ввода")
			continue
		}
		break
	}
	for {
		fmt.Print("Введите сумму конвертации: ")
		fmt.Scanln(&sum)
		if sum <= 0 {
			fmt.Println("Введено число меньше нуля или ноль")
			continue
		} else {
			break
		}
	}
	for {
		switch curr1 {
		case "usd":
			fmt.Print("Введите валюту, в которую хотите сконвертировать (eur/rub): ")
		case "eur":
			fmt.Print("Введите валюту, в которую хотите сконвертировать (usd/rub): ")
		case "rub":
			fmt.Print("Введите валюту, в которую хотите сконвертировать (usd/eur): ")
		default:
			continue
		}
		fmt.Scanln(&curr2)
		if curr2 != "usd" && curr2 != "eur" && curr2 != "rub" || curr2 == curr1 {
			fmt.Println("Ошибка ввода")
			continue
		}
		break
	}
	return sum, curr1, curr2
}

func convertOper(sum float64, curr1 string, curr2 string) (float64, error) {
	switch {
	case curr1 == "usd" && curr2 == "eur":
		return sum * usdToEur, nil
	case curr1 == "usd" && curr2 == "rub":
		return sum * usdToRub, nil
	case curr1 == "eur" && curr2 == "rub":
		return sum * eurToRub, nil
	case curr1 == "eur" && curr2 == "usd":
		return sum * (1 / usdToEur), nil
	case curr1 == "rub" && curr2 == "usd":
		return sum * (1 / usdToRub), nil
	case curr1 == "rub" && curr2 == "eur":
		return sum * (1 / eurToRub), nil
	default:
		return 0.0, errors.New("данные введены неверно")
	}
}
