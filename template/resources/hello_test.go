{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package resources_test

import (
	"context"
	"fmt"

	"github.com/goadesign/goa"
	. "{{$pkg}}/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"{{$pkg}}/resources/app/test"
)

var _ = Describe("Hello", func() {

	var (
		t    = GinkgoT()
		ctx  context.Context
		svc  = goa.New("hello-test")
		ctrl = NewHelloController(svc)
	)

	BeforeEach(func() {
		ctx = context.Background()
	})

	It("should say hello to a normal user", func() {
		var (
			name     = "tester"
			expected = fmt.Sprintf("Hello, %s!", name)
		)
		_, greeting := test.SayhelloHelloOK(t, ctx, svc, ctrl, name)
		Ω(*greeting.Greeting).Should(Equal(expected))
	})

	It("should not say hello to Newman", func() {
		var name = "newman"
		test.SayhelloHelloBadRequest(t, ctx, svc, ctrl, name)
	})

})
