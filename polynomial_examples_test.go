package quickFactor

import "fmt"

func ExamplePolynomial_Degree() {
	p := Polynomial{5, -12, 4}
	// 4x² - 12x + 5 --> highest exponent = 2

	fmt.Print(p.Degree()) // Output: 2
}

func ExamplePolynomial_F() {
	x := 7.0
	p := Polynomial{3, 5}
	// 5x + 3 -> 5(7) + 3 == 38

	fmt.Print(p.F(x)) // Output: 38
}
