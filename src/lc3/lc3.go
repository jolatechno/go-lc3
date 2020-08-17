package lc3

import (
  "encoding/binary"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
  "github.com/jolatechno/go-lc3/src/lc3/memory"
  "github.com/jolatechno/go-lc3/src/lc3/instructions"
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

func (cpu *Lc3Cpu)Load(img []byte) (n int, err error) {
  origin := binary.BigEndian.Uint16(img[0:2]) /* read origin off of the image */

  prgm := make([]uint16, int(len(img) / 2 - 1)) /* create an empty array */
  for i, j := 2, 0; i < len(img); i += 2 {
    prgm[j] = binary.BigEndian.Uint16(img[i:i + 2]) /* convert two byte to uint */
    j++
  }

  n, err = cpu.Memory().Writea(origin, prgm) /* write program to memory */
  return n*2, err /* return length in byte and error */
}
