package lc3

import (
  "fmt"
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/cpu"
  "github.com/jolatechno/go-lc3/src/lc3/instructions"
)

func TestLc3Cpu(t *testing.T) {
  instructions.DebugLc3Instructions = true

  _, ok := interface{}(&Lc3cpu).(interfaces.Cpu) /* testing if Lc3 implements the right interface */
  if !ok { /* throwing an error */
    t.Error("couldn't convert lc3 cpu to interface")
  }

  _, err := cpu.LoadFile(&Lc3cpu, "../../hello.obj") /* test the loadFile function */
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
