package main

import (
	"app-3/api"
	"app-3/bins"
	"app-3/config"
	"app-3/storage"
	"fmt"
)

func main() {
	//operation := flag.String("operation", "list", "operations of bins")
	//file := flag.String("file", "", "file name")
	//nameOfBin := flag.String("name", "Untitled", "name of bin")
	//flag.Parse()
	var mg string
	var BinList = bins.BinList{}
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
			api.CreateBin("data.json", &BinList)
		case "2":
			BinList.ScanJson("data.json")
			BinList.PrintJson()
		default:
			break Menu
		}
	}
}
