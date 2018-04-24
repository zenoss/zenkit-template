package main_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

func Test{{Name | title}}(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	rand.Seed(GinkgoRandomSeed())
	RunSpecsWithDefaultAndCustomReporters(t, "{{Name | title}} Test Suite", []Reporter{junitReporter})
}
