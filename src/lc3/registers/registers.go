package registers

import (
  "errors"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/opcode"
)

/* defining the name of all lc3 registers */
const (
    R_R0 = iota
    R_R1 = iota
    R_R2 = iota
    R_R3 = iota
    R_R4 = iota
    R_R5 = iota
    R_R6 = iota
    R_R7 = iota
    R_PC = iota  /* program counter */
    R_COND = iota
    R_COUNT = iota
)

/* defining condition flags */
const (
    FL_POS = 1 << iota /* Positive */
    FL_ZRO = 1 << iota /* Zero */
    FL_NEG = 1 << iota /* Negative */
)

/* defining the type of those register */
type Lc3Registers [R_COUNT]uint16

/* and an array to store them, which will be populated later */
var Lc3registers Lc3Registers;

/* defining the interface */
func (regs *Lc3Registers)Write(reg interface{}, value interface{}) (err error) {
  intReg, ok := reg.(uint16) /* convert reg to int */
  if !ok { /* return an error if not possible */
    return errors.New("register not understood")
  }

  intValue, ok := value.(uint16) /* convert value to int */
  if !ok { /* return an error if not possible */
    return errors.New("value not understood")
  }

  if intReg < 0 || intReg >= R_COUNT { /* check if reg is not out of range */
    return errors.New("register is out of range")
  }

  Lc3registers[intReg] = intValue /* write value into the register array */
  return nil
}

func (regs *Lc3Registers)Read(reg interface{}) (value interface{}, err error) {
  intReg, ok := reg.(uint16) /* convert reg to int */
  if !ok { /* return an error if not possible */
    return 0, errors.New("register not understood")
  }

  if intReg < 0 || intReg >= R_COUNT { /* check if reg is not out of range */
    return 0, errors.New("register is out of range")
  }

  value = Lc3registers[intReg] /* read value from memory */
  return value, nil
}

func (regs *Lc3Registers)Fetch(memory interfaces.Memory) (op interfaces.Op, err error) {
  pc, err := regs.Read(R_PC) /* get pc value */
  if err != nil { /* throw error */
    return nil, err
  }

  inst, err := memory.Read(pc) /* read the value in memory corresponding to pc */
  if err != nil { /* throw error */
    return nil, err
  }

  intInst, ok := inst.(uint16) /* convert reg to int */
  if !ok { /* return an error if not possible */
    return nil, errors.New("instruction not understood")
  }

  return &opcode.Lc3OP{ intInst }, nil /* convert this value to an opcode */
}
