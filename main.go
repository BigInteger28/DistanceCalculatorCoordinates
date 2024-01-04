package main

import (
	"fmt"
	"math/big"
)

// calculateDistanceSquared berekent het kwadraat van de Euclidische afstand tussen twee punten
func calculateDistanceSquared(x1, z1, x2, z2 *big.Int) *big.Int {
	dx := new(big.Int).Sub(x2, x1)
	dz := new(big.Int).Sub(z2, z1)

	dxSquared := new(big.Int).Mul(dx, dx)
	dzSquared := new(big.Int).Mul(dz, dz)

	return new(big.Int).Add(dxSquared, dzSquared)
}

func main() {
	for {
		// Eerste set coördinaten
		var x1, z1 string
		fmt.Println("Voer de X en Z coördinaten in van het eerste punt (grote gehele getallen):")
		fmt.Scan(&x1, &z1)

		// Tweede set coördinaten
		var x2, z2 string
		fmt.Println("Voer de X en Z coördinaten in van het tweede punt (grote gehele getallen):")
		fmt.Scan(&x2, &z2)

		// Omzetten van string naar big.Int
		bigX1 := new(big.Int)
		bigX1.SetString(x1, 10)
		bigZ1 := new(big.Int)
		bigZ1.SetString(z1, 10)
		bigX2 := new(big.Int)
		bigX2.SetString(x2, 10)
		bigZ2 := new(big.Int)
		bigZ2.SetString(z2, 10)

		// Afstand berekenen
		distanceSquared := calculateDistanceSquared(bigX1, bigZ1, bigX2, bigZ2)
		fmt.Printf("Het kwadraat van de afstand tussen de twee punten is: %s\n", distanceSquared.String())
		fmt.Println("\n\n")
	}
}
