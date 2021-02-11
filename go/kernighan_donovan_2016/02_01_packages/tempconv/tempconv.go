// package of tempratures conversions
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
    AbsoluteZeroC Celsius = -273.15
    FeezingC Celsius = 0.00
    BoilingC Celsius = 100.00
)

func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
    return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
    return fmt.Sprintf("%g°K", k)
}

// CtoF converts Celsius temprature to Fahrenheit
func CtoF (c Celsius) Fahrenheit {
    return Fahrenheit(c * 9 / 5 + 32)
}

// CtoK converts Celsius temprature to Kelvin
func CtoK (c Celsius) Kelvin {
    return Kelvin(c - AbsoluteZeroC)
}

// FtoC converts Fahrenheit temprature to Celsius
func FtoC (f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}

// FtoK converts Fahrenheit temprature to Kelvin
func FtoK (f Fahrenheit) Kelvin {
    return Kelvin((f - 32) * (5 / 9) - AbsoluteZeroC)
}

// KtoC converts Kelvin temprature to Celsius
func KtoC (k Kelvin) Celsius {
    return Celsius(k + AbsoluteZeroC)
}

// KtoF converts Kelvin temprature to Fahrenheit
func KtoF (k Kelvin) Fahrenheit {
    return Fahrenheit(k * 9 / 5 + 32 - AbsoluteZeroC)
}
