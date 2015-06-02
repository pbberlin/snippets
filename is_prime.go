package main

import "fmt"
import "math/big"

func main() {

	cntr := 0
	for i := 0; i < 120; i++ {

		j := big.NewInt(int64(i))

		isPrime := j.ProbablyPrime(2)

		if isPrime {
			// fmt.Printf("%3vYES ", i)
			fmt.Printf("%3v ", i)
			cntr++
		} else {
			// fmt.Printf("%3v no ", i)
		}

		// if i%12 == 11 {
		// 	fmt.Printf("\n")
		// }
		if cntr%12 == 11 {
			cntr++
			fmt.Printf("\n")
		}

	}
}
