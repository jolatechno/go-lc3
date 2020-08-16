package instructions

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
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

/* update flag */
func update_flags(res uint16, regs interfaces.Registers) (err error) {
    if (res == 0) {
        return regs.Write(registers.R_COND, registers.FL_ZRO)
    } else if res >> 15 == 1 { /* a 1 in the left-most bit indicates negative */
        return regs.Write(registers.R_COND, registers.FL_NEG)
    }
    return regs.Write(registers.R_COND, registers.FL_POS)
}

/* sign_extend */
func sign_extend(x uint16, bit_count int) uint16 {
    if (x >> (bit_count - 1)) & 1 == 1 {
        x |= (0xFFFF << bit_count);
    }
    return x;
}
