package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	binList []Bin
}

func (bin *Bin) newBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
}

func (list *BinList) newBinList() *BinList {
	return &BinList{
		binList: []Bin{},
	}
}

func (list *BinList) appendBin(bin *Bin) {
	list.binList = append(list.binList, *bin)
}

func main() {
	fmt.Println()
}
