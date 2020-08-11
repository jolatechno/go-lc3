package instructions

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

/* defining the instruction set */
type Lc3InstructionSet [OP_COUNT]*Lc3Instruction

/* defining the interface */
func (set *Lc3InstructionSet)Get(op interfaces.Op) (instr interfaces.Instruction, err error) {
  intOp, ok := op.Instruction().(uint16) /* convert instruciton to int */
  if !ok { /* return an error if not possible */
    return nil, errors.New("instruction not understood")
  }

  if intOp < 0 || intOp >= OP_COUNT { /* check if instruction is not out of range */
    return nil, errors.New("instruction out of range")
  }

  instr = Lc3instructions[intOp] /* read the instruction */
  return instr, nil
}

/* instructionset definition */
var Lc3instructions = [OP_COUNT]*Lc3Instruction{
  &Lc3Instruction{
    OP: OP_BR,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_ADD,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_LD,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_ST,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_JSR,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_AND,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_LDR,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_STR,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_RTI,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_NOT,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_LDI,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_STI,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_JMP,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_RES,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_LEA,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_TRAP,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) {
      /* TODO */
      return false, errors.New("Not yet implemented")
    },
  },
}
