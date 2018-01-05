package cuckoofilter

import (
	"github.com/spaolacci/murmur3"
	"math/rand"
	"time"
	"errors"
)
type CuckooFilter struct {
	Tables []*CuckooTable
	Size int
}
func NewCuckooFilter(capacity int, bucket_size int) *CuckooFilter {
	tables := make([]*CuckooTable, capacity)
	for i:=0; i < cap(tables); i++ {
		tables[i] = NewCuckooTable(bucket_size)
	}
	return &CuckooFilter{tables,0}	
}
func (cf *CuckooFilter) Insert(item string) (int,error) {
	// insert the item:string 
	// and return its index
	idx1, idx2 := cf.Indices(item)
	fingerprint := cf.Fingerprint(item)
	// fmt.Println(idx1,idx2, fingerprint)
  if cf.Tables[idx1].Insert(fingerprint) {
		cf.Size++
		return idx1,nil
	}
	if cf.Tables[idx2].Insert(fingerprint) {
		cf.Size++
		return idx2,nil
	}
	// if both indices are full, we need to swap all current entries.
	// First of all, randomly pick up the index between idx1 and idx2.
	// Afterwards, swap that item in its bucket
	var random int
	if idx2 > idx1 {
		rand.Seed(time.Now().UnixNano())
		random = rand.Intn(idx2-idx1)
		random += idx1
	} else if idx1 > idx2{
		random = rand.Intn(idx1-idx2)
		random += idx2
	} else {
		random = idx1
	}
	return cf.Rehash(random, fingerprint)
}

func (cf *CuckooFilter) Rehash(idx int, placement uint16) (int,error) {
	
	for i:=0 ;i < cap(cf.Tables)/2; i++ {
		placement = cf.Tables[idx].Swap(placement)
		idx = (idx ^ cf.Index(Fingerprint2Bytes(placement))) % cap(cf.Tables)

		if cf.Tables[idx].Insert(placement) {
			cf.Size++
			return idx,nil
		}
	}
	return -1, errors.New("Cuckoo Filter has filled up")
}

func (cf *CuckooFilter) Remove(item string) bool {
	fingerprint := cf.Fingerprint(item)
	idx1, idx2 := cf.Indices(item)

	if cf.Tables[idx1].Remove(fingerprint) {
		cf.Size--
		return true
	}

	if cf.Tables[idx2].Remove(fingerprint) {
		cf.Size--
		return true
	}

	// Both indices do not exist the item. It is not in the buckoo table
	return false
}

func (cf *CuckooFilter) Contains(item string) bool {
	fingerprint := cf.Fingerprint(item)
	idx1, idx2 := cf.Indices(item)

	if cf.Tables[idx1].Contains(fingerprint) || cf.Tables[idx2].Contains(fingerprint) {
		return true
	} else {
		return false
	}
}

func (cf *CuckooFilter) Fingerprint(item string) uint16 {
	// According fastforward labs, they adopted 2 bytes MSB for fingerprint 
	// reference: https://github.com/fastforwardlabs/cuckoofilter
	return uint16(murmur3.Sum64([]byte(item)) >> (64-16))
	// return murmur3.Sum64([]byte(str_item))
}

func (cf *CuckooFilter) Index(hash []byte) int {
	return int(murmur3.Sum64(hash) % uint64(cap(cf.Tables)))
}

func (cf *CuckooFilter) Indices(item string) (int, int) {
	// obtain the index first time
	idx1 := cf.Index([]byte(item))

	// obtain fingerprint
	fp := cf.Fingerprint(item)

	// derive the new index from fingerprint
	// second index assign as first idx xor the new one
	// , derived from Index(fingerprint)

	idx2 := idx1 ^ cf.Index(Fingerprint2Bytes(fp))
	return idx1, idx2 % cap(cf.Tables)
}
