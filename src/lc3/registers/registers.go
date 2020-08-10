package registers

import (

)

const
(
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

type lc3register uint16

var lc3registers [R_COUNT]*lc3register;
