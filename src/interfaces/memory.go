package interfaces

import (
)

type Memory interface {
  Write(address interface{}, value interface{}) (err error) /* write a single value to memory */
  Read(address interface{}) (value interface{}, err error) /* read a single value to memory */
  Reset() /* reset the memory */
}
