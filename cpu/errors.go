package cpu

import "fmt"

type AssemblerErrorType string

const (
	INVALID_OPERAND_COUNT AssemblerErrorType = "invalid operand count"
	INVALID_ADDRESS                          = "invalid address"
	INVALID_VALUE                            = "invalid value"
	INVALID_REGISTER                         = "invalid register"
	INVALID_LABEL                            = "invalid label"
	INVALID_OPCODE                           = "invalid opcode"
)

type AssemblerError struct {
	Type    AssemblerErrorType
	Opcode  uint8
	Opname  string
	Message string
	Line    int
}

func NewAssemblerError(
	t AssemblerErrorType,
	l int,
	opcode uint8,
	opname string,
	message string,
) *AssemblerError {
	return &AssemblerError{
		Type:    t,
		Opcode:  opcode,
		Opname:  opname,
		Line:    l,
		Message: message,
	}
}

func (e *AssemblerError) Error() string {
	return fmt.Sprintf("Assembler error on line %d: %s - %s", e.Line+1, e.Type, e.Message)
}
