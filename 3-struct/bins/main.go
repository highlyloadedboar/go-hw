package bins

import (
	"3-sruct/files"
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

func newBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
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

func main() {
	bin := newBin(
		"1", true, time.Now().Local(), "name",
	)
	fmt.Println(bin.name)

	files.ReadFile()

}
