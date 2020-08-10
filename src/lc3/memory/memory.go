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
  /* TODO */
  return errors.New("Not implemented")
}

func (mem *Lc3Mem)Writea(address interface{}, values []interface{}) (n int, err error) {
  /* TODO */
  return 0, errors.New("Not implemented")
}

func (mem *Lc3Mem)Read(address interface{}) (value interface{}, err error) {
  /* TODO */
  return nil, errors.New("Not implemented")
}

func (mem *Lc3Mem)Reada(address interface{}, values []interface{}) (n int, err error) {
  /* TODO */
  return 0, errors.New("Not implemented")
}
