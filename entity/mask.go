package entity

const MaxComponents = 128

// ComponentMask struct is a simple bitset with room for MaxComponents.
// It's used to check, add and remove bits cheaply
type ComponentMask struct {
	bits [MaxComponents / 64]uint64
}

func (m *ComponentMask) Has(compId uint) bool {
	word, bit := compId/64, compId%64
	return (m.bits[word] & (1 << bit)) != 0
}

func (m *ComponentMask) Set(compId uint) {
	word, bit := compId/64, compId%64
	m.bits[word] |= (1 << bit)
}

func (m *ComponentMask) Clear(compId uint) {
	word, bit := compId/64, compId%64
	m.bits[word] &^= (1 << bit)
}

func (m *ComponentMask) Reset() {
	for i := range m.bits {
		m.bits[i] = 0
	}
}
