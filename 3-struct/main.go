package main

import (
	"app-3/bins"
	"app-3/storage"
	"fmt"
)

func main() {
	BinList := bins.BinList{}
	var mg string
	fmt.Println("---Менеджер бинов---")
Menu:
	for {
		bins.PromptData("Введите команду: ", &mg)
		switch mg {
		case "1":
			BinList.NewBin()
			storage.SaveJson("data.json", BinList)
		case "2":
			BinList = *storage.ScanJson("data.json")
			fmt.Println(BinList)
		default:
			break Menu
		}
	}
}
