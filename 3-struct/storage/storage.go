package storage

import (
	"app-3/bins"
	"app-3/file"
	"encoding/json"
	"github.com/fatih/color"
	"os"
)

func SaveJson(name string, bins bins.BinList) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		color.Red("Ошибка файла!\n")
		return
	}
	defer file.Close()

	data, err := json.Marshal(bins)
	if err != nil {
		color.Red("Ошибка преобразования в json!\n")
		return
	}

	_, err = file.Write(data)
	if err != nil {
		color.Red("Не удалось записать данные в json файл!\n")
		return
	}

	color.Green("Данные записаны успешно!\n")
}

func ScanJson(name string) *bins.BinList {
	return file.ReadJson(name)
}
