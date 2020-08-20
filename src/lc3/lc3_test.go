package lc3

import (
  "fmt"
  "testing"
  "github.com/jolatechno/go-lc3/src/cpu"
  "github.com/jolatechno/go-lc3/src/lc3/instructions"
)

func TestLc3Cpu(t *testing.T) {
  instructions.DebugLc3Instructions = true

  lc3cpu := New() /* define a new cpu interface */

  _, err := cpu.LoadFile(lc3cpu, "../../hello.obj") /* test the loadFile function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else {
    fmt.Println("") /* jump a line */
    err = cpu.Run(lc3cpu) /* test the run function */
    fmt.Println("") /* jump a line */
    if err != nil { /* throwing an error */
      t.Error(err)
    }
  }
}
