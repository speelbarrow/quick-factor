package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
	"math/rand"
)

var _ = Describe("the Factor function", func() {
	rand.Seed(GinkgoRandomSeed())

	It("should correctly factor polynomial functions", func() {
		By("using a simple polynomial function")
		p := Polynomial{10, 7, 1}

		f := Factor(&p)
		Expect(*f.Unfactored).To(Equal(p))

		Expect(f.Raw).To(ContainSubstring("(x + 2)"))
		Expect(f.Raw).To(ContainSubstring("(x + 5)"))

		Expect(f.XIntercepts).To(ContainElement(-2))
		Expect(f.XIntercepts).To(ContainElement(-5))

		Expect(f.Zeroes).To(ContainElement("x + 2"))
		Expect(f.Zeroes).To(ContainElement("x + 5"))

		By("using a randomly generated polynomial function")
		var (
			product, sum float64
			products     = rand.Perm(50)
		)
		for _, p := range products {
			if p == 1 {
				continue
			}

			divisors := rand.Perm(p)
			for _, divisor := range divisors {
				if d := float64(p) / float64(divisor); d == math.Round(d) {
					product = float64(p)
					sum = float64(divisor) + d
					break
				}
			}

			if product != 0 && sum != 0 {
				break
			}
		}
		p = Polynomial{product, sum, 1}
		f = Factor(&p)

		Expect(*f.Unfactored).To(Equal(p))

		Expect(f.XIntercepts).To(HaveLen(2))
		Expect(f.XIntercepts[0] + f.XIntercepts[1]).To(Equal(sum))
		Expect(f.XIntercepts[0] * f.XIntercepts[1]).To(Equal(product))
		x := float64(rand.Intn(19) + 1)
		Expect((x - f.XIntercepts[0]) * (x - f.XIntercepts[1])).To(Equal(p.F(x)))

		Expect(f.Zeroes).To(HaveLen(2))
	})
})
