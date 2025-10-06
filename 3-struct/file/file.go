package file

import (
	"github.com/fatih/color"
	"os"
)

func ReadJson(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		color.Red("Файл не найден! \n")
		return nil, err
	}
	return data, nil
}
