package instructions

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
)

/* function to read register and convert it to uint 16 */
func readReg(regs interfaces.Registers, r uint16) (value uint16) {
  v, err := regs.Read(r) /* read register */
  if err != nil { /* throw an error */
    panic(err)
  }

  value, ok := v.(uint16) /* convert it to uint16 */
  if !ok { /* throw an error */
    panic("can't convert register to int")
  }

  return value
}

/* function to handle panics */
func handle(err error) {
  if err != nil {
    panic(err)
  }
}

/* convert any recover to error for return */
func recover_all() (err error) {
  r := recover()
  if r != nil {
    switch x := r.(type) {
    case string:
      return errors.New(x)

    case error:
      return x

    default:
      return errors.New("Unknown panic")
    }
  }

  return nil
}

/* instructionset definition */
var Lc3instructionSet = Lc3InstructionSet {
  [OP_COUNT]*Lc3Instruction{
    &Lc3Instruction{
      OP: OP_BR,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        var pc_offset uint16 = sign_extend(param & 0x1ff, 9)
  			var flag uint16 = (param >> 9) & 0x7

        cond := readReg(regs, registers.R_COND) /* read cond register */
  			if (flag & cond) != 0 {
          pc := readReg(regs, registers.R_PC) /* read pc register */
  				handle(regs.Write(registers.R_COND, pc + pc_offset)) /* write pc to rtegister */
  			}

        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_ADD,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        var r0 uint16 = (param >> 9) & 0x7 /* destination register (DR) */
        var r1 uint16 = (param >> 6) & 0x7 /* first operand (SR1) */
        var imm_flag uint16 = (param >> 5) & 0x1 /* whether we are in immediate mode */
        var res uint16

        v1 := readReg(regs, r1) /* read first register */
        if imm_flag == 1 { /* second param is a value */
            var imm5 uint16 = sign_extend(param & 0x1F, 5) /* read value */

            res = v1 + imm5 /* compute value */
        } else { /* second param is a register */
            var r2 uint16 = param & 0x7 /* read register */
            v2 := readReg(regs, r2) /* read second register */
            res = v1 + v2 /* compute value */
        }

        handle(regs.Write(r0, res)) /* write result to register */
        handle(update_flags(res, regs)) /* update flags */
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_LD,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_ST,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_JSR,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_AND,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_LDR,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_STR,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_RTI,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_NOT,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_LDI,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_STI,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_JMP,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_RES,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_LEA,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },

    &Lc3Instruction{
      OP: OP_TRAP,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        defer func() { /* recover panics */
          err = recover_all()
        }()

        /* TODO */
        panic(errors.New("Not yet implemented"))
        return true, err
      },
    },
  },
}
