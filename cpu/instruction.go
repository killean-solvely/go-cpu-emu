package cpu

type Opcode uint8

const (
	OP_LOAD_RV  Opcode = iota // Load a value into a register
	OP_LOAD_RR                // Load a register into another register
	OP_LOADM_RA               // Load a value from stored memory into a register
	OP_STORE_RA               // Stores a register in stored memory
	OP_STORE_AV               // Stores a value in stored memory
	OP_STORE_RR               // Stores register2 in the address that register1 contains
	OP_ADD_RR                 // Add the right register to the left register, storing the result in the left register
	OP_ADD_RV                 // Add the register and value together, storing in the register
	OP_SUB_RR                 // Subtract the right register from the left register, storing the result in the left register
	OP_SUB_RV                 // Subtract the value from the register, storing in the register
	OP_MUL_RR                 // Multiply the left register by the right register, storing the result in the left register
	OP_MUL_RV                 // Multiply the register by the value, storing in the register
	OP_DIV_RR                 // Divide the left register by the right register, storing the result in the left register
	OP_DIV_RV                 // Divide the register by the value, storing in the register
	OP_MOD_RR                 // Modulo the left register by the right register, storing the result in the left register
	OP_MOD_RV                 // Modulo the register by the value, storing in the register
	OP_AND_RR                 // Bitwise AND the left register with the right register, storing the result in the left register
	OP_AND_RV                 // Bitwise AND the register with the value, storing in the register
	OP_OR_RR                  // Bitwise OR the left register with the right register, storing the result in the left register
	OP_OR_RV                  // Bitwise OR the register with the value, storing in the register
	OP_XOR_RR                 // Bitwise XOR the left register with the right register, storing the result in the left register
	OP_XOR_RV                 // Bitwise XOR the register with the value, storing in the register
	OP_NOT_R                  // Bitwise NOT a register
	OP_SHL_R                  // Bitwise shift left a register
	OP_SHR_R                  // Bitwise shift right a register
	OP_INC_R                  // Increment a register
	OP_DEC_R                  // Decrement a register
	OP_PUSH_R                 // Push a register onto the stack
	OP_PUSH_V                 // Push a value onto the stack
	OP_POP_NONE               // Pop a value off the stack
	OP_POP_R                  // Pop a value off the stack into a register
	OP_CMP_RR                 // Compare two registers, setting the flags
	OP_CMP_RV                 // Compare a register to a value, setting the flags
	OP_JMP_A                  // Jump to an address
	OP_JMP_R                  // Jump to a register
	OP_JE_A                   // Jump if equal to an address
	OP_JE_R                   // Jump if equal to an register
	OP_JNE_A                  // Jump if not equal to an address
	OP_JNE_R                  // Jump if not equal to an register
	OP_JG_A                   // Jump if greater than an address
	OP_JG_R                   // Jump if greater than an register
	OP_JGE_A                  // Jump if greater than or equal to an address
	OP_JGE_R                  // Jump if greater than or equal to an register
	OP_JL_A                   // Jump if less than an address
	OP_JL_R                   // Jump if less than an register
	OP_JLE_A                  // Jump if less than or equal to an address
	OP_JLE_R                  // Jump if less than or equal to an register
	OP_CALL_A                 // Call a function at an address, pushing the next instruction onto the stack
	OP_CALL_R                 // Call a function at a register, pushing the next instruction onto the stack
	OP_RET_NONE               // Return from a function, popping the return address off the stack and jumping to it
	OP_PRINT_V                // Print a value
	OP_PRINT_R                // Print a register
	OP_PRINTS_A               // Print a string from stored memory
	OP_HLT_NONE               // Halt execution
)

type InstructionType uint8

// Instruction types
const (
	INST_R InstructionType = iota
	INST_RR
	INST_RA
	INST_RV
	INST_A
	INST_AV
	INST_V
	INST_NONE

	INST_AL
	INST_RL
)

type OpcodeKey struct {
	OpcodeName string
	Type       InstructionType
}

