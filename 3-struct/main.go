package main

import (
	"flag"
	"fmt"
	"gobasics/api"
	"gobasics/bins"
	"gobasics/config"
	"gobasics/storage"
)

func main() {
	configStuct := config.NewConfig()

	createFlag := flag.Bool("create", false, "Create function")
	updateFlag := flag.Bool("update", false, "Update function")
	deleteFlag := flag.Bool("delete", false, "Delete function")
	getFlag := flag.Bool("get", false, "Get function")
	listFlag := flag.Bool("list", false, "List function")

	fileFlag := flag.String("file", "", "File path")
	idFlag := flag.String("id", "", "Bin id")
	nameFlag := flag.String("name", "", "Bin name")

	flag.Parse()

	if *createFlag {
		fmt.Print("create", *fileFlag, *nameFlag)

	}
	if *updateFlag {
		fmt.Print("update", *fileFlag, *idFlag)
	}
	if *deleteFlag {
		fmt.Print("delete", *idFlag)
	}
	if *getFlag {
		fmt.Print("get", *idFlag)
	}
	if *listFlag {
		fmt.Print("list")
	}

	api.GetBin(configStuct)
	db := storage.NewStorage("bins.json")
	binList := bins.NewBinListWithDb(db)
	fmt.Print(binList)
}
