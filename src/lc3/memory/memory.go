package memory

import (
  "math"
)

const (
  UINT16_MAX = math.MaxUint16   /* maximum adress value for a uint16 */
)

/* defining the type of lc3 memory */
type lc3type uint16

/* and creating an array representing this memory */
var lc3mem [UINT16_MAX]lc3type
