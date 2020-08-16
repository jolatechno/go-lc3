package memory

import (
  "testing"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

func TestLc3Memory(t *testing.T) {
  test_memory, ok := interface{}(&Lc3mem).(interfaces.Memory) /* testing if Lc3mem implements the write interface */
  if !ok { /* throwing an error */
    t.Error("couldn't convert lc3 memory to interface")
  }

  address, value := 0, 12
  err := test_memory.Write(address, value) /* checking the Write function */
  if err != nil { /* throwing an error */
    t.Error(err)
  }

  read_value, err := test_memory.Read(address) /* checking the Read function */
  if err != nil { /* throwing an error */
    t.Error(err)
  }
  if read_value != value { /* checking if read and write value match */
    t.Error("lc3 memory read and write values don't match")
  }

  addresses, values := []uint16{0, 1}, []uint16{13, 14}
  n, err := test_memory.Writea(addresses, values) /* checking the Write function */
  if err != nil { /* throwing an error */
    t.Error(err)
  }
  if n != len(values) { /* checking if writea value is of right lentgh */
    t.Error("lc3 memory writea didn't write the whole array")
  }

  read_values := make([]uint16, len(values))
  n, err = test_memory.Reada(addresses, read_values) /* checking the Read function */
  if err != nil { /* throwing an error */
    t.Error(err)
  }
  if n != len(values) { /* checking if reada value is of right lentgh */
    t.Error("lc3 memory writea didn't write the whole array")
  }
  for i, v := range values {
    if read_values[i] != v { /* checking if reada and writea value match */
      t.Error("lc3 memory reada and writea values don't match")
    }
  }
}
