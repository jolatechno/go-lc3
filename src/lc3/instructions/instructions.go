package instructions

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

/* defining a basic interface for lc3 instructions */
type Lc3Instruction struct {
  OP uint8 /* instruction opcode, uint4 would have been more appropriate but isn't present in standard go */
  Exec func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) /* function serving intstruction's execution */
}

/* defining the interface */
func (instru *Lc3Instruction)Run(memory interfaces.Memory, registers interfaces.Registers, op interfaces.Op) (next bool, err error) {
  params := op.Params() /* redaing parameters */
  intParam, ok := params.(uint16) /* convert param to int */
  if !ok { /* return an error if not possible */
    return false, errors.New("register not understood")
  }

  return instru.Exec(memory, registers, intParam) /* execute the instruction */
}
