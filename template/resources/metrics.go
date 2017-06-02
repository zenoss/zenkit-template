{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package resources

import (
	"encoding/json"

	"github.com/goadesign/goa"
	"{{$pkg}}/resources/app"
	"github.com/zenoss/zenkit"
)

// MetricsController implements the metrics resource.
type MetricsController struct {
	*goa.Controller
}

// NewMetricsController creates a metrics controller.
func NewMetricsController(service *goa.Service) *MetricsController {
	return &MetricsController{Controller: service.NewController("MetricsController")}
}

// GetMetrics runs the getMetrics action.
func (c *MetricsController) GetMetrics(ctx *app.GetMetricsMetricsContext) error {
	// MetricsController_GetMetrics: start_implement
	registry := zenkit.ContextMetrics(ctx)
	if registry == nil {
		// No registry was registered; must not be using metrics middleware.
		return ctx.OK([]byte("{}"))
	}
	encoder := json.NewEncoder(ctx.ResponseData)
	if ctx.Pretty {
		encoder.SetIndent("", "    ")
	}
	if err := encoder.Encode(registry); err != nil {
		return err
	}
	// MetricsController_GetMetrics: end_implement
	return nil
}
