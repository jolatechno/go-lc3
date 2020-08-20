package memory

import (
  "testing"
)

var (
  address uint16 = 0
  value uint16 = 12
  instruction uint16 = value >> 12
  addresses uint16 = 0
  values = []uint16{13, 14}
)

func TestLc3Memory(t *testing.T) {
  lc3mem := New() /* create a new memory interface */

  err := lc3mem.Write(address, value) /* checking the Write function */
  if err != nil { /* throwing an error */
    t.Error(err)
  }

  read_value, err := lc3mem.Read(address) /* checking the Read function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if read_value != value { /* checking if read and write value match */
    t.Error("lc3 memory read and write values don't match")
  }

  lc3mem.Reset() /* reset memory */

  zero_value, err := lc3mem.Read(address) /* checking the Read function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if zero_value != uint16(0) { /* checking if read value has been reset */
    t.Error("lc3 memory hasn't been reset")
  }
}
