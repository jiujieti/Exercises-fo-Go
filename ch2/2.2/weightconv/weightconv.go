package weightconv

import (
	"fmt"
)

type Kilograms float64
type Pounds float64

func (k Kilograms) String() string { return fmt.Sprintf("%gkg", k) }

func (p Pounds) String() string { return fmt.Sprintf("%glbs", p) }
