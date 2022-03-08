package service_test

import (
	"context"

	. "github.com/zenoss/{{Name}}/service"
	// proto "github.com/zenoss/zing-proto/v11/go/{{Name}}"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {

	var (
		ctx context.Context
		// svc proto.{{Name | title}}Service
	)

	BeforeEach(func() {
		ctx = context.Background()
		// svc = New{{Name | title}}ServiceServer()
	})

	It("should do something {{Name}}ish", func() {
		Expect(true).Should(BeTrue())
	})

})
