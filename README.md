# Instruction Set

Order follows for each instructions
```
OP REG
OP REG REG
OP REG ADDR
OP REG VAL
OP ADDR
OP ADDR VAL
OP VAL
OP
```


Load a value into a register

`LOAD REG VAL`

Load a value from memory into a register

`LOADM REG ADDR`

Store a register into memory

`STORE REG ADDR`

Store a value into memory

`STORE ADDR VAL`

Store the right register at the left registers value as the memory address

`STORE REG REG`

Add two registers and store the result in the left register

`ADD REG REG`

Add a value to a register and store the result in the register

`ADD REG VAL`

Subtract two registers and store the result in the left register

`SUB REG REG`

Subtract a value from a register and store the result in the register

`SUB REG VAL`

Multiply two registers and store the result in the left register

`MUL REG REG`

Multiply a register by a value and store the result in the register

`MUL REG VAL`

Divide two registers and store the result in the left register

`DIV REG REG`

Divide a register by a value and store the result in the register

`DIV REG VAL`

Modulo two registers and store the result in the left register

`MOD REG REG`

Modulo a register by a value and store the result in the register

`MOD REG VAL`

And two registers and store the result in the left register

`AND REG REG`

And a register by a value and store the result in the register

`AND REG VAL`

Or two registers and store the result in the left register

`OR REG REG`

Or a register by a value and store the result in the register

`OR REG VAL`

Xor two registers and store the result in the left register

`XOR REG REG`

Xor a register by a value and store the result in the register

`XOR REG VAL`

Not a register and store the result in the register

`NOT REG`

Shift a register left one bit and store the result in the register

`SHL REG`

Shift a register right one bit and store the result in the register

`SHR REG`

Increment a register by one

`INC REG`

Decrement a register by one

`DEC REG`

Push a value onto the stack

`PUSH VAL`

Push the value in a register onto the stack

`PUSH REG`

Pop a value off the stack (destroys the value)

`POP`

Pop a value off the stack and store it in a register

`POP REG`

Compares two registers, setting the flags

`CMP REG REG`

Compares a register to a value, setting the flags

`CMP REG VAL`

Jump to an address

`JMP ADDR`

Jump to a register

`JMP REG`

Jump to an address if the equal flag is set

`JE ADDR`

Jump to a register if the equal flag is set

`JE REG`

Jump to an address if the not equal flag is set

`JNE ADDR`

Jump to a register if the not equal flag is set

`JNE REG`

Jump to an address if the greater flag is set

`JG ADDR`

Jump to a register if the greater flag is set

`JG REG`

Jump to an address if the greater flag and equal flag are set

`JGE ADDR`

Jump to a register if the greater flag and equal flag are set

`JGE REG`

Jump to an address if the less flag is set

`JL ADDR`

Jump to a register if the less flag is set

`JL REG`

Jump to an address if the less flag and equal flag are set

`JLE ADDR`

Jump to a register if the less flag and equal flag are set

`JLE REG`

Call a function at an address

`CALL ADDR`

Call a function at a register

`CALL REG`

Return from a function

`RET`

Print a value

`PRINT VAL`

Print the value in a register

`PRINT REG`

Print a string at an address

`PRINTS ADDR`

End the program

`HLT`
