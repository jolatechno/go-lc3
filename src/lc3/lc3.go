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
  Instrs instructions.Lc3InstructionSet
  Mem memory.Lc3Mem
  Regs registers.Lc3Registers
}

func New() (cpu interfaces.Cpu) { /* function to create a new lc3 cpu */
  return &Lc3Cpu{
    instructions.Lc3instructionSet,
    *memory.New().(*memory.Lc3Mem),
    *registers.New().(*registers.Lc3Registers),
  }
}

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

  for i := 2; i < len(img); i += 2 {
    value := binary.BigEndian.Uint16(img[i:i + 2]) /* convert two byte to uint */
    err = cpu.Memory().Write(origin, value) /* load value to memory */
    if err != nil { /* throwing an error */
      return i - 2, err
    }

    origin++
  }

  return len(img) - 2, err
}
