package cpu

import "testing"

func TestMemoryReadWriteByte(t *testing.T) {
	mem := NewMemory()

	mem.WriteByte(10, 42)
	value := mem.ReadByte(10)

	if value != 42 {
		t.Errorf("Expected memory at address 10 to be 42, got %d", value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for out of bounds memory read")
		}
	}()

	mem.WriteByte(1000000000, 55)
}

func TestMemoryReadWriteWord(t *testing.T) {
	mem := NewMemory()

	mem.WriteWord(10, 42)
	value := mem.ReadWord(10)

	if value != 42 {
		t.Errorf("Expected memory at address 10 to be 42, got %d", value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for out of bounds memory read")
		}
	}()

	mem.WriteByte(1000000000, 55)
}
