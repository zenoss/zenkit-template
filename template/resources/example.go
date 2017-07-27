{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package resources

import (
	"fmt"
	"strings"

	"github.com/goadesign/goa"
	metrics "github.com/rcrowley/go-metrics"
	"{{$pkg}}/resources/app"
	"github.com/zenoss/zenkit"
)

// ExampleController implements the example resource.
type ExampleController struct {
	*goa.Controller
}

// NewExampleController creates a example controller.
func NewExampleController(service *goa.Service) *ExampleController {
	return &ExampleController{Controller: service.NewController("ExampleController")}
}

// Add runs the add action.
func (c *ExampleController) Add(ctx *app.AddExampleContext) error {
	// ExampleController_Add: start_implement
	defer zenkit.LogEntryAndExit(ctx)()
	defer zenkit.MeasureTime(ctx)()

	total := ctx.A + ctx.B

	// Keep a running total for no reason
	ctr := metrics.GetOrRegisterCounter("{{Name}}.Add.total", zenkit.ContextMetrics(ctx))
	ctr.Inc(int64(total))

	return ctx.OK(&app.XTesterSum{Total: total})

	// ExampleController_Add: end_implement
	res := &app.XTesterSum{}
	return ctx.OK(res)
}

// Greet runs the greet action.
func (c *ExampleController) Greet(ctx *app.GreetExampleContext) error {
	// ExampleController_Greet: start_implement

	defer zenkit.LogEntryAndExit(ctx)()
	defer zenkit.MeasureTime(ctx)()
	if strings.ToLower(ctx.Name) == "newman" {
		return ctx.BadRequest()
	}
	result := fmt.Sprintf("Hello, %s!", ctx.Name)
	return ctx.OK(&app.X{{Name | title}}Greeting{Greeting: result})

	// ExampleController_Greet: end_implement
	res := &app.X{{Name | title}}Greeting{}
	return ctx.OK(res)
}
