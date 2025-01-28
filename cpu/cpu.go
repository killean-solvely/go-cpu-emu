package cpu

import "fmt"

const RegisterCount = 4

// Register names
const (
	R0 = iota
	R1
	R2
	R3
)

var RegisterMap = map[string]uint8{
	"R0": R0,
	"R1": R1,
	"R2": R2,
	"R3": R3,
}

type Flags struct {
	Equal   uint8
	Greater uint8
	Less    uint8
}

func (f *Flags) Compare(a, b uint16) {
	if a == b {
		f.Equal = 1
	} else {
		f.Equal = 0
	}

	if a > b {
		f.Greater = 1
	} else {
		f.Greater = 0
	}

	if a < b {
		f.Less = 1
	} else {
		f.Less = 0
	}
}

type CPU struct {
	Registers      [RegisterCount]uint16
	Stack          *Stack
	Flags          Flags
	ProgramCounter uint32
}

func NewCPU() *CPU {
	cpu := &CPU{
		Registers:      [RegisterCount]uint16{},
		Stack:          NewStack(),
		ProgramCounter: 0,
	}

	return cpu
}

func (c *CPU) Execute(memory *Memory) {
	c.ProgramCounter = CodeMemoryStart

	for {
		halted := c.executeNext(memory)
		if halted {
			return
		}
	}
}

func (c *CPU) executeNext(memory *Memory) (halted bool) {
	halted = false

	opcode := memory.ReadByte(c.ProgramCounter)
	c.ProgramCounter++

	switch Opcode(opcode) {
	case OP_LOAD_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Registers[reg] = value

	case OP_LOAD_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Registers[reg1] = c.Registers[reg2]

	case OP_LOADM_RA:
		reg, address := c.prepRAInstruction(memory)
		c.Registers[reg] = memory.ReadStoredMemoryWord(uint32(address))

	case OP_STORE_RA:
		reg, address := c.prepRAInstruction(memory)
		memory.WriteStoredMemoryWord(uint32(address), c.Registers[reg])

	case OP_STORE_AV:
		address, value := c.prepAVInstruction(memory)
		memory.WriteStoredMemoryWord(uint32(address), value)

	case OP_STORE_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		memory.WriteStoredMemoryWord(uint32(c.Registers[reg1]), c.Registers[reg2])

	case OP_ADD_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Registers[reg1] += c.Registers[reg2]

	case OP_ADD_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Registers[reg] += value

	case OP_SUB_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Registers[reg1] -= c.Registers[reg2]

	case OP_SUB_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Registers[reg] -= value

	case OP_MUL_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Registers[reg1] *= c.Registers[reg2]

	case OP_MUL_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Registers[reg] *= value

	case OP_DIV_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Registers[reg1] /= c.Registers[reg2]

	case OP_DIV_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Registers[reg] /= value

	case OP_MOD_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Registers[reg1] %= c.Registers[reg2]

	case OP_MOD_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Registers[reg] %= value

	case OP_AND_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Registers[reg1] &= c.Registers[reg2]

	case OP_AND_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Registers[reg] &= value

	case OP_OR_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Registers[reg1] |= c.Registers[reg2]

	case OP_OR_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Registers[reg] |= value

	case OP_XOR_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Registers[reg1] ^= c.Registers[reg2]

	case OP_XOR_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Registers[reg] ^= value

	case OP_NOT_R:
		reg := c.prepRInstruction(memory)
		c.Registers[reg] = ^c.Registers[reg]

	case OP_SHL_R:
		reg := c.prepRInstruction(memory)
		c.Registers[reg] <<= 1

	case OP_SHR_R:
		reg := c.prepRInstruction(memory)
		c.Registers[reg] >>= 1

	case OP_INC_R:
		reg := c.prepRInstruction(memory)
		c.Registers[reg]++

	case OP_DEC_R:
		reg := c.prepRInstruction(memory)
		c.Registers[reg]--

	case OP_PUSH_R:
		reg := c.prepRInstruction(memory)
		c.Stack.Push(c.Registers[reg])

	case OP_PUSH_V:
		value := c.prepVInstruction(memory)
		c.Stack.Push(value)

	case OP_POP_NONE:
		c.prepNoneInstruction()
		c.Stack.Pop()

	case OP_POP_R:
		reg := c.prepRInstruction(memory)
		c.Registers[reg] = c.Stack.Pop()

	case OP_CMP_RR:
		reg1, reg2 := c.prepRRInstruction(memory)
		c.Flags.Compare(c.Registers[reg1], c.Registers[reg2])

	case OP_CMP_RV:
		reg, value := c.prepRVInstruction(memory)
		c.Flags.Compare(c.Registers[reg], value)

	case OP_JMP_A:
		address := c.prepAInstruction(memory)
		c.ProgramCounter = uint32(address)

	case OP_JMP_R:
		reg := c.prepRInstruction(memory)
		c.ProgramCounter = uint32(c.Registers[reg])

	case OP_JE_A:
		address := c.prepAInstruction(memory)
		if c.Flags.Equal == 1 {
			c.ProgramCounter = uint32(address)
		}

	case OP_JE_R:
		reg := c.prepRInstruction(memory)
		if c.Flags.Equal == 1 {
			c.ProgramCounter = uint32(c.Registers[reg])
		}

	case OP_JNE_A:
		address := c.prepAInstruction(memory)
		if c.Flags.Equal == 0 {
			c.ProgramCounter = uint32(address)
		}

	case OP_JNE_R:
		reg := c.prepRInstruction(memory)
		if c.Flags.Equal == 0 {
			c.ProgramCounter = uint32(c.Registers[reg])
		}

	case OP_JG_A:
		address := c.prepAInstruction(memory)
		if c.Flags.Greater == 1 {
			c.ProgramCounter = uint32(address)
		}

	case OP_JG_R:
		reg := c.prepRInstruction(memory)
		if c.Flags.Greater == 1 {
			c.ProgramCounter = uint32(c.Registers[reg])
		}

	case OP_JGE_A:
		address := c.prepAInstruction(memory)
		if c.Flags.Greater == 1 || c.Flags.Equal == 1 {
			c.ProgramCounter = uint32(address)
		}

	case OP_JGE_R:
		reg := c.prepRInstruction(memory)
		if c.Flags.Greater == 1 || c.Flags.Equal == 1 {
			c.ProgramCounter = uint32(c.Registers[reg])
		}

	case OP_JL_A:
		address := c.prepAInstruction(memory)
		if c.Flags.Less == 1 {
			c.ProgramCounter = uint32(address)
		}

	case OP_JL_R:
		reg := c.prepRInstruction(memory)
		if c.Flags.Less == 1 {
			c.ProgramCounter = uint32(c.Registers[reg])
		}

	case OP_JLE_A:
		address := c.prepAInstruction(memory)
		if c.Flags.Less == 1 || c.Flags.Equal == 1 {
			c.ProgramCounter = uint32(address)
		}

	case OP_JLE_R:
		reg := c.prepRInstruction(memory)
		if c.Flags.Less == 1 || c.Flags.Equal == 1 {
			c.ProgramCounter = uint32(c.Registers[reg])
		}

	case OP_CALL_A:
		address := c.prepAInstruction(memory)
		c.Stack.Push(uint16(c.ProgramCounter))
		c.ProgramCounter = uint32(address)

	case OP_CALL_R:
		reg := c.prepRInstruction(memory)
		c.Stack.Push(uint16(c.ProgramCounter))
		c.ProgramCounter = uint32(c.Registers[reg])

	case OP_RET_NONE:
		c.prepNoneInstruction()
		c.ProgramCounter = uint32(c.Stack.Pop())

	case OP_PRINT_V:
		value := c.prepVInstruction(memory)
		fmt.Println(value)

	case OP_PRINT_R:
		reg := c.prepRInstruction(memory)
		fmt.Println(c.Registers[reg])

	case OP_PRINTS_A:
		address := c.prepAInstruction(memory)
		// Build the string up from memory. The string is null-terminated.
		var str []byte
		for {
			value := memory.ReadStoredMemoryByte(uint32(address))
			if value == 0 {
				break
			}
			str = append(str, value)
			address++
		}
		fmt.Print(string(str))

	case OP_HLT_NONE:
		halted = true

	default:
		panic("Unknown opcode")
	}

	return halted
}

