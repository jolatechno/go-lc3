package instructions

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

/* defining the name of all lc3 opcode */
const (
    OP_BR = iota     /* branch */
    OP_ADD = iota    /* add  */
    OP_LD = iota     /* load */
    OP_ST = iota     /* store */
    OP_JSR = iota    /* jump register */
    OP_AND = iota    /* bitwise and */
    OP_LDR = iota    /* load register */
    OP_STR = iota    /* store register */
    OP_RTI = iota    /* unused */
    OP_NOT = iota    /* bitwise not */
    OP_LDI = iota    /* load indirect */
    OP_STI = iota    /* store indirect */
    OP_JMP = iota    /* jump */
    OP_RES = iota    /* reserved (unused) */
    OP_LEA = iota    /* load effective address */
    OP_TRAP = iota   /* execute trap */
    OP_COUNT = iota
)

/* defining a basic interface for lc3 instructions */
type Lc3Instruction struct {
  OP uint8 /* instruction opcode, uint4 would have been more appropriate but isn't present in standard go */
  Exec func(memory interfaces.Memory, interfaces interfaces.Registers, params []interface{}) (err error) /* function serving intstruction's execution */
}

/* and an array to store them */
var Lc3instructions [OP_COUNT]*Lc3Instruction; // later defined in other files (one per instruction)

/* defining the interface */
func (instru *Lc3Instruction)Run(memory interface{}, registers interface{}, params []interface{}) (err error) {
  /* TODO */
  return errors.New("Not implemented")
}
