package main

import (
	"app-3/api"
	"app-3/bins"
	"app-3/config"
	"app-3/storage"
	"fmt"
)

func main() {
	BinList := bins.BinList{}
	var mg string
	fmt.Println(api.GetKey(config.NewConfig("Key.env")))
	fmt.Println("---Менеджер бинов---")
Menu:
	for {
		bins.PromptData("Введите команду: ", &mg)
		switch mg {
		case "1":
			BinList.ScanJson("data.json")
			BinList.NewBin()
			storage.SaveJson("data.json", BinList)
		case "2":
			BinList.ScanJson("data.json")
			BinList.PrintJson()
		default:
			break Menu
		}
	}
}
