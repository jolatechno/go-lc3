package instructions

import (
  "errors"
  "fmt"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
)

/* instructionset definition */
var Lc3instructionSet = Lc3InstructionSet {
  [OP_COUNT]*Lc3Instruction {
    &Lc3Instruction { /* OP_BR instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        pc_offset := signExtend(param & 0x1ff, 9)
  			flag := (param >> 9) & 0x7

        cond := readReg(regs, registers.R_COND) /* read cond register */
  			if (flag & cond) != 0 {
          pc := readReg(regs, registers.R_PC) /* read pc register */
  				handle(regs.Write(registers.R_COND, pc + pc_offset)) /* write pc to rtegister */
  			}

        return true, err
      },
    },

    &Lc3Instruction { /* OP_ADD instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        dr := (param >> 9) & 0x7
  			sr1 := (param >> 6) & 0x7
  			mode := (param >> 5) & 0x1
        var res uint16

        r1 := readReg(regs, sr1) /* read first register */
        if mode == 1 { /* second param is a value */
            immediate := signExtend(param & 0x1F, 5) /* read value */
            res = r1 + immediate /* compute value */
        } else { /* second param is a register */
            sr2 := param & 0x7 /* read register */
            r2 := readReg(regs, sr2) /* read second register */
            res = r1 + r2 /* compute value */
        }

        handle(regs.Write(dr, res)) /* write result to register */
        handle(updateFlags(regs, res)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_LD instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        dr := (param >> 9) & 0x7
  			pc_offset := signExtend(param & 0x1ff, 9)

        pc := readReg(regs, registers.R_PC) /* read pc register */
  			res := readMem(memory, pc + pc_offset) /* read value off of memory */

        handle(regs.Write(dr, res)) /* write res to register */
  			handle(updateFlags(regs, res)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_ST instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        sr := (param >> 9) & 0x7
  			pc_offset := signExtend(param & 0x1ff, 9)

        pc := readReg(regs, registers.R_PC) /* read pc register */
        s := readReg(regs, sr) /* read sr register */

        handle(memory.Write(pc + pc_offset, s)) /* write to memory */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_JSR instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        r := (param >> 6) & 0x7
  			long_pc_Offset := signExtend(param & 0x7ff, 11)
  			long_flag := (param >> 11) & 1

        pc := readReg(regs, registers.R_PC) /* read pc register */
  			handle(regs.Write(registers.R_R7, pc)) /* write it to R7 */
  			if long_flag != 0 {
  				handle(regs.Write(registers.R_PC, pc + long_pc_Offset)) /*increment by long offset */
  			} else {
          v := readReg(regs, r) /* read r register */
  				handle(regs.Write(registers.R_PC, v)) /* assign v to pc */
  			}

        return true, err
      },
    },

    &Lc3Instruction { /* OP_AND instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        dr := (param >> 9) & 0x7
  			sr1 := (param >> 6) & 0x7
  			mode := (param >> 5) & 0x1
        var res uint16

        r1 := readReg(regs, sr1) /* read first register */
  			if mode == 1 {
  				imm5 := signExtend(param & 0x1F, 5)
  				res =  r1 & imm5 /* compute value */
  			} else {
  				sr2 := param & 0x7
          r2 := readReg(regs, sr2) /* read second register */
  				res = r1 & r2 /* compute value */
  			}

        handle(regs.Write(dr, res)) /* write res to register */
  			handle(updateFlags(regs, res)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_LDR instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        dr := (param >> 9) & 0x7
  			sr1 := (param >> 6) & 0x7
  			offset := signExtend(param & 0x3F, 6)

        r1 := readReg(regs, sr1) /* read first register */
  			res := readMem(memory, r1 + offset) /* read res off of memory */

        handle(regs.Write(dr, res)) /* write res to register */
        handle(updateFlags(regs, res)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_STR instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        sr1 := (param >> 9) & 0x7
  			sr2 := (param >> 6) & 0x7
  			offset := signExtend(param & 0x3F, 6)

        r1 := readReg(regs, sr1) /* read first register */
        r2 := readReg(regs, sr2) /* read second register */

        handle(memory.Write(r2 + offset, r1)) /* write to memory */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_RTI instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        return false, errors.New("Unknown opcode") /* return an error */
      },
    },

    &Lc3Instruction { /* OP_NOT instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        dr := (param >> 9) & 0x7
  			sr1 := (param >> 6) & 0x7

  			res := ^readReg(regs, sr1) /* read bitwise not of sr1 register */

        handle(regs.Write(dr, res)) /* write res to register */
        handle(updateFlags(regs, res)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_LDI instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        dr := (param >> 9) & 0x7
  			pc_offset := signExtend(param & 0x1ff, 9)

        pc := readReg(regs, registers.R_PC) /* read pc register */
        res := readMem(memory, readMem(memory, pc + pc_offset)) /* read res off of memory */

        handle(regs.Write(dr, res)) /* write res to register */
        handle(updateFlags(regs, res)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_STI instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        sr1 := (param >> 9) & 0x7
  			pc_offset := signExtend(param & 0x1ff, 9)

        pc := readReg(regs, registers.R_PC) /* read pc register */
        r1 := readReg(regs, sr1) /* read first register */

        mem := readMem(memory, pc + pc_offset) /* read off of memory */
        handle(memory.Write(mem, r1)) /* write to memory */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_JMP instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        sr1 := (param >> 6) & 0x7
        r1 := readReg(regs, sr1) /* read first register */
  			handle(regs.Write(registers.R_PC, r1)) /* assign r1 to pc */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_RES instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        return false, errors.New("Unknown opcode") /* return an error */
      },
    },

    &Lc3Instruction { /* OP_LEA instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        dr := (param >> 9) & 0x7
  			pc_offset := signExtend(param & 0x1ff, 9)

  			pc := readReg(regs, registers.R_PC) /* read pc register */
        res := pc + pc_offset

        handle(regs.Write(dr, res)) /* write res to register */
        handle(updateFlags(regs, res)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction { /* OP_TRAP instruction */
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recoverAll()
        }()

        switch param & 0xFF {
        case TRAP_GETC:
          c := getChar() /* read pressed key */
          handle(regs.Write(registers.R_R0, c)) /* write it to r0 register */

        case TRAP_OUT:
          r := readReg(regs, registers.R_R0) /* read r0 register */
          fmt.Printf("%c", rune(r))

        case TRAP_PUTS:
          r := readReg(regs, registers.R_R0) /* read r0 register */
          for mem := uint16(1); mem != 0; {
            mem = readMem(memory, r) /* read value off of memory */
            fmt.Printf("%c", rune(mem))
            r++ /* increment r */
          }

        case TRAP_IN:
          fmt.Print("Enter a character: ")
          c := getChar() /* read pressed key */
          handle(regs.Write(registers.R_R0, c)) /* write it to r0 register */

        case TRAP_PUTSP:
          r := readReg(regs, registers.R_R0) /* read r0 register */
          for {
            mem := readMem(memory, r) /* read value off of memory */
            if mem == 0 { /* if 0 exit */
              return true, err
            }
            r1 := rune(mem & 0xFF)
            fmt.Printf("%c", r1)
            r2 := rune(mem >> 8)
            if r2 != 0 {
              fmt.Printf("%c", r2)
            }
            r++ /* increment r */
          }

        case TRAP_HALT:
          fmt.Println("____halting cpu____")
          return false, err
        }
        return true, err
      },
    },
  },
}
