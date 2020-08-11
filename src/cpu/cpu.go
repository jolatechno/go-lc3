package cpu

import (
  "github.com/jolatechno/go-lc3/src/interfaces"
)

/* function to run the vm */
func Run(memory interfaces.Memory, registers interfaces.Registers, instructions interfaces.InstructionSet) (err error) {
  for { /* main loop */
    op, err := registers.Fetch(memory) /* fetching the next instruction */
    if err != nil { /* throw error */
      return err
    }

    inst, err := instructions.Get(op) /* get the instruction opcode */
    if err != nil { /* throw error */
      return err
    }

    next, err := inst.Run(memory, registers, op) /* execute the instruction */
    if err != nil {
      return err
    }

    if !next { /* end the loop if the program finished */
      return nil
    }
  }
}
