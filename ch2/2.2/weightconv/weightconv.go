package weightconv

import (
	"fmt"
)

type Kilograms float64
type Pounds float64

func (k Kilograms) String() string { return fmt.Sprintf("%.6gkg", k) }

func (p Pounds) String() string { return fmt.Sprintf("%.6glbs", p) }
