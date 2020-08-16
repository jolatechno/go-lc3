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
    t.Error("lc3 memory read and write calue don't match")
  }
}
