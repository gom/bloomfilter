package bloomfilter

import "testing"

func TestBasic(t *testing.T) {
	var l uint = 63
	b := NewBitArray(l)
	if b.Len() != l {
		t.Errorf("Len should be ", l)
	}

	v := uint(37)
	if b.Has(v) {
		t.Errorf("%s should not be in", v)
	}
	b.Set(v)
	if !b.Has(v) {
		t.Errorf("%s should be in", v)
	}
	b.Delete(v)
	if b.Has(v) {
		t.Errorf("%s should not be in", v)
	}

	f := uint(17)
	if b.Has(f) {
		t.Errorf("%s should not be in", f)
	}
}

func TestBorderValues(t *testing.T) {
	test_data := []uint{0, 1, 2, 62, 63, 64, 65, 126, 127, 128, 129}
	b := NewBitArray(uint(130))

	for _, v := range test_data {
		if b.Has(v) {
			t.Errorf("%s should not be in", v)
		}
		b.Set(v)
		if !b.Has(v) {
			t.Errorf("%s should be in", v)
		}
		b.Delete(v)
		if b.Has(v) {
			t.Errorf("%s should not be in", v)
		}
	}
}

func TestInvalidValues(t *testing.T) {
	var l uint = 63
	b := NewBitArray(l)

	v := l + 1
	b.Set(v)
	if b.Has(v) {
		t.Errorf("%s should not be in", v)
	}

	b.Delete(v)
	if b.Has(v) {
		t.Errorf("%s should not be in", v)
	}
}
