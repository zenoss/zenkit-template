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

// Metrics runs the metrics action.
func (c *MetricsController) Metrics(ctx *app.MetricsMetricsContext) error {
	// MetricsController_Metrics: start_implement
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
	// MetricsController_Metrics: end_implement
	return nil
}

// Ping runs the ping action.
func (c *MetricsController) Ping(ctx *app.PingMetricsContext) error {
	// MetricsController_Ping: start_implement
	// MetricsController_Ping: end_implement
	return nil
}
