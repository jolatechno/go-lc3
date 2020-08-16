package cpu

import (
  "github.com/jolatechno/go-lc3/src/interfaces"
)

/* function to load a program */
func Load(cpu interfaces.Cpu, offset interface{}, prgm interface{}) (n int, err error) {
  return cpu.Memory().Writea(offset, prgm) /* write prgm to memory at offset */
}

/* function to run the vm */
func Run(cpu interfaces.Cpu) (err error) {
  for { /* main loop */
    pc, err := cpu.Registers().Fetch() /* fetching the next instruction */
    if err != nil { /* throw error */
      return err
    }

    op, err := cpu.Memory().Fetch(pc) /* fetching the next instruction */
    if err != nil { /* throw error */
      return err
    }

    inst, err := cpu.InstructionSet().Get(op) /* get the instruction opcode */
    if err != nil { /* throw error */
      return err
    }

    next, err := inst.Run(cpu.Memory(), cpu.Registers(), op) /* execute the instruction */
    if err != nil {
      return err
    }

    if !next { /* end the loop if the program finished */
      return nil
    }
  }
}
