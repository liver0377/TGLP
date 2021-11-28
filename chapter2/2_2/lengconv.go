package lengconv

import (
	"fmt"
)

type Foot float64
type Meter float64

func (f Foot) Sting() string { return fmt.Sprintf("%g", f)}
func (m Meter) String() string { return fmt.Sprintf("%g", m)}

