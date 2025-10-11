package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var menu = map[string]func([]float64){
	"avg": Avg,
	"sum": Sum,
	"med": Med,
}

func main() {
	fmt.Println("\033[32m---Калькулятор чисел---\033[0m")
	for {
		var message string
		numSl := numbers()
		fmt.Println("Слайс чисел: ", numSl)
		operation(numSl)
		message = PromptData("Ввести новые числа? (да/нет): ")
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
		mg = PromptData("Введите операцию (avg/sum/med): ")
		menuFunc := menu[mg]
		menuFunc(sl)
		mg = PromptData("Произвести новую операцию? (да/нет): ")
		if mg == "да" {
			continue
		}
		break
	}
}

func Med(sl []float64) {
	var res float64
	sort.Float64s(sl)
	if len(sl)%2 == 0 {
		res = (sl[len(sl)/2] + sl[len(sl)/2-1]) / 2
	} else {
		res = sl[len(sl)/2]
	}
	fmt.Printf("Медиана чисел: %.2f \n", res)
}

func Avg(sl []float64) {
	var res float64
	var sum float64
	for _, x := range sl {
		sum += x
	}
	res = sum / float64(len(sl))
	fmt.Printf("Среднее значение списка чисел: %.2f \n", res)
}

func Sum(sl []float64) {
	var sum float64
	for _, x := range sl {
		sum += x
	}
	fmt.Println("Сумма чисел: ", sum)
}

func numbers() []float64 {
	var s string
	numSlice := make([]float64, 0)
	s = PromptData("Введите числа через запятую: ")
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

func PromptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
