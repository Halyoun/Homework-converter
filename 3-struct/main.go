package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func main() {
	BinList := []Bin{}
	newBin(&BinList)
	fmt.Println("BinList:", BinList)
}

func newBin(binList *[]Bin) {
	bin := Bin{}
	promptData("Введите id: ", &bin.id)
	promptData("Задайте приватность (true/false): ", &bin.private)
	promptData("Введите имя: ", &bin.name)
	bin.createdAt = time.Now()
	*binList = append(*binList, bin)
}

func promptData(s string, p any) {
	fmt.Print(s)
	fmt.Scanln(p)
}
