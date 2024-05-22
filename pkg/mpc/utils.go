package mpc

import (
	"errors"
	"math/big"
)

// ModInverse calculates the modular inverse of a number a modulo m.
func ModInverse(a, m *big.Int) (*big.Int, error) {
	if a.Sign() == 0 {
		return big.NewInt(0), nil
	}

	g, x := big.NewInt(0), big.NewInt(0)
	g.GCD(x, nil, a, m)
	if g.Cmp(big.NewInt(1)) != 0 {
		return nil, errors.New("modular inverse does not exist")
	}

	x.Mod(x, m)
	if x.Sign() < 0 {
		x.Add(x, m)
	}

	return x, nil
}

// PolynomialEval evaluates a polynomial at a given point x.
func PolynomialEval(coeffs []*big.Int, x, prime *big.Int) *big.Int {
	result := big.NewInt(0)

	for i := range coeffs {
		term := big.NewInt(0)
		term.Exp(x, big.NewInt(int64(i)), prime)
		term.Mul(term, coeffs[i])
		term.Mod(term, prime)
		result.Add(result, term)
		result.Mod(result, prime)
	}

	return result
}
