// Define types and functions related to factoring.
package main

// A polynomial in factored form.
type FactoredPolynomial struct {
	Raw         string
	XIntercepts []float64
	Zeroes      []string
	Unfactored  *Polynomial
}

func Factor(unfactored *Polynomial) *FactoredPolynomial {
	return nil
}
