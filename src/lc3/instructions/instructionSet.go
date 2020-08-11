package instructions

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

var Lc3instructions = [OP_COUNT]interfaces.Instruction{
  &Lc3Instruction{
    OP: OP_BR,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_ADD,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_LD,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_ST,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_JSR,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_AND,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_LDR,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_STR,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_RTI,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_NOT,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_LDI,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_STI,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_JMP,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_RES,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_LEA,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },

  &Lc3Instruction{
    OP: OP_TRAP,
    Exec: func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (err error) {
      /* TODO */
      return errors.New("Not yet implemented")
    },
  },
}
