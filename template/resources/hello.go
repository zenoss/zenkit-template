{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package resources

import (
	"fmt"
	"strings"

	"github.com/goadesign/goa"
	"{{$pkg}}/resources/app"
	"github.com/zenoss/zenkit"
)

// HelloController implements the hello resource.
type HelloController struct {
	*goa.Controller
}

// NewHelloController creates a hello controller.
func NewHelloController(service *goa.Service) *HelloController {
	return &HelloController{Controller: service.NewController("HelloController")}
}

// Sayhello runs the sayhello action.
func (c *HelloController) Sayhello(ctx *app.SayhelloHelloContext) error {
	// HelloController_Sayhello: start_implement
	defer zenkit.LogEntryAndExit(ctx)()
	defer zenkit.MeasureTime(ctx)()
	if strings.ToLower(ctx.Name) == "newman" {
		return ctx.BadRequest()
	}
	result := fmt.Sprintf("Hello, %s!", ctx.Name)
	return ctx.OK([]byte(result))
	// HelloController_Sayhello: end_implement
	return nil
}
