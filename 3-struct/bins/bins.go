package bins

import (
	"encoding/json"
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
	Bins []Bin `json:"bins"`
}

type BinListWithDb struct {
	BinList
	db Db
}

func NewBinListWithDb(db Db) *BinListWithDb {
	binArrBytes, _ := db.ReadStorage()

	var binArr BinList
	err := json.Unmarshal(binArrBytes, &binArr)
	if err != nil {
		binArr.Bins = []Bin{}
	}

	return &BinListWithDb{
		BinList: BinList{
			Bins: binArr.Bins,
		},
		db: db,
	}
}
