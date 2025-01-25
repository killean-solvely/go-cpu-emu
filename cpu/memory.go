package cpu

const (
	StoredMemorySize = 55
	TotalMemorySize  = 256
	CodeMemoryStart  = StoredMemorySize
)

type Memory struct {
	Data [TotalMemorySize]uint8
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Read(address uint16) uint8 {
	if address >= TotalMemorySize {
		panic("Memory read out of bounds")
	}
	return m.Data[address]
}

func (m *Memory) Write(address uint16, value uint8) {
	if address >= TotalMemorySize {
		panic("Memory write out of bounds")
	}
	m.Data[address] = value
}

func (m *Memory) ReadStoredMemory(address uint16) uint8 {
	if address >= StoredMemorySize {
		panic("Stored memory read out of bounds")
	}
	return m.Data[address]
}

func (m *Memory) WriteStoredMemory(address uint16, value uint8) {
	if address >= StoredMemorySize {
		panic("Stored memory write out of bounds")
	}
	m.Data[address] = value
}

func (m *Memory) LoadCode(code []uint8) {
	if len(code) > (TotalMemorySize - CodeMemoryStart) {
		panic("Code exceeds available memory space")
	}
	for i, b := range code {
		m.Write(uint16(CodeMemoryStart+i), b)
	}
}