func (c *CPU) prepRRInstruction(memory *Memory) (reg1, reg2 uint8) {
	reg1 = memory.ReadByte(c.ProgramCounter)
	c.ProgramCounter++
	reg2 = memory.ReadByte(c.ProgramCounter)
	c.ProgramCounter++
	return reg1, reg2
}

func (c *CPU) prepRVInstruction(memory *Memory) (reg uint8, value uint16) {
	reg = memory.ReadByte(c.ProgramCounter)
	c.ProgramCounter++
	value = memory.ReadWord(c.ProgramCounter)
	c.ProgramCounter += 2
	return reg, value
}

func (c *CPU) prepRAInstruction(memory *Memory) (reg uint8, address uint16) {
	reg = memory.ReadByte(c.ProgramCounter)
	c.ProgramCounter++
	address = memory.ReadWord(c.ProgramCounter)
	c.ProgramCounter += 2
	return reg, address
}

func (c *CPU) prepRInstruction(memory *Memory) (reg uint8) {
	reg = memory.ReadByte(c.ProgramCounter)
	c.ProgramCounter++
	return reg
}

func (c *CPU) prepAVInstruction(memory *Memory) (address, value uint16) {
	address = memory.ReadWord(c.ProgramCounter)
	c.ProgramCounter += 2
	value = memory.ReadWord(c.ProgramCounter)
	c.ProgramCounter += 2
	return address, value
}

func (c *CPU) prepAInstruction(memory *Memory) (address uint16) {
	address = memory.ReadWord(c.ProgramCounter)
	c.ProgramCounter += 2
	return address
}

func (c *CPU) prepVInstruction(memory *Memory) (value uint16) {
	value = memory.ReadWord(c.ProgramCounter)
	c.ProgramCounter += 2
	return value
}

func (c *CPU) prepNoneInstruction() {
	// Do nothing (currently)
}
