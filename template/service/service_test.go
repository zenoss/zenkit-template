package service_test

import (
	"context"

	. "github.com/zenoss/{{Name}}/service"
	proto "github.com/zenoss/zing-proto/go/{{Name}}"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {

	var (
		ctx context.Context
		svc proto.{{Name | title}}ServiceServer
	)

	BeforeEach(func() {
		ctx = context.Background()
		svc = New{{Name | title}}Service()
	})

	It("should do something {{Name}}ish", func() {
		Ω(true).Should(BeTrue())
	})

})
