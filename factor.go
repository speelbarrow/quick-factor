// Define types and functions related to factoring.
package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"sync"
)

// A polynomial in factored form.
type FactoredPolynomial struct {
	Raw         string
	XIntercepts []float64
	Unfactored  Polynomial
}

// Factor a polynomial function. Round numbers displayed in the 'Raw' field to 'roundTo' decimal places. If the function is not factorable, the 'Raw' field will read "not factorable".
func Factor(unfactored Polynomial, roundTo int) *FactoredPolynomial {
	r := &FactoredPolynomial{Unfactored: unfactored}

	for unfactored.Degree() > 2 {
		divisor := simplestFactor(&unfactored)
		unfactored = *dividePolynomial(&unfactored, divisor)
		r.Raw += formatFactoredForm(divisor, roundTo)
		r.XIntercepts = append(r.XIntercepts, divisor)
	}

	if d := unfactored.Degree(); d == 2 {
		f := factorTrinomial(&unfactored, roundTo)
		switch f.(type) {
		case *[2]float64:
			for _, v := range *f.(*[2]float64) {
				v *= -1

				r.Raw += formatFactoredForm(v, roundTo)
				r.XIntercepts = append(r.XIntercepts, v)
			}
		case *[2]string:
			r.Raw = fmt.Sprintf("(x - (%s))(x - (%s))", (*f.(*[2]string))[0], (*f.(*[2]string))[1])
		}
	} else if d == 1 {
		x := (unfactored[0] * -1) / unfactored[1]
		r.Raw = formatFactoredForm(x, roundTo)
		r.XIntercepts = []float64{x}
	} else { // In this case the degree must be equal to 0
		r.Raw = "not factorable"
	}

	return r
}

// A helper function to transform a number into it's representation in factored form.
func formatFactoredForm(x float64, roundTo int) string {
	sign := "-"
	if x != math.Abs(x) {
		sign = "+"
	}

	return fmt.Sprintf(fmt.Sprintf("(x %s %%.%dg)", sign, roundTo), math.Abs(x))
}

// A helper function that factors trinomials only. Panics if the provided Polynomial has a degree that is not equal to 2. Will return either *[2]float64 or *[2]string.
func factorTrinomial(trinomial *Polynomial, roundTo int) interface{} {
	if d := trinomial.Degree(); d != 2 {
		log.Panicf("expected Polynomial with degree 2, got degree %d\n", d)
	}

	if (*trinomial)[0] != 0 {
		if r := trinomialByGrouping(trinomial); r != nil {
			return r
		}
	}

	return trinomialByQuadraticFormula(trinomial, roundTo)
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
		for a, b := math.Abs((*trinomial)[1])+1, -1.; a*b >= product; a, b = a+1, b-1 {
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

// A helper function that factors trinomials by using the quadratic formula. Will return either *[2]float64 or *[2]string.
func trinomialByQuadraticFormula(trinomial *Polynomial, roundTo int) interface{} {
	bSquaredMinus4ac := math.Pow((*trinomial)[1], 2) - (4 * (*trinomial)[2] * (*trinomial)[0])
	if bSquaredMinus4ac < 0 {
		firstHalf := fmt.Sprintf(fmt.Sprintf("(%%.%dg ", roundTo), (*trinomial)[1]*-1)
		secondHalf := fmt.Sprintf(fmt.Sprintf(" √%%.%[1]dg) / %%.%[1]dg", roundTo), bSquaredMinus4ac, 2*(*trinomial)[2])

		return &[2]string{
			firstHalf + "+" + secondHalf,
			firstHalf + "-" + secondHalf,
		}
	} else {
		negativeB := (*trinomial)[1] * -1
		squareRootOfBSquaredMinus4ac := math.Sqrt(bSquaredMinus4ac)
		twoA := 2 * (*trinomial)[2]

		// Multiply by -1 so that when it is flipped back in the Factor function it results in a valid answer.
		return &[2]float64{
			((negativeB + squareRootOfBSquaredMinus4ac) / twoA) * -1,
			((negativeB - squareRootOfBSquaredMinus4ac) / twoA) * -1,
		}
	}
}

// A helper function to divide a polynomial with a degree > 2. Panics if division results in a remainder.
func dividePolynomial(p *Polynomial, divisor float64) *Polynomial {
	r := Polynomial{(*p)[len(*p)-1]}

	for i, last := len(*p)-2, (*p)[len(*p)-1]*divisor; i > 0; i, last = i-1, r[0]*divisor {
		r = append(Polynomial{(*p)[i] + last}, r...)
	}

	if rem := (*p)[0] + (r[0] * divisor); rem != 0 {
		log.Panicf(
			"invalid division - found remainder\n"+
				"Original: %v\n"+
				"Factored: %v\n"+
				"Divisor: %g\n"+
				"Remainder: %g\n",
			*p, r, divisor, rem,
		)
	}

	return &r
}

// A helper function to determine the simplest factor of a polynomial function with a degree < 2. Will return 0 if there is no factor.
func simplestFactor(p *Polynomial) float64 {
	var (
		aFactors = findFactorsOfN((*p)[len(*p)-1])
		cFactors = findFactorsOfN((*p)[0])
		r        = make(chan float64, len(aFactors)*len(cFactors)) // Create buffered channel so it doesn't block when values are sent after first value (even though only the first value is read)
		wg       sync.WaitGroup
	)

	for _, a := range aFactors {
		for _, c := range cFactors {
			wg.Add(1)
			go func(a, c float64) {
				defer wg.Done()

				if p.F(c/a) == 0 {
					r <- c / a
				} else if p.F((c/a)*-1) == 0 {
					r <- (c / a) * -1
				}
			}(a, c)
		}
	}

	done := make(chan bool)
	defer close(done)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case f := <-r:
		<-done // Wait for all goroutines to finish before closing 'r' to avoid sending on a closed channel
		close(r)
		return f
	case <-done:
		return 0
	}
}

// A helper function that finds all the factors of 'n'.
func findFactorsOfN(n float64) []float64 {
	r := []float64{1, n}

	if n != math.Round(n) {
		return r
	}

	var (
		rLock sync.RWMutex
		wg    sync.WaitGroup
	)
	for i := 2.; i < n; i++ {
		wg.Add(1)
		go func(i float64) {
			defer wg.Done()

			if d := n / i; d == math.Round(d) {
				rLock.Lock()
				r = append(r, d)
				rLock.Unlock()
			}
		}(i)
	}

	wg.Wait()
	sort.Slice(r, func(i, j int) bool {
		if r[i] > 0 && r[j] < 0 {
			return true
		} else if r[i] < 0 && r[j] > 0 {
			return false
		} else {
			return r[i] < r[j]
		}
	})
	return r
}
