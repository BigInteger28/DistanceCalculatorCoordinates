package main

import (
	"fmt"
	"math/big"
	"strings"
)

// calculateDistance berekent de Euclidische afstand tussen twee punten
func calculateDistance(x1, z1, x2, z2 *big.Int) *big.Int {
	dx := new(big.Int).Sub(x2, x1)
	dz := new(big.Int).Sub(z2, z1)

	dxSquared := new(big.Int).Mul(dx, dx)
	dzSquared := new(big.Int).Mul(dz, dz)

	distanceSquared := new(big.Int).Add(dxSquared, dzSquared)

	return distanceSquared.Sqrt(distanceSquared)
}

// formatBigNumber formatteert een big.Int om een spatie toe te voegen elke duizend eenheden
func formatBigNumber(number *big.Int) string {
	numberStr := number.String()
	var formattedStrBuilder strings.Builder
	length := len(numberStr)

	for i, digit := range numberStr {
		// Voeg het cijfer toe
		formattedStrBuilder.WriteRune(digit)

		// Voeg een spatie toe na elke derde cijfer, behalve aan het eind van het nummer
		if (length-i-1)%3 == 0 && i < length-1 {
			formattedStrBuilder.WriteRune(' ')
		}
	}

	return formattedStrBuilder.String()
}

func main() {
	for {
		// Eerste set coördinaten
		var x1, z1 string
		fmt.Println("Voer de X en Z coördinaten in van het eerste punt (gescheiden door een spatie):")
		fmt.Scan(&x1, &z1)

		// Tweede set coördinaten
		var x2, z2 string
		fmt.Println("Voer de X en Z coördinaten in van het tweede punt (gescheiden door een spatie):")
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

		// Afstand berekenen en formatteren
		distance := calculateDistance(bigX1, bigZ1, bigX2, bigZ2)
		formattedDistance := formatBigNumber(distance)

		fmt.Printf("De afstand tussen de twee punten is: %s\n", formattedDistance)
		fmt.Println("\n\n")
	}
}
