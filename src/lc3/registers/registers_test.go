package registers

import (
  "testing"

)

var (
  reg uint16 = 2
  value uint16 = 8
  instruction uint16 = 2
  mem_value uint16 = instruction << 12
)

func TestLc3Registers(t *testing.T) {
  lc3registers := New() /* create a new registers set */

  err := lc3registers.Write(reg, value) /* check the Write function */
  if err != nil { /* throwing an error */
    t.Error(err)
  }

  read_reg, err := lc3registers.Read(reg) /* check the Param function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if read_reg != value { /* check if read and write match */
    t.Error("lc3 register read and write don't match")
  }

  pc, err := lc3registers.Read(R_PC) /* reading the pc */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else {
    pc_fetch, err := lc3registers.Fetch() /* Fetch the pc */
    if err != nil { /* throwing an error */
      t.Error(err)
    } else if pc != pc_fetch { /* check if pc match */
      t.Error("lc3 pc don't match")
    }
  }

  lc3registers.Reset() /* reset registers */

  zero_reg, err := lc3registers.Read(reg) /* read register */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if zero_reg != uint16(0) {  /* check if registers have been reset */
    t.Error("lc3 registers haven't been reset")
  }
}
