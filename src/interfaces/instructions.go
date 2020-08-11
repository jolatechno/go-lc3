package interfaces

import (

)

/* interface to define an instruction */
type Instruction interface {
  Run(memory Memory, registers Registers, params []interface{}) (err error) /* execute an instruction given specific state and parameters */
}

/* interface to define an instruction */
type Op interface {
  Params() (params []interface{}) /* get params */
  Instruction() (op interface{}) /* get instruction */
  ToMemory() (value interface{}) /* convert opcode to memory alocateable value */
  FromMemory(value interface{}) (err error) /* convert memory value to opcode */
}
