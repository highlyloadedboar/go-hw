package bins

import (
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	binList []Bin
}

func newBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

func newBinList() *BinList {
	return &BinList{
		binList: []Bin{},
	}
}

func (list *BinList) appendBin(bin *Bin) {
	list.binList = append(list.binList, *bin)
}
