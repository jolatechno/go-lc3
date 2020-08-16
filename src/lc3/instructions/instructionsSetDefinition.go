package instructions

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

/* instructionset definition */
var Lc3instructionSet = Lc3InstructionSet {
  [OP_COUNT]*Lc3Instruction{
    &Lc3Instruction{
      OP: OP_BR,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_ADD,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        var r0 uint16 = (param >> 9) & 0x7 /* destination register (DR) */
        var r1 uint16 = (param >> 6) & 0x7 /* first operand (SR1) */
        var imm_flag uint16 = (param >> 5) & 0x1 /* whether we are in immediate mode */
        var res uint16

        v1, err := regs.Read(r1) /* read first register */
        if err != nil { /* throw an error */
          return false, err
        }

        v1int, ok := v1.(uint16) /* convert it to uint16 */
        if !ok { /* throw an error */
          return false, errors.New("can't convert register to int")
        }

        if imm_flag == 1 { /* second param is a value */
            var imm5 uint16 = sign_extend(param & 0x1F, 5) /* read value */

            res = v1int + imm5 /* compute value */
        } else { /* second param is a register */
            var r2 uint16 = param & 0x7 /* read register */

            v2, err := regs.Read(r2) /* read second register */
            if err != nil { /* throw an error */
              return false, err
            }

            v2int, ok := v2.(uint16)  /* convert it to uint16 */
            if !ok { /* throw an error */
              return false, errors.New("can't convert register to int")
            }

            res = v1int + v2int /* compute value */
        }

        err = regs.Write(r0, res) /* write result to register */
        if err != nil { /* throw an error */
          return false, err
        }

        return true, update_flags(res, regs) /* update flags */
      },
    },

    &Lc3Instruction{
      OP: OP_LD,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_ST,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_JSR,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_AND,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_LDR,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_STR,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_RTI,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_NOT,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_LDI,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_STI,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_JMP,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_RES,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_LEA,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },

    &Lc3Instruction{
      OP: OP_TRAP,
      Exec: func(memory interfaces.Memory, regs interfaces.Registers, param uint16) (next bool, err error) {
        /* TODO */
        return false, errors.New("Not yet implemented")
      },
    },
  },
}
