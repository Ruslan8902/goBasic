package bins

import "time"

type Db interface {
	WriteStorage(content []byte, storage string)
	ReadStorage(storage string) BinList
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

func NewBinList(db Db, storage string, arr ...Bin) *BinListWithDb {
	binArr := db.ReadStorage(storage)
	binArr.Bins = append(binArr.Bins, arr...)

	return &BinListWithDb{
		BinList: BinList{
			Bins: binArr.Bins,
		},
		db: db,
	}
}
