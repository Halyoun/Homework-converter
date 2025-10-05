package bins

import (
	"fmt"
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
	WriteJSON()
}

func PromptData(s string, p any) {
	fmt.Print(s)
	fmt.Scanln(p)
}
