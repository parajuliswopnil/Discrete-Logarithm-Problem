package main

import (
	"fmt"

	"github.com/parajuliswopnil/Discrete-Logarithm-Problem/sbgm"
)

var (
	generator int64 = 13 // generator of the field Z50
	beta      int64 = 31 // generator ^ x = beta : for example 13^x = 47, the discrete logarithm problem in this case is to find the value of x 
)

func main() {
	sbgm.CreateFieldZ50(generator) 
	m := sbgm.CalculateM() 
	invGenerator := sbgm.InverseOfGenerator(generator)
	genPowMinusM := sbgm.ComputeGeneratorPowMinusM(invGenerator, m)
	sbgm.CreateXbTable(generator, m)

	gamma := beta

	var discreteLog int64

	for xg := 0; xg < int(m)-1; xg++ {
		fmt.Println(gamma)
		xb, ok := sbgm.XbTable[gamma]
		if ok {
			discreteLog = int64(xg)*m + xb
			break
		} else {
			gamma = (gamma * genPowMinusM) % 50
		}
	}
	fmt.Println(discreteLog)
}
