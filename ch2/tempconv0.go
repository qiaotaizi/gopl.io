package main

import (
	"fmt"
	"gopl.io/ch2/tempconv"
)

func main() {
	//K->C
	var k tempconv.Kelvin=273.15
	fmt.Printf("%g°K = %g°C\n",k,tempconv.KToC(k))

	//K->F
	fmt.Printf("%g°K = %g°F\n",k,tempconv.KToF(k))

	var f tempconv.Fahrenheit=100
	//F->K
	fmt.Printf("%g°F = %g°K\n",f,tempconv.FToK(f))
	//F->C
	fmt.Printf("%g°F = %g°C\n",f,tempconv.FToC(f))

	var c tempconv.Celsius=tempconv.AbsZeroC;
	//C->K
	fmt.Printf("%g°C = %g°K\n",c,tempconv.CToK(c))
	//C->F
	fmt.Printf("%g°C = %g°F\n",c,tempconv.CToF(c))

}
