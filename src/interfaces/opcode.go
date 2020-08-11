package interfaces

import (
)

/* interface to define an instruction set */
type InstructionSet interface {
  Get(op Op) (inst Instruction, err error)
}
