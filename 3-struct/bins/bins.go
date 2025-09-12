package bins

import "time"

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

func NewBinList(arr ...Bin) *BinList {
	var binArr []Bin
	binArr = append(binArr, arr...)

	return &BinList{
		Bins: binArr,
	}
}
