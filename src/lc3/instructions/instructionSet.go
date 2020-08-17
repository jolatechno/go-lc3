package instructions

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

/* defining the instruction set */
type Lc3InstructionSet struct {
  Instructions [OP_COUNT]*Lc3Instruction
}

/* defining the interface */
func (set *Lc3InstructionSet)Get(op interfaces.Op) (instr interfaces.Instruction, err error) {
  intOp, ok := op.Instruction().(uint16) /* convert instruciton to int */
  if !ok { /* return an error if not possible */
    return nil, errors.New("instruction not understood")
  }

  if intOp < 0 || intOp >= OP_COUNT { /* check if instruction is not out of range */
    return nil, errors.New("instruction out of range")
  }

  instr = set.Instructions[intOp] /* read the instruction */
  return instr, nil
}
