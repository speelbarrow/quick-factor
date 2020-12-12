package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
	"math/rand"
)

var _ = Describe("the Polynomial type's methods", func() {
	rand.Seed(GinkgoRandomSeed())

	Describe("the Degree method", func() {
		It("should return the degree of the polynomial", func() {
			By("creating a Polynomial with a length between 2 and 10")
			p := Polynomial{}

			l := rand.Intn(8) + 2
			for i := 1; i <= l; i++ {
				p = append(p, float64(i))
			}

			Expect(p.Degree()).To(Equal(l - 1))
		})
		It("should not acknowledge highest indices with values of 0", func() {
			By("creating a Polynomial with highest indices at 0")
			p := Polynomial{3, 2, 1, 0, 0, 0}

			Expect(p.Degree()).To(Equal(2))
		})
	})
	Describe("the F method", func() {
		It("should return the function value with the given 'x' value", func() {
			By("generating a 2 - 5 term polynomial")
			p := Polynomial{}

			l := rand.Intn(3) + 2
			for i := 1; i <= l; i++ {
				p = append(p, float64(rand.Intn(9)+1))
			}

			By("using a randomly generated 'x' value")
			x := float64(rand.Intn(2999)+1) / 100

			By("calculating the expected value")
			var expected float64
			for i, v := range p {
				expected += v * math.Pow(x, float64(i))
			}

			Expect(p.F(x)).To(Equal(expected))
		})
	})
})
