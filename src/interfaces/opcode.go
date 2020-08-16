package interfaces

import (
)

/* interface to define an opcode */
type Op interface {
  Params() (params interface{}) /* get params */
  Instruction() (op interface{}) /* get instruction */
}
