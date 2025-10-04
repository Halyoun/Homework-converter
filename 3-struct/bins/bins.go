package bins

import (
	"fmt"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBin(binList *[]Bin) {
	bin := Bin{}
	promptData("Введите id: ", &bin.Id)
	promptData("Задайте приватность (true/false): ", &bin.Private)
	promptData("Введите имя: ", &bin.Name)
	bin.CreatedAt = time.Now()
	*binList = append(*binList, bin)
}

func promptData(s string, p any) {
	fmt.Print(s)
	fmt.Scanln(p)
}
