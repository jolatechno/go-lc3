package interfaces

import (

)

type Registers interface {
  Write(reg interface{}, value interface{}) (err error) /* read the value of a register */
  Read(reg interface{}) (value interface{}, err error) /* read the value of a register */
  Fetch(memory Memory) (instruction interface{}, err error) /* fetch an instruction off of memory */
}
