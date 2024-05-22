package mpc

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// implements Shamir's secret sharing algorithm

type ShamirSecretSharing struct {
	Prime     *big.Int
	Shares    []*big.Int
	Threshold int
}

func NewShamirSecretSharing(secret *big.Int, Threshold, numShares int) (*ShamirSecretSharing, error) {
	Prime, err := rand.Prime(rand.Reader, secret.BitLen()+numShares)
	if err != nil {
		return nil, err
	}

	sss := &ShamirSecretSharing{
		Prime:     Prime,
		Shares:    make([]*big.Int, numShares),
		Threshold: Threshold,
	}

	return sss, sss.Share(secret)
}

func (sss *ShamirSecretSharing) Share(secret *big.Int) error {
	coeffs := make([]*big.Int, sss.Threshold)
	coeffs[0] = new(big.Int).Set(secret)

	// Generate random coefficients for the polynomial
	for i := 1; i < sss.Threshold; i++ {
		coeffs[i], _ = rand.Int(rand.Reader, sss.Prime)
	}

	// Calculate Shares using Lagrange interpolation
	for i := range sss.Shares {
		x := big.NewInt(int64(i + 1))
		share := big.NewInt(0)

		for j := 0; j < sss.Threshold; j++ {
			term := big.NewInt(0)
			term.Mul(coeffs[j], x.Exp(x, big.NewInt(int64(j)), nil))
			for k := 0; k < j; k++ {
				denom := big.NewInt(int64(k + 1))
				denom.Sub(x, denom)
				if denom.Sign() == 0 {
					continue
				}
				term.Div(term, denom)
			}
			share.Add(share, term)
		}
		share.Mod(share, sss.Prime)
		sss.Shares[i] = share
	}

	return nil
}

func (sss *ShamirSecretSharing) Combine(Shares []*big.Int) (*big.Int, error) {
	if len(Shares) < sss.Threshold {
		return nil, errors.New("insufficient number of Shares")
	}

	secret := big.NewInt(0)
	product := big.NewInt(1)

	for i, x := range Shares[:sss.Threshold] {
		xi := big.NewInt(int64(i + 1))
		product.Mul(product, xi)

		term := big.NewInt(1)
		for j := range Shares[:sss.Threshold] {
			if i == j {
				continue
			}
			div := big.NewInt(0).Sub(xi, big.NewInt(int64(j+1)))
			div.ModInverse(div, sss.Prime)
			term.Mul(term, div)
		}

		term.Mul(term, x)
		term.Mod(term, sss.Prime)
		secret.Add(secret, term)
		secret.Mod(secret, sss.Prime)
	}

	product.ModInverse(product, sss.Prime)
	secret.Mul(secret, product)
	secret.Mod(secret, sss.Prime)

	return secret, nil
}
