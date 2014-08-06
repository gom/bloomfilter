package bloomfilter

type BitArray struct {
	boxes  []uint64
	length uint
}

const (
	bitsPerBox uint   = 64
	baseBit    uint64 = 1
)

func NewB(length uint) *BitArray {
	return &BitArray{make([]uint64, box_pos(length)+1), length}
}

func (b *BitArray) Len() uint {
	return b.length
}

func (b *BitArray) Set(i uint) {
	if i > b.length {
		return
	}
	b.boxes[box_pos(i)] |= (baseBit << bit_pos(i))

}

func (b *BitArray) Delete(i uint) {
	if i > b.length {
		return
	}
	b.boxes[box_pos(i)] ^= (baseBit << bit_pos(i))
}

func (b *BitArray) Has(i uint) bool {
	if i > b.length {
		return false
	}
	return (b.boxes[box_pos(i)] & (baseBit << bit_pos(i))) != 0
}

func box_pos(hint uint) uint {
	return hint >> 6
}

func bit_pos(i uint) uint {
	return i & (bitsPerBox - 1)
}
