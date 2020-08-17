package instructions

import (
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/opcode"
)

var (
  test_opcode = &opcode.Lc3OP{ 0 << 12 }
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
