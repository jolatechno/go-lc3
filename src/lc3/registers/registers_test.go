package registers

import (
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/memory"
)

var (
  reg uint16 = 2
  value uint16 = 8
  test_memory = &memory.Lc3mem
  instruction uint16 = 2
  mem_value uint16 = instruction << 12
)

func TestLc3Registers(t *testing.T) {
  _, ok := interface{}(&Lc3registers).(interfaces.Registers) /* testing if Lc3registers implements the write interface */
  if !ok { /* throwing an error */
    t.Error("couldn't convert lc3 registers to interface")
  }

  err := (&Lc3registers).Write(reg, value) /* check the Write function */
  if err != nil { /* throwing an error */
    t.Error(err)
  }

  read_reg, err := (&Lc3registers).Read(reg) /* check the Param function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if read_reg != value { /* check if read and write match */
    t.Error("lc3 register read and write don't match")
  }

  pc, err := (&Lc3registers).Read(R_PC) /* reading the pc */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else {
    err = test_memory.Write(pc, mem_value) /* writting instruction to memory */
    if err != nil { /* throwing an error */
      t.Error("couldn't test fetch because of memory write error")
    } else {
      op, err := (&Lc3registers).Fetch(test_memory) /* checking the fetch function */
      if err != nil { /* throwing an error */
        t.Error(err)
      } else {
        if op.Instruction() != instruction { /* check if param and write match */
          t.Error("lc3 register fetch instruction don't match")
        }

        if op.Params() != mem_value { /* check if param and write match */
          t.Error("lc3 register fetch params don't match")
        }
      }
    }
  }

}
