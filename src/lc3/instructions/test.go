package instructions

import (
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/opcode"
)

var (
  test_instruction_set interfaces.InstructionSet
  test_opcode = &opcode.Lc3OP{ 0 }
)

func TestLc3InstructionSet(t *testing.T) {
  test_instruction_set, ok := interface{}(&Lc3instructionSet).(interfaces.InstructionSet)
  if !ok {
    t.Error("couldn't convert lc3 instructionSet to interface")
  }

  _, err := test_instruction_set.Get(test_opcode)
  if err != nil {
    t.Error(err)
  }
}
