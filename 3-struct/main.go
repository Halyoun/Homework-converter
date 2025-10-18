package main

import (
	"app-3/api"
	"app-3/bins"
	"app-3/storage"
	"flag"
	"fmt"
)

func main() {
	operation := flag.String("operation", "list", "operations of bins")
	file := flag.String("file", "data.json", "file name")
	name := flag.String("name", "Untitled", "name of bin")
	id := flag.String("id", "", "id of bin")
	flag.Parse()
	var BinList = bins.BinList{}
	//fmt.Println(api.GetKey(config.NewConfig("Key.env")))
	fmt.Println("---Менеджер бинов---")
	switch *operation {
	case "create":
		BinList.ScanJson(*file)
		BinList.NewBin()
		storage.SaveJson(*file, BinList)
		api.CreateBin(*file, *name, &BinList)
	case "update":
		BinList.ScanJson("data.json")
		BinList.PrintJson()
		api.UpdateBin(*file, *id, &BinList)
	case "delete":
		api.DeleteBin(*id)
	case "get":
		BinList = api.GetBin(*id)
		for _, bin := range BinList.Bins {
			fmt.Println(bin)
		}
	case "list":
		BinList.ScanJson(*file)
		for _, bin := range BinList.Bins {
			fmt.Println(bin.Id, bin.Name)
		}
	}
}
