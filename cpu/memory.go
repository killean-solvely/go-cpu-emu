package cpu

const (
	StoredMemorySize = 2048 // 2kb
	TotalMemorySize  = 65536
	CodeMemoryStart  = StoredMemorySize
)

type Memory struct {
	Data [TotalMemorySize]uint8
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) ReadByte(address uint32) uint8 {
	if address >= TotalMemorySize {
		panic("Memory read out of bounds")
	}
	return m.Data[address]
}

func (m *Memory) ReadWord(address uint32) uint16 {
	if address >= TotalMemorySize {
		panic("Memory read out of bounds")
	}
	return bytesToUint16(m.Data[address : address+2])
}

func (m *Memory) WriteByte(address uint32, value uint8) {
	if address >= TotalMemorySize {
		panic("Memory write out of bounds")
	}
	m.Data[address] = value
}

func (m *Memory) WriteWord(address uint32, value uint16) {
	if address >= TotalMemorySize {
		panic("Memory write out of bounds")
	}
	bytes := uint16ToBytes(value)
	m.Data[address] = bytes[0]
	m.Data[address+1] = bytes[1]
}

func (m *Memory) ReadStoredMemoryByte(address uint32) uint8 {
	if address >= StoredMemorySize {
		panic("Stored memory read out of bounds")
	}
	return m.Data[address]
}

func (m *Memory) ReadStoredMemoryWord(address uint32) uint16 {
	if address >= StoredMemorySize {
		panic("Stored memory read out of bounds")
	}
	return bytesToUint16(m.Data[address : address+2])
}

func (m *Memory) WriteStoredMemoryByte(address uint32, value uint8) {
	if address >= StoredMemorySize {
		panic("Stored memory write out of bounds")
	}
	m.Data[address] = value
}

func (m *Memory) WriteStoredMemoryWord(address uint32, value uint16) {
	if address >= StoredMemorySize {
		panic("Stored memory write out of bounds")
	}
	bytes := uint16ToBytes(value)
	m.Data[address] = bytes[0]
	m.Data[address+1] = bytes[1]
}

func (m *Memory) LoadCode(code []uint8) {
	if len(code) > (TotalMemorySize - CodeMemoryStart) {
		panic("Code exceeds available memory space")
	}
	for i, b := range code {
		m.WriteByte(uint32(CodeMemoryStart+i), b)
	}
}
