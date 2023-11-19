package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.Exp(big.NewInt(2), big.NewInt(1000), nil)
	b.Exp(big.NewInt(2), big.NewInt(1001), nil)

	multiplicationResult := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s\n", multiplicationResult)

	divisionResult := new(big.Rat).SetFrac(a, b)
	fmt.Printf("Деление: %s\n", divisionResult.FloatString(6))

	additionResult := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s\n", additionResult)

	subtractionResult := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s\n", subtractionResult)
}
