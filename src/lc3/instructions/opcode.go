package instructions

import (
  "errors"
)

/* defining the type of lc3 opcode */
type Lc3OP struct {
  Op int8 /* instruction opcode, uint4 would have been more appropriate but isn't present in standard go */
  Parms []int16 /* list of additional parameters */
}

/* defining the interface */
func (opcode *Lc3OP)Params() (params []interface{}) {
  /* TODO */
  return nil
}

func (opcode *Lc3OP)Instruction() (op interface{}) {
  /* TODO */
  return nil
}

func (opcode *Lc3OP)ToMemory() (value interface{}) {
  /* TODO */
  return nil
}

func (opcode *Lc3OP)FromMemory(value interface{}) (err error) {
  /* TODO */
  return errors.New("Not implemented")
}
