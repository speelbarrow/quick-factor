package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
	"math/rand"
)

var _ = Describe("the Factor function", func() {
	rand.Seed(GinkgoRandomSeed())

	It("should factor a simple polynomial function", func() {
		p := Polynomial{10, 7, 1}
		f := Factor(p, 2)

		Expect(f.Unfactored).To(Equal(p))

		Expect(f.Raw).To(ContainSubstring("(x + 2)"))
		Expect(f.Raw).To(ContainSubstring("(x + 5)"))

		Expect(f.XIntercepts).To(HaveLen(2))
		Expect(f.XIntercepts).To(ContainElements(-2., -5.))
	})
	It("should be able to factor trinomial functions using the quadratic formula", func() {
		p := Polynomial{-14, 5, 1}
		f := Factor(p, 2)

		Expect(f.Unfactored).To(Equal(p))

		Expect(f.Raw).To(ContainSubstring("(x - 2)"))
		Expect(f.Raw).To(ContainSubstring("(x + 7)"))

		Expect(f.XIntercepts).To(HaveLen(2))
		Expect(f.XIntercepts).To(ContainElements(2., -7.))
	})
	It("should factor a longer polynomial function", func() {
		p := Polynomial{1, -4, 0, 7, 2}
		f := Factor(p, 2)

		Expect(f.Unfactored).To(Equal(p))

		Expect(f.Raw).To(ContainSubstring("(x + 1)"))
		Expect(f.Raw).To(ContainSubstring("(x - 0.5)"))
		Expect(f.Raw).To(ContainSubstring("(x - 0.3)"))
		Expect(f.Raw).To(ContainSubstring("(x + 3.3)"))

		Expect(f.XIntercepts).To(HaveLen(4))
		By("rounding x intercepts to avoid incorrect test fails")

		for i, v := range f.XIntercepts {
			f.XIntercepts[i] = math.Round(v*10) / 10
		}

		Expect(f.XIntercepts).To(ContainElements(-1., 0.5, 0.3, -3.3))
	})
	It("should factor a randomly generated simple polynomial function", func() {
		var (
			product, sum float64
			products     = rand.Perm(50)
		)
		for _, p := range products {
			if p <= 1 {
				continue
			}
			if rand.Intn(1) == 1 {
				p *= -1
			}

			divisors := rand.Perm(p)
			for _, divisor := range divisors {
				if divisor == 0 {
					continue
				}

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
		p := Polynomial{product, sum, 1}
		f := Factor(p, 2)

		Expect(f.Unfactored).To(Equal(p))

		Expect(f.XIntercepts).To(HaveLen(2))
		Expect(math.Abs(f.XIntercepts[0] + f.XIntercepts[1])).To(Equal(sum))
		Expect(f.XIntercepts[0] * f.XIntercepts[1]).To(Equal(product))
		x := float64(rand.Intn(19) + 1)
		Expect((x - f.XIntercepts[0]) * (x - f.XIntercepts[1])).To(Equal(p.F(x)))
	})
	It("should be able to factor binomial functions", func() {
		p := Polynomial{-2, 3}
		f := Factor(p, 2)

		Expect(f.Unfactored).To(Equal(p))

		Expect(f.Raw).To(Equal("(x - 0.67)"))

		Expect(f.XIntercepts).To(Equal([]float64{2. / 3.}))
	})
	It("should not try to factor a monomial or empty Polynomial", func() {
		By("using a monomial")
		p := Polynomial{5}
		f := Factor(p, 2)

		Expect(f.Unfactored).To(Equal(p))
		Expect(f.Raw).To(Equal("not factorable"))
		Expect(f.XIntercepts).To(HaveLen(0))

		By("using an empty Polynomial")
		p = Polynomial{}
		f = Factor(p, 2)

		Expect(f.Raw).To(Equal("not factorable"))
		Expect(f.XIntercepts).To(HaveLen(0))
	})
})
