package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
)

var _ = Describe("the Polynomial type's methods", func() {
	rand.Seed(GinkgoRandomSeed())

	Describe("the Degree method", func() {
		It("should return the length of the array - 1", func() {
			By("creating a new Polynomial with a length between 2 and 10")
			p := Polynomial{}

			l := rand.Intn(8) + 2
			for i := 1; i <= l; i++ {
				p = append(p, float64(i))
			}

			By("checking that the method returns the correct value")
			Expect(p.Degree()).To(Equal(l - 1))
		})
	})
})
