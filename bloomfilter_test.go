package bloomfilter

import "testing"

func TestAddAndCheck(t *testing.T) {
	bf := New(1000, 3)
	s1 := "Anderson"
	s2 := "Bred"
	s3 := "Chad"
	s4 := "Dave"

	bf.Add(s1).Add(s2).CheckAndAdd(s3)
	for _, s := range []string{s1, s2, s3} {
		if !bf.Check(s) {
			t.Errorf("%s should be in.", s)
		}
	}
	if bf.Check(s4) {
		t.Errorf("%s should not be in.", s4)
	}

	bf.Clear()
	if bf.Check(s1) {
		t.Errorf("%s should not be in.", s1)
	}
}

func TestClear(t *testing.T) {
	bf := New(1000, 3)
	s := "Anderson"

	bf.Add(s)
	if !bf.Check(s) {
		t.Errorf("%s should be in", s)
	}

	bf.Clear()
	if bf.Check(s) {
		t.Errorf("%s should not be in", s)
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

func TestEstimate(t *testing.T) {
	n := uint(65536)
	p := 0.005
	marging := 0.001

	bf := NewWithEstimate(n, p)
	rate := bf.FalsePositiveRate(n)
	if rate < (p-marging) || (p+marging) < rate {
		t.Errorf("Too small/big estimate: %f", rate)
	}
}

func BenchmarkCheckAndAddSeparated(b *testing.B) {
	bf := New(10000, 4)
	s := "foobar"

	for i := 0; i < b.N; i++ {
		ts := s + string(i)
		bf.Check(ts)
		bf.Add(ts)
	}
}

func BenchmarkCheckAndAddCombined(b *testing.B) {
	bf := New(10000, 4)
	s := "foobar"

	for i := 0; i < b.N; i++ {
		ts := s + string(i)
		bf.CheckAndAdd(ts)
	}
}
