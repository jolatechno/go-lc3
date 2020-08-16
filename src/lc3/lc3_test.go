package lc3

import (
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/cpu"
)

var (
  prgm []uint16 = []uint16{ 0 }
)

func TestLc3Cpu(t *testing.T) {
  _, ok := interface{}(&Lc3cpu).(interfaces.Cpu) /* testing if Lc3 implements the right interface */
  if !ok { /* throwing an error */
    t.Error("couldn't convert lc3 cpu to interface")
  }

  n, err := cpu.Load(&Lc3cpu, Offset, prgm) /* test the load function */
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
