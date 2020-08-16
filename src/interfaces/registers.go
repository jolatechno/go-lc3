package interfaces

import (
)

type Registers interface {
  Write(reg interface{}, value interface{}) (err error) /* read the value of a register */
  Read(reg interface{}) (value interface{}, err error) /* read the value of a register */
  Fetch() (value interface{}, err error) /* fetch the pc off of memory */
}
