package cpu

import "testing"

func TestMemoryReadWrite(t *testing.T) {
	mem := NewMemory()

	mem.Write(10, 42)
	value := mem.Read(10)

	if value != 42 {
		t.Errorf("Expected memory at address 10 to be 42, got %d", value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for out of bounds memory read")
		}
	}()

	mem.Write(300, 55)
}
