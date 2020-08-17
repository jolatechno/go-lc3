package interfaces

import (
)

type Memory interface {
  Write(address interface{}, value interface{}) (err error) /* write a single value to memory */
  Writea(address interface{}, values interface{}) (n int, err error) /* write an array of values to memory */
  Read(address interface{}) (value interface{}, err error) /* read a single value to memory */
  Reada(address interface{}, values interface{}) (n int, err error) /* write an array of values to memory */
  Fetch(pc interface{}) (op Op, err error) /* fetch an instruction off of memory */
}
