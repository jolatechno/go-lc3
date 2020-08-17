package memory

import (
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

var (
  address uint16 = 0
  value uint16 = 12
  instruction uint16 = value >> 12
  addresses uint16 = 0
  values = []uint16{13, 14}
)

func TestLc3Memory(t *testing.T) {
  _, ok := interface{}(&Lc3mem).(interfaces.Memory) /* testing if Lc3mem implements the right interface */
  if !ok { /* throwing an error */
    t.Error("couldn't convert lc3 memory to interface")
  }


  err := (&Lc3mem).Write(address, value) /* checking the Write function */
  if err != nil { /* throwing an error */
    t.Error(err)
  }

  read_value, err := (&Lc3mem).Read(address) /* checking the Read function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if read_value != value { /* checking if read and write value match */
    t.Error("lc3 memory read and write values don't match")
  }

  read_op, err := (&Lc3mem).Fetch(address) /* checking the Read function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if read_op.Instruction() != instruction { /* checking if read and write instruction match */
    t.Error("lc3 memory fetch and write instruction don't match")
  } else if read_op.Params() != value { /* checking if read and write instruction match */
    t.Error("lc3 memory fetch and write params don't match")
  }

  n, err := (&Lc3mem).Writea(addresses, values) /* checking the Write function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if n != len(values) { /* checking if writea value is of right lentgh */
    t.Error("lc3 memory writea didn't write the whole array")
  }

  read_values := interface{}(make([]uint16, len(values)))
  n, err = (&Lc3mem).Reada(addresses, read_values) /* checking the Read function */
  if err != nil { /* throwing an error */
    t.Error(err)
  } else if n != len(values) { /* checking if reada value is of right lentgh */
    t.Error("lc3 memory writea didn't write the whole array")
  } else {
    for i, v := range read_values.([]uint16) {
      if values[i] != v { /* checking if reada and writea value match */
        t.Error("lc3 memory reada and writea values don't match")
      }
    }
  }
}
