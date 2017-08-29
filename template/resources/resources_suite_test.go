package resources_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

func TestResources(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	rand.Seed(GinkgoRandomSeed())
	RunSpecsWithDefaultAndCustomReporters(t, "{{camel Name "-" | title}} Resources Suite", []Reporter{junitReporter})
}
