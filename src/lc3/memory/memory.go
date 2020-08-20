package memory

import (
  "math"
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

const (
  UINT16_MAX = math.MaxUint16   /* maximum adress value for a uint16 */
)

/* defining the type of lc3 memory */
type Lc3Mem struct{
  Buff [UINT16_MAX]uint16
}

func New() (mem interfaces.Memory) { /* function to create a new lc3mem interface */
  var lc3mem [UINT16_MAX]uint16
  return &Lc3Mem{ lc3mem }
}

/* defining the interface */
func (mem *Lc3Mem)Write(address interface{}, value interface{}) (err error) {
  intAddr, ok := address.(uint16) /* convert adress to int */
  if !ok { /* return an error if not possible */
    return errors.New("adress not understood")
  }

  if intAddr < 0 || intAddr >= UINT16_MAX { /* check if adress is not out of range */
    return errors.New("adress is out of range")
  }

  intValue, ok := value.(uint16) /* convert value to int */
  if !ok { /* return an error if not possible */
    return errors.New("value not understood")
  }

  mem.Buff[intAddr] = intValue
  return nil
}

func (mem *Lc3Mem)Read(address interface{}) (value interface{}, err error) {
  intAddr, ok := address.(uint16) /* convert adress to int */
  if !ok { /* return an error if not possible */
    return 0, errors.New("adress not understood")
  }

  if intAddr < 0 || intAddr >= UINT16_MAX { /* check if adress is not out of range */
    return 0, errors.New("adress is out of range")
  }

  value = mem.Buff[intAddr] /* read value off of memory */
  return value, nil
}

func (mem *Lc3Mem)Reset() {
  (*mem) = *(New().(*Lc3Mem))
}
