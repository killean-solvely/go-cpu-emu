package cpu

type Opcode uint8

const (
	OP_LOAD          Opcode = iota // Load a value into a register
	OP_STORE_VAL                   // Store a value in stored memory
	OP_STORE_REG                   // Store a value from a register in stored memory
	OP_STORE_REG_REG               // Store a value from a register in stored memory, using a register as the address
	OP_LOAD_MEM                    // Load a value from stored memory into a register
	OP_ADD                         // Add values from two registers
	OP_SUB                         // Subtract values from two registers
	OP_MUL                         // Multiply values from two registers
	OP_DIV                         // Divide values from two registers
	OP_MOD                         // Modulo values from two registers
	OP_AND                         // Bitwise AND values from two registers
	OP_OR                          // Bitwise OR values from two registers
	OP_XOR                         // Bitwise XOR values from two registers
	OP_NOT                         // Bitwise NOT values from a register
	OP_SHL                         // Bitwise shift left values from a register
	OP_SHR                         // Bitwise shift right values from a register
	OP_INC                         // Increment a register
	OP_DEC                         // Decrement a register
	OP_JMP                         // Jump to an address
	OP_JMP_REG                     // Jump to an address in a register
	OP_PUSH                        // Push a value onto the stack
	OP_PUSH_REG                    // Push a value from a register onto the stack
	OP_POP                         // Pop a value from the stack
	OP_POP_REG                     // Pop a value from the stack into a register
	OP_CMP_REG_VAL                 // Compare a register with a value
	OP_CMP_REG_REG                 // Compare two registers
	OP_JE                          // Jump if equal
	OP_JNE                         // Jump if not equal
	OP_JG                          // Jump if greater
	OP_JGE                         // Jump if greater or equal
	OP_JL                          // Jump if less
	OP_JLE                         // Jump if less or equal
	OP_CALL                        // Call a function. Pushes the next instruction address onto the stack
	OP_RET                         // Return from a function. Uses the last value on the stack as the return address
	OP_PRINT                       // Print a value
	OP_PRINT_REG                   // Print a value from a register
	OP_PRINT_MEM                   // Print a string from stored memory
	OP_HLT                         // Halt execution
)

type InstructionType uint8

// Instruction types
const (
	INST_R_R InstructionType = iota
	INST_R_V
	INST_R_A
	INST_R_L
	INST_A_V
	INST_A_L
	INST_A
	INST_V
	INST_R
	INST_NONE
)

type OpcodeDefinition struct {
	Type   InstructionType
	Opcode Opcode
}

var OpcodeMap = map[string]OpcodeDefinition{
	"LOAD":          {Type: INST_R_V, Opcode: OP_LOAD},
	"STORE_VAL":     {Type: INST_A_V, Opcode: OP_STORE_VAL},
	"STORE_REG":     {Type: INST_R_A, Opcode: OP_STORE_REG},
	"STORE_REG_REG": {Type: INST_R_R, Opcode: OP_STORE_REG_REG},
	"LOAD_MEM":      {Type: INST_R_A, Opcode: OP_LOAD_MEM},
	"ADD":           {Type: INST_R_R, Opcode: OP_ADD},
	"SUB":           {Type: INST_R_R, Opcode: OP_SUB},
	"MUL":           {Type: INST_R_R, Opcode: OP_MUL},
	"DIV":           {Type: INST_R_R, Opcode: OP_DIV},
	"MOD":           {Type: INST_R_R, Opcode: OP_MOD},
	"AND":           {Type: INST_R_R, Opcode: OP_AND},
	"OR":            {Type: INST_R_R, Opcode: OP_OR},
	"XOR":           {Type: INST_R_R, Opcode: OP_XOR},
	"NOT":           {Type: INST_R, Opcode: OP_NOT},
	"SHL":           {Type: INST_R, Opcode: OP_SHL},
	"SHR":           {Type: INST_R, Opcode: OP_SHR},
	"INC":           {Type: INST_R, Opcode: OP_INC},
	"DEC":           {Type: INST_R, Opcode: OP_DEC},
	"JMP":           {Type: INST_A_L, Opcode: OP_JMP},
	"JMP_REG":       {Type: INST_R_L, Opcode: OP_JMP_REG},
	"PUSH":          {Type: INST_V, Opcode: OP_PUSH},
	"PUSH_REG":      {Type: INST_R, Opcode: OP_PUSH_REG},
	"POP":           {Type: INST_NONE, Opcode: OP_POP},
	"POP_REG":       {Type: INST_R, Opcode: OP_POP_REG},
	"CMP_REG_VAL":   {Type: INST_R_V, Opcode: OP_CMP_REG_VAL},
	"CMP_REG_REG":   {Type: INST_R_R, Opcode: OP_CMP_REG_REG},
	"JE":            {Type: INST_A_L, Opcode: OP_JE},
	"JNE":           {Type: INST_A_L, Opcode: OP_JNE},
	"JG":            {Type: INST_A_L, Opcode: OP_JG},
	"JGE":           {Type: INST_A_L, Opcode: OP_JGE},
	"JL":            {Type: INST_A_L, Opcode: OP_JL},
	"JLE":           {Type: INST_A_L, Opcode: OP_JLE},
	"CALL":          {Type: INST_A_L, Opcode: OP_CALL},
	"RET":           {Type: INST_NONE, Opcode: OP_RET},
	"PRINT":         {Type: INST_V, Opcode: OP_PRINT},
	"PRINT_REG":     {Type: INST_R, Opcode: OP_PRINT_REG},
	"PRINT_MEM":     {Type: INST_A, Opcode: OP_PRINT_MEM},
	"HLT":           {Type: INST_NONE, Opcode: OP_HLT},
}

var InstructionSizeMap = map[InstructionType]int{
	INST_R_R:  3,
	INST_R_V:  3,
	INST_R_A:  3,
	INST_R_L:  2,
	INST_A_V:  3,
	INST_A_L:  2,
	INST_A:    2,
	INST_V:    2,
	INST_R:    2,
	INST_NONE: 1,
}
