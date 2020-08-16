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
    OP_LDR = iota    /* load register */ /* TODO */
    OP_STR = iota    /* store register */ /* TODO */
    OP_RTI = iota    /* unused */ /* TODO */
    OP_NOT = iota    /* bitwise not */ /* TODO */
    OP_LDI = iota    /* load indirect */ /* TODO */
    OP_STI = iota    /* store indirect */ /* TODO */
    OP_JMP = iota    /* jump */
    OP_RES = iota    /* reserved (unused) */ /* TODO */
    OP_LEA = iota    /* load effective address */ /* TODO */
    OP_TRAP = iota   /* execute trap */
    OP_COUNT = iota
)

const (
	TRAP_GETC  = 0x20 // get character from keyboard
	TRAP_OUT   = 0x21 // output a character
	TRAP_PUTS  = 0x22 // output a word string
	TRAP_IN    = 0x23 // input a string
	TRAP_PUTSP = 0x24 // output a byte string
	TRAP_HALT  = 0x25 // halt the program
)

const (
	MR_KBSR = 0xFE00 // keyboard status
	MR_KBDR = 0xFE02 // keyboard data
)

const (
	FL_POS = 1 << 0
	FL_ZRO = 1 << 1
	FL_NEG = 1 << 2
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
