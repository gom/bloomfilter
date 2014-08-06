package bloomfilter

import (
	"fmt"
	"hash"
	"hash/fnv"
	"math"
	"strconv"
)

type BloomFilter2 struct {
	m      uint
	k      uint
	bits   *BitArray
	hasher hash.Hash
}

func New2(m, k uint) *BloomFilter2 {
	return &BloomFilter2{m, k, NewB(m), fnv.New64()}
}

func NewWithEstimate2(n uint, p float64) *BloomFilter2 {
	m, k := estimate2(n, p)
	return New2(m, k)
}

func estimate2(n uint, p float64) (uint, uint) {
	m := -1 * (float64(n) * math.Log(p) / math.Pow(math.Log(2), 2))
	k := math.Ceil((float64(m) / float64(n)) * math.Log(2))
	return uint(m), uint(k)
}

func (bf *BloomFilter2) Add(str string) *BloomFilter2 {
	for _, h := range bf.hashes(str) {
		bf.bits.Set(h)
	}
	return bf
}

func (bf *BloomFilter2) Check(str string) bool {
	for _, h := range bf.hashes(str) {
		if !bf.bits.Has(h) {
			return false
		}
	}
	return true
}

func (bf *BloomFilter2) CheckAndAdd(str string) bool {
	result := true
	for _, h := range bf.hashes(str) {
		if !bf.bits.Has(h) {
			result = false
			bf.bits.Set(h)
		}
	}
	return result
}

func (bf *BloomFilter2) Clear() *BloomFilter2 {
	bf.bits = NewB(bf.m)
	bf.hasher.Reset()
	return bf
}

/**
 * http://en.wikipedia.org/wiki/Bloom_filter
 */
func (bf *BloomFilter2) FalsePositiveRate(n uint) float64 {
	return math.Pow(1-math.Exp((-1*float64(bf.k*n))/float64(bf.m)), float64(bf.k))
}

func (bf *BloomFilter2) hashes(str string) []uint {
	res := make([]uint, bf.k)
	for i := uint(0); i < bf.k; i++ {
		res[i] = bf.hash(str + string(i))
	}
	return res
}

func (bf *BloomFilter2) hash(str string) uint {
	bf.hasher.Reset()
	bf.hasher.Write([]byte(str))
	sum := fmt.Sprintf("%x", bf.hasher.Sum(nil))

	sub_sum, _ := strconv.ParseUint(sum, 16, 64)
	return uint(sub_sum % uint64(bf.m))
}
