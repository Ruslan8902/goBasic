package bins

import (
	"encoding/json"
	"fmt"
	"time"
)

type Db interface {
	WriteStorage([]byte)
	ReadStorage() ([]byte, error)
}

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

type BinList struct {
	Bins []Bin
}

type BinListWithDb struct {
	BinList
	Db Db
}

func NewBinListWithDb(db Db) *BinListWithDb {
	binArrBytes, _ := db.ReadStorage()

	var binArr []Bin
	err := json.Unmarshal(binArrBytes, &binArr)
	if err != nil {
		return &BinListWithDb{
			BinList: BinList{
				Bins: []Bin{},
			},
			Db: db,
		}
	}

	return &BinListWithDb{
		BinList: BinList{
			Bins: binArr,
		},
		Db: db,
	}
}

func (b *BinListWithDb) PrintBinList() {
	for i, bin := range b.Bins {
		fmt.Printf("%v. %v: %v\n", i+1, bin.Id, bin.Name)
	}
}
