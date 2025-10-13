package bins

import (
	"app-3/file"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins      []Bin     `json:"bins"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (list *BinList) NewBin() {
	bin := Bin{}
	PromptData("Введите id: ", &bin.Id)
	PromptData("Задайте приватность (true/false): ", &bin.Private)
	PromptData("Введите имя: ", &bin.Name)
	bin.CreatedAt = time.Now()
	list.Bins = append(list.Bins, bin)
	list.UpdatedAt = time.Now()
}

func (list *BinList) ScanJson(name string) {
	data, err := file.Read(name)
	err = json.Unmarshal(data, &list)
	if err != nil {
		color.Red("Файл не является json или не удалось преобразовать!")
		return
	}
}

func (list *BinList) PrintJson() {
	fmt.Println("\033[31m---- Список бинов ----\033[m")
	for _, bin := range list.Bins {
		fmt.Println("\033[32m---------------------------\033[m")
		fmt.Println("ID: ", bin.Id)
		fmt.Println("Name: ", bin.Name)
		fmt.Println("Created at: ", bin.CreatedAt)
		fmt.Println("Private: ", bin.Private)

	}
	fmt.Println("\033[32m---------------------------\033[m")
	fmt.Println("\033[31mПоследнее обновление: \033[m", list.UpdatedAt)
}

func PromptData(s string, p any) {
	fmt.Print(s)
	fmt.Scanln(p)
}
