package file

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func ReadJson(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		color.Red("Файл не найден! \n")
		return nil, err
	}
	if strings.Contains(name, ".json") {
		fmt.Println("\033[31mФайл не является json!\033[0m")
		return nil, nil
	}
	return data, nil
}
