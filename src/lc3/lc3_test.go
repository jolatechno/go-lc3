package lc3

import (
  "fmt"
  "testing"
  //"encoding/binary"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/cpu"
  "github.com/jolatechno/go-lc3/src/lc3/instructions"
  //"github.com/jolatechno/go-lc3/src/lc3/registers"
)

var (
  //offset uint16 = registers.PC_START
  //halt uint16 = instructions.OP_TRAP << 12 | instructions.TRAP_HALT
  //prgm = make([]byte, 4)
)

func TestLc3Cpu(t *testing.T) {
  instructions.DebugLc3Instructions = true

  _, ok := interface{}(&Lc3cpu).(interfaces.Cpu) /* testing if Lc3 implements the right interface */
  if !ok { /* throwing an error */
    t.Error("couldn't convert lc3 cpu to interface")
  }

  /* load halt value in memory */
  //binary.BigEndian.PutUint16(prgm[0:], offset)
	//binary.BigEndian.PutUint16(prgm[2:], halt)

  //n, err := (&Lc3cpu).Load(prgm) /* test the load function */
  //if err != nil { /* throwing an error */
  //  t.Error(err)
  //} else if n != len(prgm) - 2 { /* check the written length */
  //  t.Error("couldn't write the whole program")
  //} else {
  //  fmt.Println("") /* jump a line */
  //  err = cpu.Run(&Lc3cpu) /* test the run function */
  //  if err != nil { /* throwing an error */
  //    t.Error(err)
  //  }
  //}

  _, err := cpu.LoadFile(&Lc3cpu, "hello.obj") /* test the loadFile function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else {
    fmt.Println("") /* jump a line */
    err = cpu.Run(&Lc3cpu) /* test the run function */
    fmt.Println("") /* jump a line */
    if err != nil { /* throwing an error */
      t.Error(err)
    }
  }
}
