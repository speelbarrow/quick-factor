package main

import "fmt"

func ExamplePolynomial_Degree() {
	p := Polynomial{5, -12, 4}
	// 4x² - 12x + 5 --> highest exponent = 2

	fmt.Println(p.Degree()) // Output: 2
}
