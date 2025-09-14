package main

import (
	"fmt"
	"gobasics/bins"
	"gobasics/storage"
)

func main() {
	db := storage.NewStorage("bins.json")
	binList := bins.NewBinListWithDb(db)
	fmt.Print(binList)
}
