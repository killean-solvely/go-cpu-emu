package cpu

const (
	OP_LOAD        = iota // Load a value into a register
	OP_STORE_VAL          // Store a value in stored memory
	OP_STORE_REG          // Store a value from a register in stored memory
	OP_LOAD_MEM           // Load a value from stored memory into a register
	OP_ADD                // Add values from two registers
	OP_JMP                // Jump to an address
	OP_PUSH               // Push a value onto the stack
	OP_PUSH_REG           // Push a value from a register onto the stack
	OP_POP                // Pop a value from the stack
	OP_POP_REG            // Pop a value from the stack into a register
	OP_CMP_REG_VAL        // Compare a register with a value
	OP_CMP_REG_REG        // Compare two registers
	OP_JE                 // Jump if equal
	OP_JNE                // Jump if not equal
	OP_JG                 // Jump if greater
	OP_JGE                // Jump if greater or equal
	OP_JL                 // Jump if less
	OP_JLE                // Jump if less or equal
	OP_HLT                // Halt execution
)

type Instruction struct {
	Opcode   uint8
	Operand1 uint8
	Operand2 uint8
}
