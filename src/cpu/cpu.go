package cpu

import (
  "io/ioutil"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

/* function to load a program */
func LoadFile(cpu interfaces.Cpu, filename string) (n int, err error) {
  prgm, err := ioutil.ReadFile(filename) /* read file */
  if err != nil { /* throw an error */
    return 0, err
  }

  return cpu.Load(prgm) /* write prgm to memory */
}

/* function to run the vm */
func Run(cpu interfaces.Cpu) (err error) {
  for { /* main loop */
    pc, err := cpu.Registers().Fetch() /* fetching the next instruction */
    if err != nil { /* throw error */
      return err
    }

    op, err := cpu.Memory().Read(pc) /* fetching the next instruction */
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
