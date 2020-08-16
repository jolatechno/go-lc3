package opcode

import (
)

/* defining the type of lc3 opcode */
type Lc3OP struct {
  Value uint16 /* memory type */
}

/* defining the interface */
func (opcode *Lc3OP)Params() (params interface{}) {
  return opcode.Value /* return parameters */
}

func (opcode *Lc3OP)Instruction() (op interface{}) {
  return opcode.Value >> 12 /* return the opcode */
}
