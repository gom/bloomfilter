package bloomfilter

import (
	"math/rand"
	"testing"
)

func TestAddAndCheck(t *testing.T) {
	bf := New(1000, 3)
	s1 := "Anderson"
	s2 := "Bred"
	s3 := "Chad"
	s4 := "Dave"

	bf.Add(s1).Add(s2).CheckAndAdd(s3)
	for _, s := range []string{s1, s2, s3} {
		if !bf.Check(s) {
			t.Errorf("%s shourd be in.", s)
		}
	}
	if bf.Check(s4) {
		t.Errorf("%s shourd not be in.", s4)
	}
}

func TestFalsePositiveRate(t *testing.T) {
	bf := New(10000, 3)
	n := uint(100)
	rate := bf.FalsePositiveRate(n)
	if rate > 0.0001 {
		t.Errorf("Too high false positive rate: %f", rate)
	}
}

func BenchmarkCheckAndAddSeparated(b *testing.B) {
	bf := New(10000, 4)
	s := "foobar"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ts := s + string(rand.Intn(10000000))
			bf.Check(ts)
			bf.Add(ts)
		}
	})
}

func BenchmarkCheckAndAddCombined(b *testing.B) {
	bf := New(10000, 4)
	s := "foobar"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ts := s + string(rand.Intn(10000000))
			bf.CheckAndAdd(ts)
		}
	})
}
