package cpu

const RegisterCount = 4

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
	return &CPU{
		Registers:      [RegisterCount]uint8{},
		Stack:          NewStack(),
		ProgramCounter: 0,
	}
}

func (c *CPU) Execute(memory *Memory) {
	c.ProgramCounter = CodeMemoryStart

	for {
		opcode := memory.Read(c.ProgramCounter)
		c.ProgramCounter++

		switch opcode {
		case OP_LOAD:
			reg := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			value := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			c.Registers[reg] = value

		case OP_STORE_VAL:
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			value := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			memory.WriteStoredMemory(uint16(address), value)

		case OP_STORE_REG:
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			reg := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			memory.WriteStoredMemory(uint16(address), c.Registers[reg])

		case OP_LOAD_MEM:
			reg := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			c.Registers[reg] = memory.ReadStoredMemory(uint16(address))

		case OP_ADD:
			reg1 := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			reg2 := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			c.Registers[reg1] += c.Registers[reg2]

		case OP_JMP:
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter = uint16(address)

		case OP_PUSH:
			value := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			c.Stack.Push(value)

		case OP_PUSH_REG:
			reg := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			c.Stack.Push(c.Registers[reg])

		case OP_POP:
			c.Stack.Pop()
			c.ProgramCounter++

		case OP_POP_REG:
			reg := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			c.Registers[reg] = c.Stack.Pop()

		case OP_CMP_REG_VAL:
			reg := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			value := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			c.Flags.Compare(c.Registers[reg], value)

		case OP_CMP_REG_REG:
			reg1 := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			reg2 := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			c.Flags.Compare(c.Registers[reg1], c.Registers[reg2])

		case OP_JE:
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			if c.Flags.Equal == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JNE:
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			if c.Flags.Equal == 0 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JG:
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			if c.Flags.Greater == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JGE:
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			if c.Flags.Greater == 1 || c.Flags.Equal == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JL:
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			if c.Flags.Less == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_JLE:
			address := memory.Read(c.ProgramCounter)
			c.ProgramCounter++
			if c.Flags.Less == 1 || c.Flags.Equal == 1 {
				c.ProgramCounter = uint16(address)
			}

		case OP_HLT:
			return

		default:
			panic("Unknown opcode")
		}
	}
}
