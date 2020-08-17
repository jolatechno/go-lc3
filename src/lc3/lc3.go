package lc3

import (
  "errors"
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

/* function to convert two bytes to uint16 */
func byte2uint16(b1 byte, b2 byte) uint16 {
    _ = b1 // bounds check hint to compiler; see golang.org/issue/14808
    return uint16(b1) | uint16(b2) << 8
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
  if len(img) % 2 == 1 { /* check if the array is of even length */
    return 0, errors.New("can't convert byte array of uneven length to uint16 array")
  }

  prgm := make([]uint16, int(len(img) / 2)) /* create an empty array */
  for i, _ := range prgm { /* iterate over the array */
    prgm[i] = byte2uint16(img[2*i], img[2*i + 1]) /* convert two byte to uint */
  }

  n, err = cpu.Memory().Writea(uint16(0), prgm) /* write program to memory */
  return n*2, err /* return length in byte and error */
}
