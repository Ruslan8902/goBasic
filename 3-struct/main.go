package main

import (
	"fmt"
	"gobasics/api"
	"gobasics/bins"
	"gobasics/config"
	"gobasics/storage"
)

func main() {
	configStuct := config.NewConfig()
	api.APIfunction(configStuct)
	db := storage.NewStorage("bins.json")
	binList := bins.NewBinListWithDb(db)
	fmt.Print(binList)
}
