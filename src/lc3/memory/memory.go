package memory

import (
  "math"
)

const (
  UINT16_MAX = math.MaxUint16
)

type lc3type uint16
var lc3mem [UINT16_MAX]lc3type
