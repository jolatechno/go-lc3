package interfaces

import (
)

/* interface to define an instruction */
type Instruction interface {
  Run(memory Memory, registers Registers, op Op) (next bool, err error) /* execute an instruction given specific state and parameters */
}

/* interface to define an opcode */
type Op interface {
  Params() (params interface{}) /* get params */
  Instruction() (op interface{}) /* get instruction */
}
