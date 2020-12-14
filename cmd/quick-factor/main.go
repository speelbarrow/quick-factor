// Command-line tool to find factors of a polynomial function.
package main

import (
	"bufio"
	"fmt"
	. "github.com/noah-friedman/quick-factor"
	"log"
	"os"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)

func ReadUint() uint {
	if s, e := r.ReadString('\n'); e != nil {
		log.Panicf("ERROR: failed to read input - %e\n", e)
	} else {
		if i, e := strconv.ParseUint(strings.TrimRight(s, "\n"), 10, 32); e != nil {
			log.Fatalln("FATAL: invalid input - must be a positive integer")
		} else {
			return uint(i)
		}
	}

	return 0
}

func ReadFloat64() float64 {
	if s, e := r.ReadString('\n'); e != nil {
		log.Panicf("ERROR: failed to read input - %e\n", e)
	} else {
		if f, e := strconv.ParseFloat(strings.TrimRight(s, "\n"), 64); e != nil {
			log.Fatalln("FATAL: invalid input - must be a number")
		} else {
			return f
		}
	}

	// Only to satisfy compiler, this return statement is never executed.
	return 0
}

func main() {
	fmt.Println("Quick Factor v0.1.0")
	fmt.Println("Developed by Noah Friedman")
	fmt.Println("--------------------------")

	fmt.Print("\nWhat is the degree of the polynomial you'd like to use?: ")
	degree := ReadUint()

	fmt.Println("\nPlease enter the coefficient for each term of the polynomial in the prompts below. Leave a prompt blank to represent 0.")

	p := Polynomial{}
	for i := int(degree); i >= 0; i-- {
		fmt.Printf("Coefficient for x^%d: ", i)
		p = append(Polynomial{ReadFloat64()}, p...)
	}

	fmt.Println("\nChoose an action:")
	fmt.Println("1. Factor")
	fmt.Println("2. Solve with X")
	fmt.Print("Enter your choice: ")
	choice := ReadUint()

	fmt.Print("How many decimal places to round to?: ")
	roundTo := int(ReadUint())

	switch choice {
	case 1:
		f := Factor(p, roundTo)

		fmt.Printf("\nResult: %s\nX Intercepts: %v\n", f.Raw, f.XIntercepts)
	case 2:
		fmt.Print("What value of X?: ")

		fmt.Printf(fmt.Sprintf("\nResult: %%.%dg", roundTo), p.F(ReadFloat64()))
	}
}
