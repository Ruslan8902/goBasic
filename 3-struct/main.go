package main

import (
	"fmt"
	"gobasics/bins"
	"gobasics/config"
	"gobasics/storage"
)

func main() {
	configStuct := config.NewConfig()
	fmt.Println(configStuct)
	db := storage.NewStorage("bins.json")
	binList := bins.NewBinListWithDb(db)
	fmt.Print(binList)
}
