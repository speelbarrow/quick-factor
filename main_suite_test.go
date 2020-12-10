package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestQuickFactor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "quick-factor")
}
