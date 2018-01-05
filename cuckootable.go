package cuckoofilter
import (
	"time"
	"math/rand"
)
type CuckooTable struct {
		Buckets []uint16
		Size int
}

func NewCuckooTable(size int) *CuckooTable {
	buckets := make([]uint16, size)
	return &CuckooTable{buckets,0}
}
func (ct *CuckooTable) Insert(fingerprint uint16) bool{
	// insert a fingerprint, 
	// if current bucket is not full
	if ct.Size < cap(ct.Buckets) {
		ct.Buckets[ct.Size] = fingerprint
		ct.Size++
		return true
	} else {
		// bucket is full. cuckoofilter copes with the a failed insert. 
		return false
	}
}

func (ct *CuckooTable) Remove(fingerprint uint16) bool {
	for i,v := range ct.Buckets {
		if v == fingerprint {
			ct.Buckets = append(ct.Buckets[:i],ct.Buckets[i+1:]...)
			ct.Size--
			return true
		}
	}
	return false
}

//
// while occuring collision, swapping buckets behaves as cuckoo birds.
//
func (ct *CuckooTable) Swap(fingerprint uint16) uint16 {
	rand.Seed(time.Now().UnixNano())  
	idx := rand.Intn(4)
	selected := ct.Buckets[idx]
	ct.Buckets[idx] = fingerprint
	return selected
}

func (ct *CuckooTable) Contains(fingerprint uint16) bool {
	for _,v := range ct.Buckets {
		if v == fingerprint {
			return true
		}
	}
	return false
}