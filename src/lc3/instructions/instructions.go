package instructions

import (

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
type lc3instruction struct {
  OP uint8 /* instruction opcode, uint4 would have been more appropriate but isn't present in standard go */
  Exec func( /* function serving intstruction's execution */
    Memory interface{},
    Registers interface{},
  )
}

/* and an array to store them */
var lc3instructions [OP_COUNT]*lc3instruction; // later defined in other files (one per instruction)
