package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("\033[32m---Калькулятор чисел---\033[0m")
	for {
		var message string
		numSl := numbers()
		fmt.Println("Слайс чисел: ", numSl)
		operation(numSl)
		fmt.Print("Ввести новые числа? (да/нет): ")
		fmt.Scanln(&message)
		if message == "да" {
			continue
		}
		fmt.Println("\033[31mЗавершение программы.\033[0m")
		break
	}
}

func operation(sl []float64) {
	var mg string
	for {
		fmt.Print("Введите операцию (AVG/SUM/MED): ")
		fmt.Scanln(&mg)
		switch mg {
		case "AVG":
			var res float64
			var sum float64
			for _, x := range sl {
				sum += x
			}
			res = sum / float64(len(sl))
			fmt.Printf("Среднее значение списка чисел: %.2f \n", res)

		case "SUM":
			var sum float64
			for _, x := range sl {
				sum += x
			}
			fmt.Println("Сумма чисел: ", sum)
		case "MED":
			var res float64
			sort.Float64s(sl)
			if len(sl)%2 == 0 {
				res = (sl[len(sl)/2] + sl[len(sl)/2-1]) / 2
			} else {
				res = sl[len(sl)/2]
			}
			fmt.Printf("Медиана чисел: %.2f \n", res)
		default:
			fmt.Println("\033[31mОперация введена неверно, попробуйте её раз.\033[0m")
			continue
		}
		fmt.Print("Произвести новую операцию? (да/нет): ")
		fmt.Scanln(&mg)
		if mg == "да" {
			continue
		}
		break
	}
}

func numbers() []float64 {
	var s string
	numSlice := make([]float64, 0)
	fmt.Print("Введите числа через запятую: ")
	fmt.Scanln(&s)
	strSlice := strings.Split(s, ",")
	for _, str := range strSlice {
		conv, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println(err)
		}
		numSlice = append(numSlice, conv)
	}
	return numSlice
}
