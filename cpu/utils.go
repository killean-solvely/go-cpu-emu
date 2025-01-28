package cpu

import "encoding/binary"

func uint16ToBytes(value uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, value)
	return buf
}

func bytesToUint16(data []byte) uint16 {
	return binary.BigEndian.Uint16(data)
}
