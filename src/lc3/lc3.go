package lc3

import (
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
  "github.com/jolatechno/go-lc3/src/lc3/memory"
  "github.com/jolatechno/go-lc3/src/lc3/instructions"
)

var (
  Offset uint16 = 0x0000
)

/* defining the lc3 cpu interface */
type Lc3Cpu struct {
  Mem memory.Lc3Mem
  Instrs instructions.Lc3InstructionSet
  Regs registers.Lc3Registers
}

/* and define the variable */
var Lc3cpu = Lc3Cpu{ memory.Lc3mem, instructions.Lc3instructionSet, registers.Lc3registers }

/* defining the interface */
func (cpu *Lc3Cpu)Memory() (memory interfaces.Memory) {
  return &cpu.Mem
}

func (cpu *Lc3Cpu)InstructionSet() (memory interfaces.InstructionSet) {
  return &cpu.Instrs
}

func (cpu *Lc3Cpu)Registers() (memory interfaces.Registers) {
  return &cpu.Regs
}
