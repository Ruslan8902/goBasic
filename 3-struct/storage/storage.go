package storage

import (
	"encoding/json"
	"fmt"
	"gobasics/bins"
	"gobasics/file"
	"os"
)

func WriteStorage(content []byte, storage string) {
	file, err := os.Create(storage)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись прошла успешно!")
}

func ReadStorage(storage string) bins.BinList {
	if file.IsJson(storage) {
		file, err := os.ReadFile(storage)
		if err != nil {
			fmt.Println(err)
			return bins.BinList{}
		}

		var arr bins.BinList
		err = json.Unmarshal(file, &arr)
		if err != nil {
			fmt.Println(err)
			return bins.BinList{}
		}

		return arr
	}
	return bins.BinList{}
}

func SaveBinListJSON(binList *bins.BinList) {
	binsBytes, err := ToBytes(binList)
	if err != nil {
		fmt.Println(err)
	}
	WriteStorage(binsBytes, "bins.json")
}

func ToBytes(binList *bins.BinList) ([]byte, error) {
	file, err := json.Marshal(binList)
	if err != nil {
		return nil, err
	}
	return file, nil
}
