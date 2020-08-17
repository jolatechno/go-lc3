package instructions

import (
)

/* defining the name of all lc3 opcode */
const (
  OP_BR uint16 = iota     /* branch */
  OP_ADD uint16 = iota    /* add  */
  OP_LD uint16 = iota     /* load */
  OP_ST uint16 = iota     /* store */
  OP_JSR uint16 = iota    /* jump register */
  OP_AND uint16 = iota    /* bitwise and */
  OP_LDR uint16 = iota    /* load register */
  OP_STR uint16 = iota    /* store register */
  OP_RTI uint16 = iota    /* unused */
  OP_NOT uint16 = iota    /* bitwise not */
  OP_LDI uint16 = iota    /* load indirect */
  OP_STI uint16 = iota    /* store indirect */
  OP_JMP uint16 = iota    /* jump */
  OP_RES uint16 = iota    /* reserved (unused) */
  OP_LEA uint16 = iota    /* load effective address */
  OP_TRAP uint16 = iota   /* execute trap */
  OP_COUNT uint16 = iota
)

const (
	TRAP_GETC  uint16 = 0x20 // get character from keyboard
	TRAP_OUT   uint16 = 0x21 // output a character
	TRAP_PUTS  uint16 = 0x22 // output a word string
	TRAP_IN    uint16 = 0x23 // input a string
	TRAP_PUTSP uint16 = 0x24 // output a byte string
	TRAP_HALT  uint16 = 0x25 // halt the program
)

const (
	MR_KBSR uint16 = 0xFE00 // keyboard status
	MR_KBDR uint16 = 0xFE02 // keyboard data
)

const (
	FL_POS uint16 = 1 << 0
	FL_ZRO uint16 = 1 << 1
	FL_NEG uint16 = 1 << 2
)
