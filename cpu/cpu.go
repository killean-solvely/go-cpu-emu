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

func (f *Flags) Compare(a, b uint8) {
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
	Registers      [RegisterCount]uint8
	Stack          *Stack
	Flags          Flags
	ProgramCounter uint16
}

func NewCPU() *CPU {
	cpu := &CPU{
		Registers:      [RegisterCount]uint8{},
		Stack:          NewStack(),
		ProgramCounter: 0,
	}

	return cpu
}

func (c *CPU) Execute(memory *Memory) {
	c.ProgramCounter = CodeMemoryStart

	for {
		opcode := memory.Read(c.ProgramCounter)
		c.ProgramCounter++

		switch Opcode(opcode) {
		case OP_LOAD:
			reg, value := c.prepRVInstruction(memory)
			c.Registers[reg] = value

		case OP_STORE_VAL:
			address, value := c.prepAVInstruction(memory)
			memory.WriteStoredMemory(uint16(address), value)

		case OP_STORE_REG:
			reg, address := c.prepRAInstruction(memory)
			memory.WriteStoredMemory(uint16(address), c.Registers[reg])

		case OP_LOAD_MEM:
			reg, address := c.prepRAInstruction(memory)
			c.Registers[reg] = memory.ReadStoredMemory(uint16(address))

		case OP_ADD:
			reg1, reg2 := c.prepRRInstruction(memory)
			c.Registers[reg1] += c.Registers[reg2]

		case OP_SUB:
			reg1, reg2 := c.prepRRInstruction(memory)
			c.Registers[reg1] -= c.Registers[reg2]

		case OP_MUL:
			reg1, reg2 := c.prepRRInstruction(memory)
			c.Registers[reg1] *= c.Registers[reg2]

		case OP_DIV:
			reg1, reg2 := c.prepRRInstruction(memory)
			c.Registers[reg1] /= c.Registers[reg2]

		case OP_MOD:
			reg1, reg2 := c.prepRRInstruction(memory)
			c.Registers[reg1] %= c.Registers[reg2]

		case OP_AND:
			reg1, reg2 := c.prepRRInstruction(memory)
			c.Registers[reg1] &= c.Registers[reg2]

		case OP_OR:
			reg1, reg2 := c.prepRRInstruction(memory)
			c.Registers[reg1] |= c.Registers[reg2]

		case OP_XOR:
			reg1, reg2 := c.prepRRInstruction(memory)
			c.Registers[reg1] ^= c.Registers[reg2]

		case OP_NOT:
			reg := c.prepRInstruction(memory)
			c.Registers[reg] = ^c.Registers[reg]

		case OP_SHL:
			reg := c.prepRInstruction(memory)
			c.Registers[reg] <<= 1

		case OP_SHR:
			reg := c.prepRInstruction(memory)
			c.Registers[reg] >>= 1

		case OP_INC:
			reg := c.prepRInstruction(memory)
			c.Registers[reg]++

		case OP_DEC:
			reg := c.prepRInstruction(memory)
			c.Registers[reg]--

		case OP_JMP:
			address := c.prepAInstruction(memory)
			c.ProgramCounter = uint16(address)

		case OP_JMP_REG:
			reg := c.prepRInstruction(memory)
			c.ProgramCounter = uint16(c.Registers[reg])

		case OP_PUSH:
			value := c.prepVInstruction(memory)
			c.Stack.Push(value)

		case OP_PUSH_REG:
			reg := c.prepRInstruction(memory)
			c.Stack.Push(c.Registers[reg])

		case OP_POP:
			c.prepNoneInstruction()
			c.Stack.Pop()

		case OP_POP_REG:
			reg := c.prepRInstruction(memory)
			c.Registers[reg] = c.Stack.Pop()

		case OP_CMP_REG_VAL:
			reg, value := c.prepRVInstruction(memory)
			c.Flags.Compare(c.Registers[reg], value)

		case OP_CMP_REG_REG:
			reg1, reg2 := c.prepRRInstruction(memory)
			c.Flags.Compare(c.Registers[reg1], c.Registers[reg2])

		case OP_JE:
			address := c.prepAInstruction(memory)
			if c.Flags.Equal == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JNE:
			address := c.prepAInstruction(memory)
			if c.Flags.Equal == 0 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JG:
			address := c.prepAInstruction(memory)
			if c.Flags.Greater == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JGE:
			address := c.prepAInstruction(memory)
			if c.Flags.Greater == 1 || c.Flags.Equal == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JL:
			address := c.prepAInstruction(memory)
			if c.Flags.Less == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JLE:
			address := c.prepAInstruction(memory)
			if c.Flags.Less == 1 || c.Flags.Equal == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_PRINT:
			value := c.prepVInstruction(memory)
			fmt.Println(value)

		case OP_PRINT_REG:
			reg := c.prepRInstruction(memory)
			fmt.Println(c.Registers[reg])

		case OP_PRINT_MEM:
			address := c.prepAInstruction(memory)
			// Build the string up from memory. The string is null-terminated.
			var str []byte
			for {
				value := memory.ReadStoredMemory(uint16(address))
				if value == 0 {
					break
				}
				str = append(str, value)
				address++
			}
			fmt.Println(string(str))

		case OP_HLT:
			return

		default:
			panic("Unknown opcode")
		}
	}
}

func (c *CPU) prepRRInstruction(memory *Memory) (reg1, reg2 uint8) {
	reg1 = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	reg2 = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	return reg1, reg2
}

func (c *CPU) prepRVInstruction(memory *Memory) (reg uint8, value uint8) {
	reg = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	value = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	return reg, value
}

func (c *CPU) prepRAInstruction(memory *Memory) (reg, address uint8) {
	reg = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	address = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	return reg, address
}

func (c *CPU) prepRInstruction(memory *Memory) (reg uint8) {
	reg = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	return reg
}

func (c *CPU) prepAVInstruction(memory *Memory) (address, value uint8) {
	address = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	value = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	return address, value
}

func (c *CPU) prepAInstruction(memory *Memory) (address uint8) {
	address = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	return address
}

func (c *CPU) prepVInstruction(memory *Memory) (value uint8) {
	value = memory.Read(c.ProgramCounter)
	c.ProgramCounter++
	return value
}

func (c *CPU) prepNoneInstruction() {
	c.ProgramCounter++
}
