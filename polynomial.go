// Define the Polynomial type and its methods.
package main

// A type that represents a polynomial function. Each array index is mapped to the corresponding term.
// For example,
//  Polynomial{5, -13, 1, 4}
//
//  Index:     0    1  2  3
// represents the polynomial function:
//              x³ + x² - 13x + 5
//
//  Power of x: 3    2      1   0
type Polynomial []float64

// Returns the degree (highest exponent) of the polynomial function.
func (p *Polynomial) Degree() int {
	return len(*p) - 1
}
