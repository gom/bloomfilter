package bloomfilter

import (
	"fmt"
	"hash"
	"hash/fnv"
	"math"
	"strconv"
)

type BloomFilter struct {
	m      uint
	k      uint
	bits   []bool
	hasher hash.Hash
}

func New(m, k uint) *BloomFilter {
	return &BloomFilter{m, k, make([]bool, m), fnv.New64()}
}

func NewWithEstimate(n uint, p float64) *BloomFilter {
	m, k := estimate(n, p)
	return New(m, k)
}

func estimate(n uint, p float64) (uint, uint) {
	m := -1 * (float64(n) * math.Log(p) / math.Pow(math.Log(2), 2))
	k := math.Ceil((float64(m) / float64(n)) * math.Log(2))
	return uint(m), uint(k)
}

func (bf *BloomFilter) Add(str string) *BloomFilter {
	for _, h := range bf.hashes(str) {
		bf.bits[h] = true
	}
	return bf
}

func (bf *BloomFilter) Check(str string) bool {
	for _, h := range bf.hashes(str) {
		if !bf.bits[h] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) CheckAndAdd(str string) bool {
	result := true
	for _, h := range bf.hashes(str) {
		if !bf.bits[h] {
			result = false
			bf.bits[h] = true
		}
	}
	return result
}

func (bf *BloomFilter) Clear() *BloomFilter {
	bf.bits = make([]bool, bf.m)
	bf.hasher.Reset()
	return bf
}

/**
 * http://en.wikipedia.org/wiki/Bloom_filter
 */
func (bf *BloomFilter) FalsePositiveRate(n uint) float64 {
	return math.Pow(1-math.Exp((-1*float64(bf.k*n))/float64(bf.m)), float64(bf.k))
}

func (bf *BloomFilter) hashes(str string) []uint {
	res := make([]uint, bf.k)
	for i := uint(0); i < bf.k; i++ {
		res[i] = bf.hash(str + string(i))
	}
	return res
}

func (bf *BloomFilter) hash(str string) uint {
	bf.hasher.Reset()
	bf.hasher.Write([]byte(str))
	sum := fmt.Sprintf("%x", bf.hasher.Sum(nil))

	sub_sum, _ := strconv.ParseUint(sum, 16, 64)
	return uint(sub_sum % uint64(bf.m))
}
