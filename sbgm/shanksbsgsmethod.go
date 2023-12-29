package sbgm



/*--------------------------------------------------------------------------------------------------------------------*/
/******             Solving the discrete logarithm problem using Shank's Baby Step Giant Step Method          *********/
/*--------------------------------------------------------------------------------------------------------------------*/


import (
	"fmt"
	"math"
	"math/big"
)

var (
	field   = []int64{}
	XbTable = map[int64]int64{}
)

// factors of 50 is 2 * 5 ^ 2, so prime factors are 2 and 5
// phi(a^n . b^m) = (a^n - a^(n-1))(b^n - b^(n-1))
func phiof50() int64 {
	return (2 - 1) * (25 - 5)
}

func CreateFieldZ50(generator int64) {
	n := big.NewInt(50) // for Z50, n = 50
	numberOfFieldElements := phiof50()

	for i := 0; i < int(numberOfFieldElements); i++ {
		powersOfGenerator := new(big.Int)
		element := powersOfGenerator.Exp(big.NewInt(generator), big.NewInt(int64(i)), n)
		field = append(field, element.Int64())
	}
	fmt.Println(field)
}

// m = ceil(sqrt(phi(n)))
func CalculateM() int64 {
	return int64(
		math.Ceil(
			math.Sqrt(
				float64(
					phiof50()))))
}

// calculating inverse for the generator such that generator * inv(generator) congurent to 1 mod n
// the calcuation here is not the standard way of calculating the inverse
func InverseOfGenerator(generator int64) int64 {
	var n int64 = 50 // for Z50, n = 50
	for _, v := range field {
		if (generator*v)%n == int64(1) {
			return v
		}
	}
	return 0
}

// compute g^-m
func ComputeGeneratorPowMinusM(invGenerator, m int64) int64 {
	var n int64 = 50 // for Z50, n = 50
	return int64(math.Pow(float64(invGenerator), float64(m))) % n
}

func CreateXbTable(generator, m int64) {
	var n int64 = 50 // for Z50, n = 50
	for i := 0; i < int(m); i++ {
		genpowxb := int64(math.Pow(float64(generator), float64(i))) % n
		XbTable[genpowxb] = int64(i)
	}
}
