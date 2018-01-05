package cuckoofilter


import (
	"fmt"
	"strconv"
	"testing"
)

func TestCuckooFilter(t *testing.T) {
	filter := NewCuckooFilter(1000,4)
	// fmt.Println(filter)
	if 0 != filter.Size {
		t.Error("amount is wrong", filter.Size)
	}

	if 1000 != cap(filter.Tables) {
		t.Error("capacity is wrong", cap(filter.Tables))
	}

	filter.Insert("James")
	if 1 != filter.Size {
		t.Error("After being inserted, amount is wrong", filter.Size)
	}

	if !filter.Contains("James") {
		t.Error("cannot find")
	}

	filter.Insert("Anderson")
	if !filter.Contains("James") || !filter.Contains("Anderson") {
		t.Error("lose one")
	}
	
	if filter.Size != 2 {
		t.Error("size wrong", filter.Size)
	}
	filter.Remove("Anderson")

	if filter.Contains("Anderson") {
		t.Error("fail to remove")
	}
}

func TestRehash(t *testing.T) {
	f := NewCuckooFilter(100,4)

	for i:=0; i < 10*100; i++ {
		_,ok := f.Insert(strconv.Itoa(i))
		// fmt.Printf("\r%d->%d",v,i)
		if ok != nil {
			fmt.Printf("fill %d\n",f.Size)
			t.Error(ok)
			return
		}
	}
}