// Define types and functions related to factoring.
package main

import (
	"fmt"
	"log"
	"math"
	"sync"
)

// A polynomial in factored form.
type FactoredPolynomial struct {
	Raw         string
	XIntercepts []float64
	Unfactored  Polynomial
}

// Factor a polynomial function. Round numbers to 'roundTo' decimal places.
func Factor(unfactored Polynomial, roundTo int) *FactoredPolynomial {
	r := &FactoredPolynomial{Unfactored: unfactored}

	if unfactored.Degree() == 2 {
		f := factorTrinomial(&unfactored)
		switch f.(type) {
		case *[2]float64:
			for _, v := range *f.(*[2]float64) {
				sign := "-"
				if a := math.Abs(v); v != a {
					sign = "+"
				}

				r.Raw += fmt.Sprintf(fmt.Sprintf("(x %s %%.%dg)", sign, roundTo), v)
				r.XIntercepts = append(r.XIntercepts, v*-1)
			}
		case *[2]string:
		}
	}

	return r
}

// A helper function that factors trinomials only. Panics if the provided Polynomial has a degree that is not equal to 2. Will return either *[2]float64 or *[2]string.
func factorTrinomial(trinomial *Polynomial) interface{} {
	if d := trinomial.Degree(); d != 2 {
		log.Panicf("expected Polynomial with degree 2, got degree %d\n", d)
	}

	if (*trinomial)[0] != 0 {
		if r := trinomialByGrouping(trinomial); r != nil {
			return r
		}
	}

	return trinomialByQuadraticFormula(trinomial)
}

// A helper function that factors trinomials by grouping. Returns nil if the trinomial cannot be factored by grouping.
func trinomialByGrouping(trinomial *Polynomial) *[2]float64 {
	var (
		product = (*trinomial)[0] * (*trinomial)[2]
		wg      sync.WaitGroup
		group   = make(chan *[2]float64)
		check   = func(a, b float64) *[2]float64 {
			defer wg.Done()

			if a*b == product {
				if (*trinomial)[1] < 0 {
					a *= -1
					b *= -1
				}

				group <- &[2]float64{a, b}
			}
			return nil
		}
	)
	defer close(group)

	if product > 0 {
		for a, b := math.Abs((*trinomial)[1])-1, 1.; a >= b; a, b = a-1, b+1 {
			wg.Add(1)
			go check(a, b)
		}
	} else if product < 0 {
		for a, b := math.Abs((*trinomial)[1])+1, -1.; a*b <= product; a, b = a+1, b-1 {
			wg.Add(1)
			go check(a, b)
		}
	}

	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case r := <-group:
		return r
	case <-done:
		return nil
	}
}

// A helper function that factors trinomials by using the quadratic formula.
func trinomialByQuadraticFormula(trinomial *Polynomial) *[2]string {
	return nil
}
