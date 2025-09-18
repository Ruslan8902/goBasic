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

	db := storage.NewStorage("bins.json")
	binListWithDb := bins.NewBinListWithDb(db)

	if *createFlag {
		fmt.Print("create", *fileFlag, *nameFlag)
		err := api.CreateBin(configStuct, binListWithDb, fileFlag, nameFlag)
		if err != nil {
			fmt.Print(err)
		}
	}
	if *updateFlag {
		fmt.Println("update", *fileFlag, *idFlag)
		respBody, err := api.UpdateBin(configStuct, binListWithDb, idFlag, fileFlag)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(respBody)
	}
	if *deleteFlag {
		respBody, err := api.DeleteBin(configStuct, binListWithDb, idFlag)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(respBody)
	}
	if *getFlag {
		fmt.Println("get", *idFlag)
		respBody, err := api.GetBin(configStuct, idFlag)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(respBody)
	}

	if *listFlag {
		fmt.Println("list")
		binListWithDb.PrintBinList()
	}
}