var OpcodeMap = map[OpcodeKey]Opcode{
	{"LOAD", INST_RV}:  OP_LOAD_RV,
	{"LOAD", INST_RR}:  OP_LOAD_RR,
	{"LOADM", INST_RA}: OP_LOADM_RA,
	{"STORE", INST_RA}: OP_STORE_RA,
	{"STORE", INST_AV}: OP_STORE_AV,
	{"STORE", INST_RR}: OP_STORE_RR,
	{"ADD", INST_RR}:   OP_ADD_RR,
	{"ADD", INST_RV}:   OP_ADD_RV,
	{"SUB", INST_RR}:   OP_SUB_RR,
	{"SUB", INST_RV}:   OP_SUB_RV,
	{"MUL", INST_RR}:   OP_MUL_RR,
	{"MUL", INST_RV}:   OP_MUL_RV,
	{"DIV", INST_RR}:   OP_DIV_RR,
	{"DIV", INST_RV}:   OP_DIV_RV,
	{"MOD", INST_RR}:   OP_MOD_RR,
	{"MOD", INST_RV}:   OP_MOD_RV,
	{"AND", INST_RR}:   OP_AND_RR,
	{"AND", INST_RV}:   OP_AND_RV,
	{"OR", INST_RR}:    OP_OR_RR,
	{"OR", INST_RV}:    OP_OR_RV,
	{"XOR", INST_RR}:   OP_XOR_RR,
	{"XOR", INST_RV}:   OP_XOR_RV,
	{"NOT", INST_R}:    OP_NOT_R,
	{"SHL", INST_R}:    OP_SHL_R,
	{"SHR", INST_R}:    OP_SHR_R,
	{"INC", INST_R}:    OP_INC_R,
	{"DEC", INST_R}:    OP_DEC_R,
	{"PUSH", INST_R}:   OP_PUSH_R,
	{"PUSH", INST_V}:   OP_PUSH_V,
	{"POP", INST_NONE}: OP_POP_NONE,
	{"POP", INST_R}:    OP_POP_R,
	{"CMP", INST_RR}:   OP_CMP_RR,
	{"CMP", INST_RV}:   OP_CMP_RV,
	{"JMP", INST_AL}:   OP_JMP_A,
	{"JMP", INST_A}:    OP_JMP_A,
	{"JMP", INST_R}:    OP_JMP_R,
	{"JE", INST_AL}:    OP_JE_A,
	{"JE", INST_A}:     OP_JE_A,
	{"JE", INST_R}:     OP_JE_R,
	{"JNE", INST_AL}:   OP_JNE_A,
	{"JNE", INST_A}:    OP_JNE_A,
	{"JNE", INST_R}:    OP_JNE_R,
	{"JG", INST_AL}:    OP_JG_A,
	{"JG", INST_A}:     OP_JG_A,
	{"JG", INST_R}:     OP_JG_R,
	{"JGE", INST_AL}:   OP_JGE_A,
	{"JGE", INST_A}:    OP_JGE_A,
	{"JGE", INST_R}:    OP_JGE_R,
	{"JL", INST_AL}:    OP_JL_A,
	{"JL", INST_A}:     OP_JL_A,
	{"JL", INST_R}:     OP_JL_R,
	{"JLE", INST_AL}:   OP_JLE_A,
	{"JLE", INST_A}:    OP_JLE_A,
	{"JLE", INST_R}:    OP_JLE_R,
	{"CALL", INST_AL}:  OP_CALL_A,
	{"CALL", INST_A}:   OP_CALL_A,
	{"CALL", INST_R}:   OP_CALL_R,
	{"RET", INST_NONE}: OP_RET_NONE,
	{"PRINT", INST_V}:  OP_PRINT_V,
	{"PRINT", INST_R}:  OP_PRINT_R,
	{"PRINTS", INST_A}: OP_PRINTS_A,
	{"HLT", INST_NONE}: OP_HLT_NONE,
}

var InstructionSizeMap = map[InstructionType]int{
	INST_RR:   3,
	INST_RV:   3,
	INST_RA:   3,
	INST_AV:   3,
	INST_A:    2,
	INST_V:    2,
	INST_R:    2,
	INST_NONE: 1,

	INST_AL: 2,
	INST_RL: 2,
}
