package file

import (
	"app-3/bins"
	"encoding/json"
	"github.com/fatih/color"
	"os"
)

func WriteJSON() {
	
}

func ReadJson(name string) *bins.BinList {
	bins := bins.BinList{}
	file, err := os.ReadFile(name)
	if err != nil {
		color.Red("Файл не найден! \n")
		return nil
	}
	err = json.Unmarshal(file, &bins)
	if err != nil {
		color.Red("Файл не является json!")
		return nil
	}
	return &bins
}
