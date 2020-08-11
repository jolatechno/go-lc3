package memory

import (
  "math"
  "errors"
)

const (
  UINT16_MAX = math.MaxUint16   /* maximum adress value for a uint16 */
)

/* defining the type of lc3 memory */
type Lc3Mem [UINT16_MAX]uint16

/* and creating an array representing this memory */
var Lc3mem Lc3Mem

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

  Lc3mem[intAddr] = intValue
  return nil
}

func (mem *Lc3Mem)Writea(address interface{}, values []interface{}) (n int, err error) {
  n = len(values) /* maximum size written */
  err = nil

  intAddr, ok := address.(uint16) /* convert adress to int */
  if !ok { /* return an error if not possible */
    return 0, errors.New("adress not understood")
  }

  if intAddr < 0 || intAddr >= UINT16_MAX { /* check if adress is not out of range */
    return 0, errors.New("adress is out of range")
  }

  if int(intAddr) + n > UINT16_MAX { /* check if maximum adress is not out of range */
    n = int(UINT16_MAX - intAddr)
    err = errors.New("Write out of range")
  }

  for i, value := range values[0: n] { /* write values to memory */
    intValue, ok := value.(uint16) /* convert value to int */
    if !ok {  /* return an error if not possible */
      return i, errors.New("value not understood")
    }
    Lc3mem[intAddr + uint16(i)] = intValue
  }

  return n, err
}

func (mem *Lc3Mem)Read(address interface{}) (value interface{}, err error) {
  intAddr, ok := address.(uint16) /* convert adress to int */
  if !ok { /* return an error if not possible */
    return 0, errors.New("adress not understood")
  }

  if intAddr < 0 || intAddr >= UINT16_MAX { /* check if adress is not out of range */
    return 0, errors.New("adress is out of range")
  }

  value = Lc3mem[intAddr] /* read value off of memory */
  return value, nil
}

func (mem *Lc3Mem)Reada(address interface{}, values []interface{}) (n int, err error) {
  n = len(values) /* maximum size written */
  err = nil

  intAddr, ok := address.(uint16) /* convert adress to int */
  if !ok { /* return an error if not possible */
    return 0, errors.New("adress not understood")
  }

  if intAddr < 0 || intAddr >= UINT16_MAX { /* check if adress is not out of range */
    return 0, errors.New("adress is out of range")
  }

  if int(intAddr) + n > UINT16_MAX { /* check if maximum adress is not out of range */
    n = int(UINT16_MAX - intAddr)
    err = errors.New("Write out of range")
  }

  for i, value := range (*mem)[intAddr: intAddr + uint16(n)] { /* read values off of memory */
    values[i] = value
  }

  return n, err
}
