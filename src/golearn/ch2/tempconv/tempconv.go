package tempconv

import "fmt"

type Celsius float64
type Fahreneit float64

const (
	AbsoluteSeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahreneit) String() string {
	return fmt.Sprintf("%g°F", f)
}
