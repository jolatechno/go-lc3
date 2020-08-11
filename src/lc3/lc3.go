package lc3

import (
  "github.com/jolatechno/go-lc3/src/cpu"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
  "github.com/jolatechno/go-lc3/src/lc3/memory"
  "github.com/jolatechno/go-lc3/src/lc3/instructions"
)

func Exec(prgm []uint16) (err error){
  pc_start := 0x3000 /* init the pc to a certain value */

  interfacePrgm := make([]interface{}, len(prgm)) /* create an array of interfaces */
  for i, v := range prgm { /* and populate it with the program */
    interfacePrgm[i] = v
  }

  memory.Lc3mem.Writea(pc_start, interfacePrgm) /* write the program to memory, starting at the pc */
  return cpu.Run(&memory.Lc3mem, &registers.Lc3registers, &instructions.Lc3instructionSet) /* run the vm */
}
