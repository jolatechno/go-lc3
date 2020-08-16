package opcode

import (
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

var (
  test_instruction uint16 = 1
  test_param uint16 = test_instruction << 12
  test_opcode = &Lc3OP{ test_param }
)

func TestLc3Op(t *testing.T) {
  test_op, ok := interface{}(test_opcode).(interfaces.Op) /* testing if Lc3Op implements the write interface */
  if !ok { /* throwing an error */
    t.Error("couldn't convert lc3 opcode to interface")
  }

  instruction := test_op.Instruction() /* check the Instruction function */
  if instruction != test_instruction { /* throwing an error */
    t.Error("lc3 opcode instruction doesn't match")
  }

  param := test_op.Params() /* check the Param function */
  if param != test_param { /* throwing an error */
    t.Error("lc3 opcode param doesn't match")
  }
}
