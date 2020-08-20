package instructions

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

/* defining a basic interface for lc3 instructions */
type Lc3Instruction struct {
  Exec func(memory interfaces.Memory, registers interfaces.Registers, param uint16) (next bool, err error) /* function serving intstruction's execution */
}

/* defining the interface */
func (instru *Lc3Instruction)Run(memory interfaces.Memory, registers interfaces.Registers, op interface{}) (next bool, err error) {
  intParam, ok := op.(uint16) /* convert op to int */
  if !ok { /* return an error if not possible */
    return false, errors.New("register not understood")
  }

  return instru.Exec(memory, registers, intParam) /* execute the instruction */
}
