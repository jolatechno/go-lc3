package instructions

import (
  "errors"
)

/* defining the type of lc3 opcode */
type Lc3OP struct {
  Value uint16 /* memory type */
}

/* defining the interface */
func (opcode *Lc3OP)Params() (params []interface{}) {
  return []interface{}{ opcode.Value } /* return parameters */
}

func (opcode *Lc3OP)Instruction() (op interface{}) {
  return opcode.Value >> 12 /* return the opcode */
}

func (opcode *Lc3OP)ToMemory() (value interface{}) {
  return opcode.Value /* return memory value */
}

func (opcode *Lc3OP)FromMemory(value interface{}) (err error) {
  intValue, ok := value.(uint16) /* convert value to int */
  if !ok { /* return an error if not possible */
    return errors.New("register not understood")
  }

  opcode.Value = intValue /* assign value */
  return nil
}
