package lc3

import (
  "github.com/jolatechno/go-lc3/src/cpu"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
  "github.com/jolatechno/go-lc3/src/lc3/memory"
  "github.com/jolatechno/go-lc3/src/lc3/instructions"
)

var (
  PC_START = 0x3000 /* init the pc to a certain value */
  Origin = 0x0000 /* origin of program load */
)

func LoadPrgm(prgm []uint16) (n int, err error) {
  interfacePrgm := make([]interface{}, len(prgm)) /* create an array of interfaces */
  for i, v := range prgm { /* and populate it with the program */
    interfacePrgm[i] = v
  }

  return memory.Lc3mem.Writea(Origin, interfacePrgm) /* write the program to memory, starting at the origin */
}

func Exec() (err error) {
  (&registers.Lc3registers).Write(registers.R_PC, Origin) /* init pc register */
  return cpu.Run(&memory.Lc3mem, &registers.Lc3registers, &instructions.Lc3instructionSet) /* run the vm */
}
