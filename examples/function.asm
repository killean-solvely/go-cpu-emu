# Call the functions
CALL numbers
CALL lowercase
CALL uppercase
HLT

# Prints out ascii numbers
numbers:
  # '0' = 48
  LOAD R0 48
  # '9' = 57
  LOAD R1 57
  # Our index variable
  LOAD R2 0

  # Print numbers from 0 to 9
  numbers_loop:
    STORE R2 R0
    INC R0
    INC R2
    CMP R0 R1
    JLE numbers_loop

  LOAD R3 10
  STORE R2 R3
  INC R2
  LOAD R3 0
  STORE R2 R3
  PRINTS 0

  RET

# Prints out ascii lowercase letters
lowercase:
  # 'a' = 97
  LOAD R0 97
  # 'z' = 122
  LOAD R1 122
  # Our index variable
  LOAD R2 0

  # Print lowercase letters from a to z
  lowercase_loop:
    STORE R2 R0
    INC R0
    INC R2
    CMP R0 R1
    JLE lowercase_loop

  LOAD R3 10
  STORE R2 R3
  INC R2
  LOAD R3 0
  STORE R2 R3
  PRINTS 0

  RET

# Prints out ascii uppercase letters
uppercase:
  # 'A' = 65
  LOAD R0 65
  # 'Z' = 90
  LOAD R1 90
  # Our index variable
  LOAD R2 0

  # Print uppercase letters from A to Z
  uppercase_loop:
    STORE R2 R0
    INC R0
    INC R2
    CMP R0 R1
    JLE uppercase_loop

  LOAD R3 10
  STORE R2 R3
  INC R2
  LOAD R3 0
  STORE R2 R3
  PRINTS 0

  RET
