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

func operation(sl []int) {
	var mg string
	for {
		fmt.Print("Введите операцию (AVG/SUM/MED): ")
		fmt.Scanln(&mg)
		switch mg {
		case "AVG":
			var res int
			var sum int
			for _, x := range sl {
				sum += x
			}
			res = sum / len(sl)
			fmt.Println("Среднее значение списка чисел: ", res)

		case "SUM":
			var sum int
			for _, x := range sl {
				sum += x
			}
			fmt.Println("Сумма чисел: ", sum)
		case "MED":
			var res int
			sort.Ints(sl)
			res = sl[len(sl)/2]
			fmt.Println("Медиана чисел: ", res)
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

func numbers() []int {
	var s string
	numSlice := make([]int, 0)
	fmt.Print("Введите числа через запятую: ")
	fmt.Scanln(&s)
	strSlice := strings.Split(s, ",")
	for _, str := range strSlice {
		conv, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		numSlice = append(numSlice, conv)
	}
	return numSlice
}
