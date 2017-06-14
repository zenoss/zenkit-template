{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package resources

import (
	"github.com/goadesign/goa"
	"{{$pkg}}/resources/app"
	"{{$pkg}}/swagger"
)

// SwaggerController implements the swagger resource.
type SwaggerController struct {
	*goa.Controller
}

// NewSwaggerController creates a swagger controller.
func NewSwaggerController(service *goa.Service) *SwaggerController {
	return &SwaggerController{Controller: service.NewController("SwaggerController")}
}

// JSON runs the json action.
func (c *SwaggerController) JSON(ctx *app.JSONSwaggerContext) error {
	// SwaggerController_JSON: start_implement
	data, err := swagger.Asset("swagger/swagger.json")
	if err != nil {
		return err
	}
	return ctx.OK(data)
	// SwaggerController_JSON: end_implement
	return nil
}
