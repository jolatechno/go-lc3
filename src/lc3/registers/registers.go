package registers

import (

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
type lc3register uint16

/* and an array to store them, which will be populated later */
var lc3registers [R_COUNT]*lc3register;
