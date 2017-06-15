{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package resources

import (
	"encoding/json"

	"github.com/goadesign/goa"
	"{{$pkg}}/resources/app"
	"{{$pkg}}/swagger"
	"github.com/zenoss/zenkit"
)

// AdminController implements the admin resource.
type AdminController struct {
	*goa.Controller
}

// NewAdminController creates a admin controller.
func NewAdminController(service *goa.Service) *AdminController {
	return &AdminController{Controller: service.NewController("AdminController")}
}

// Metrics runs the metrics action.
func (c *AdminController) Metrics(ctx *app.MetricsAdminContext) error {
	// AdminController_Metrics: start_implement
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
	// AdminController_Metrics: end_implement
	return nil
}

// Ping runs the ping action.
func (c *AdminController) Ping(ctx *app.PingAdminContext) error {
	// AdminController_Ping: start_implement
	return ctx.OK([]byte{`PONG`})
	// AdminController_Ping: end_implement
	return nil
}

// Swagger runs the swagger action.
func (c *AdminController) Swagger(ctx *app.SwaggerAdminContext) error {
	// AdminController_Swagger: start_implement
	return ctx.OK([]byte(`<!DOCTYPE html
<html>
  <head>
    <title>Topo API</title>
    <!-- needed for adaptive design -->
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!--
    ReDoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='/_admin/swagger/swagger.json'></redoc>
    <script src="https://rebilly.github.io/ReDoc/releases/latest/redoc.min.js"> </script>
  </body>
</html>
`))
	// AdminController_Swagger: end_implement
	return nil
}

// SwaggerJSON runs the swagger.json action.
func (c *AdminController) SwaggerJSON(ctx *app.SwaggerJSONAdminContext) error {
	// AdminController_SwaggerJSON: start_implement
	data, err := swagger.Asset("swagger/swagger.json")
	if err != nil {
		return err
	}
	return ctx.OK(data)
	// AdminController_SwaggerJSON: end_implement
	return nil
}
