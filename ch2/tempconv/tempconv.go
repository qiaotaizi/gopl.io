package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64;

const (
	AbsZeroC  Celsius = -273.15
	BoilingC  Celsius = 100
	FreezingC         = 0
)

//让Celsius和Fahrenheit类型分别实现Stringer接口
func (c Celsius) String() string{
	return fmt.Sprintf("%g°C",c)
}

func (f Fahrenheit) String() string{
	return fmt.Sprintf("%g°F",f)
}

func (k Kelvin) String() string{
	return fmt.Sprintf("%g°K",k)
}
