package quickFactor

import "fmt"

func ExampleFactor_simpleTrinomial() {
	p := Polynomial{10, 7, 1} // x² + 7x + 10

	f := Factor(p, 2)
	fmt.Println(f.Raw) // Output: (x + 5)(x + 2)
}

func ExampleFactor_quadraticFormula() {
	p := Polynomial{10, 2, 1} // x² + 2x + 10

	f := Factor(p, 2)
	fmt.Println(f.Raw) // Output: (x - ((-2 + √-36) / 2))(x - ((-2 - √-36) / 2))
}

func ExampleFactor_longPolynomial() {
	p := Polynomial{1, -4, 0, 7, 2} // 2x⁴ + 7x³ - 4x + 1

	f := Factor(p, 2)
	fmt.Println(f.Raw) // Output: (x + 1)(x - 0.5)(x - 0.3)(x + 3.3)
}
