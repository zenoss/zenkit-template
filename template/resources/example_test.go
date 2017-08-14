{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package resources_test

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/goadesign/goa"
	. "{{$pkg}}/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"{{$pkg}}/resources/app/test"
)

var _ = Describe("Example", func() {

	var (
		t    = GinkgoT()
		ctx  context.Context
		svc  = goa.New("example-test")
		ctrl = NewExampleController(svc)
	)

	BeforeEach(func() {
		ctx = context.Background()
	})

	Context("when the Greet resource is requested", func() {
		It("should say hello to a normal user", func() {
			var (
				name     = "tester"
				expected = fmt.Sprintf("Hello, %s!", name)
			)
			_, greeting := test.GreetExampleOK(t, ctx, svc, ctrl, name)
			Expect(greeting.Greeting).Should(Equal(expected))
		})

		It("should not say hello to Newman", func() {
			var name = "newman"
			test.GreetExampleBadRequest(t, ctx, svc, ctrl, name)
		})
	})

	Context("when the Add resource is requested", func() {
		It("should add properly", func() {
			_, sum := test.AddExampleOK(t, ctx, svc, ctrl, 9000, 1)
			// it should be over 9000, obviously
			Expect(sum.Total).Should(Equal(9001))
		})
	})

	Context("when the Words resource is requested", func() {
		It("should stream the number of words specified", func() {
			count := rand.Intn(10) + 1
			conn, err := test.WordsExampleWSTestHelper(t, ctx, svc, ctrl, count, 5)
			Ω(err).ShouldNot(HaveOccurred())
			defer conn.Close()
			buf := gbytes.BufferReader(conn.UnderlyingConn())
			for i := 0; i < count; i++ {
				Eventually(buf).Should(gbytes.Say("word"))
			}
			Eventually(buf).Should(gbytes.Say("Done!"))
		})
	})

})
