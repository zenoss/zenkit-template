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

// Redoc runs the redoc action.
func (c *SwaggerController) Redoc(ctx *app.RedocSwaggerContext) error {
	// SwaggerController_Redoc: start_implement
	return ctx.OK([]byte(`<!DOCTYPE html
<html>
  <head>
    <title>{{Description}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='/_swagger/swagger.json'></redoc>
    <script src="https://rebilly.github.io/ReDoc/releases/latest/redoc.min.js"> </script>
  </body>
</html>
`))
	// SwaggerController_Redoc: end_implement
	return nil
}
