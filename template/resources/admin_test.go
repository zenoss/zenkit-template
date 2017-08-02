{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package resources_test

import (
	"context"
	"errors"

	"github.com/zenoss/zenkit"
	metrics "github.com/rcrowley/go-metrics"
	. "{{$pkg}}/resources"
	"{{$pkg}}/resources/app/test"
	. "github.com/onsi/ginkgo"
)

// We need a registry that refuses to Marshal
type Registry struct {
	metrics.StandardRegistry
}

func (r *Registry) MarshalJSON() ([]byte, error) {
	return []byte(nil), errors.New("this is a test")
}

var _ = Describe("Admin", func() {

	var (
		t    = GinkgoT()
		ctx  context.Context
		svc  = zenkit.NewService("admin-test", false)
		ctrl = NewAdminController(svc)
	)

	BeforeEach(func() {
		ctx = context.Background()
	})

	Context("when the Metrics resource is requested", func() {

		Context("when the metrics middleware is hooked up", func() {
			BeforeEach(func() {
				registry := metrics.NewRegistry()
				ctx = zenkit.WithMetrics(ctx, registry)
			})
			Context("when the registry cannot be encoded", func() {
				BeforeEach(func() {
					ctx = zenkit.WithMetrics(ctx, &Registry{})
				})
				It("should produce an error", func() {
					test.MetricsAdminInternalServerError(t, ctx, svc, ctrl, true)
				})
			})
			Context("when the registry can be encoded", func() {
				It("should respond OK", func() {
					test.MetricsAdminOK(t, ctx, svc, ctrl, true)
				})
			})
		})

		Context("when the metrics middleware isn't hooked up", func() {
			It("should respond OK", func() {
				test.MetricsAdminOK(t, ctx, svc, ctrl, false)
			})
		})
	})

	Context("when the Ping resource is requested", func() {
		It("should respond OK", func() {
			test.PingAdminOK(t, ctx, svc, ctrl)
		})
	})

	Context("when the Swagger resource is requested", func() {
		It("should respond OK", func() {
			test.SwaggerAdminOK(t, ctx, svc, ctrl)
		})
	})

	Context("when the SwaggerJSON resource is requested", func() {
		originalAsset := SwaggerJSONAsset
		Context("when the swagger.json asset is missing", func() {
			BeforeEach(func() {
				SwaggerJSONAsset = "none"
			})
			AfterEach(func() {
				SwaggerJSONAsset = originalAsset
			})
			It("should respond with an InternalServerError", func() {
				test.SwaggerJSONAdminInternalServerError(t, ctx, svc, ctrl)
			})
		})
		Context("when the swagger.json asset is available", func() {
			It("should respond OK", func() {
				test.SwaggerJSONAdminOK(t, ctx, svc, ctrl)
			})
		})
	})
})
