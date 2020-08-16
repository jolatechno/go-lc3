package interfaces

import (
)

/* interface to define a cpu */
type Cpu interface {
  Memory() (memory Memory) /* get the memory of the cpu */
  Registers() (registers Registers) /* get the registers of the cpu */
  InstructionSet() (instructions InstructionSet) /* get the instructionsset of the cpu */
}
