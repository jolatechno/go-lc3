package instructions

import (
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

var (
  test_opcode uint16 = 0 << 12
)

func TestLc3InstructionSet(t *testing.T) {
  _, ok := interface{}(&Lc3instructionSet).(interfaces.InstructionSet) /* checking if Lc3instructionSet implement the right interface*/
  if !ok { /* throwing an error */
    t.Error("couldn't convert lc3 instructionSet to interface")
  }

  _, err := (&Lc3instructionSet).Get(test_opcode) /* testing the Get method */
  if err != nil { /* throwing an error */
    t.Error(err)
  }
}
