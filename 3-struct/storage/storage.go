package storage

import (
	"app-3/bins"
	"encoding/json"
	"github.com/fatih/color"
	"os"
)

func SaveJson(name string, bins bins.BinList) {
	file, err := os.Create(name)
	if err != nil {
		color.Red("Ошибка создания/открытия файла!\n")
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
