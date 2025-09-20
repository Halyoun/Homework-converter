package main

import "fmt"

const usdToEur = 0.8514
const usdToRub = 83.59
const eurToRub = usdToRub / usdToEur

func main() {
	fmt.Println("Eur to Rub", eurToRub)
}
