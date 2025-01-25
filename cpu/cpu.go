package cpu

const RegisterCount = 4

type CPU struct {
	Registers      [RegisterCount]uint8
	Stack          *Stack
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

		case OP_HLT:
			return

		default:
			panic("Unknown opcode")
		}
	}
}
