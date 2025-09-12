package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
}

type BinList struct {
	bins []Bin
}

func NewBinList(arr ...Bin) *BinList {
	var binArr []Bin
	binArr = append(binArr, arr...)

	return &BinList{
		bins: binArr,
	}
}
