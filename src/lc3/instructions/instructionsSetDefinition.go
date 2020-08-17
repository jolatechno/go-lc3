package instructions

import (
  "errors"
  "fmt"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
)

/* instructionset definition */
var Lc3instructionSet = Lc3InstructionSet {
  [OP_COUNT]*Lc3Instruction{
    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        pc_offset := sign_extend(param & 0x1ff, 9)
  			flag := (param >> 9) & 0x7

        cond := readReg(regs, registers.R_COND) /* read cond register */
  			if (flag & cond) != 0 {
          pc := readReg(regs, registers.R_PC) /* read pc register */
  				handle(regs.Write(registers.R_COND, pc + pc_offset)) /* write pc to rtegister */
  			}

        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        dr := (param >> 9) & 0x7
  			sr1 := (param >> 6) & 0x7
  			mode := (param >> 5) & 0x1
        var res uint16

        r1 := readReg(regs, sr1) /* read first register */
        if mode == 1 { /* second param is a value */
            immediate := sign_extend(param & 0x1F, 5) /* read value */
            res = r1 + immediate /* compute value */
        } else { /* second param is a register */
            sr2 := param & 0x7 /* read register */
            r2 := readReg(regs, sr2) /* read second register */
            res = r1 + r2 /* compute value */
        }

        handle(regs.Write(dr, res)) /* write result to register */
        handle(update_flags(res, regs)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        dr := (param >> 9) & 0x7
  			pc_offset := sign_extend(param&0x1ff, 9)

        pc := readReg(regs, registers.R_PC) /* read pc register */
  			res := readMem(memory, pc + pc_offset) /* read value off of memory */

        handle(regs.Write(dr, res)) /* write res to register */
  			handle(update_flags(res, regs)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        sr := (param >> 9) & 0x7
  			pc_offset := sign_extend(param & 0x1ff, 9)

        pc := readReg(regs, registers.R_PC) /* read pc register */
        s := readReg(regs, sr) /* read sr register */

        handle(memory.Write(pc + pc_offset, s)) /* write to memory */

        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        r := (param >> 6) & 0x7
  			long_pc_Offset := sign_extend(param & 0x7ff, 11)
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

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        dr := (param >> 9) & 0x7
  			sr1 := (param >> 6) & 0x7
  			mode := (param >> 5) & 0x1
        var res uint16

        r1 := readReg(regs, sr1) /* read first register */
  			if mode == 1 {
  				imm5 := sign_extend(param & 0x1F, 5)
  				res =  r1 & imm5 /* compute value */
  			} else {
  				sr2 := param & 0x7
          r2 := readReg(regs, sr2) /* read second register */
  				res = r1 & r2 /* compute value */
  			}

        handle(regs.Write(dr, res)) /* write res to register */
  			handle(update_flags(res, regs)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        return false, errors.New("Unknown opcode") /* return an error */
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        dr := (param >> 9) & 0x7
  			sr1 := (param >> 6) & 0x7

  			res := ^readReg(regs, sr1) /* read bitwise not of sr1 register */

        handle(regs.Write(dr, res)) /* write res to register */
        handle(update_flags(res, regs)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        sr1 := (param >> 6) & 0x7
        r1 := readReg(regs, sr1) /* read first register */
  			handle(regs.Write(registers.R_PC, r1)) /* assign r1 to pc */

        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        return false, errors.New("Unknown opcode") /* return an error */
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        dr := (param >> 9) & 0x7
  			pc_offset := sign_extend(param & 0x1ff, 9)

  			pc := readReg(regs, registers.R_PC) /* read pc register */
        res := pc + pc_offset

        handle(regs.Write(dr, res)) /* write res to register */
        handle(update_flags(res, regs)) /* update flags */

        return true, err
      },
    },

    &Lc3Instruction{
      func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        switch param & 0xFF {
        case TRAP_GETC:
          c := GetChar() /* read pressed key */
          handle(regs.Write(registers.R_R0, c)) /* write it to r0 register */

        case TRAP_OUT:
          r := readReg(regs, registers.R_R0) /* read r0 register */
          fmt.Printf("%c", rune(r))

        case TRAP_PUTS:
          r := readReg(regs, registers.R_R0) /* read r0 register */
          for {
            mem := readMem(memory, r) /* read value off of memory */
            if mem == 0 { /* if 0 exit */
              return true, err
            }
            fmt.Printf("%c", rune(mem))
            r++ /* increment r */
          }

        case TRAP_IN:
          fmt.Print("Enter a character: ")
          c := GetChar() /* read pressed key */
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
          return false, err
        }
        return true, err
      },
    },
  },
}
