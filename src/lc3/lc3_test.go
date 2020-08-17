package lc3

import (
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/cpu"
  "github.com/jolatechno/go-lc3/src/lc3/instructions"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
)

var (
  offset uint16 = registers.PC_START
  halt uint16 = instructions.OP_TRAP << 12 | instructions.TRAP_HALT
  prgm = make([]byte, 2*registers.PC_START + 2)
)

func TestLc3Cpu(t *testing.T) {
  _, ok := interface{}(&Lc3cpu).(interfaces.Cpu) /* testing if Lc3 implements the right interface */
  if !ok { /* throwing an error */
    t.Error("couldn't convert lc3 cpu to interface")
  }

  /* load halt value in memory */
  prgm[2*registers.PC_START] = byte(halt & 0xFF)
  prgm[2*registers.PC_START + 1] = byte(halt >> 8)

  n, err := (&Lc3cpu).Load(prgm) /* test the load function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if n != len(prgm) { /* check the written length */
    t.Error("couldn't write the whole program")
  } else {
    err = cpu.Run(&Lc3cpu) /* test the run function */
    if err != nil { /* throwing an error */
      t.Error(err)
    }
  }
}
