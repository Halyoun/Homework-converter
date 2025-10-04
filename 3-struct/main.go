package main

import (
	"app-3/bins"
	"fmt"
)

func main() {
	BinList := []bins.Bin{}
	bins.NewBin(&BinList)
	fmt.Println("BinList:", BinList)
}
