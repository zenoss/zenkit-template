package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	"github.com/zenoss/zenkit"
)

var _ = Resource("admin", func() {
	Security(zenkit.JWT(), func() {
		Scope(zenkit.ScopeAPIAdmin)
	})
	BasePath("/_admin")
	Action("ping", func() {
		Description("Respond with a 200 if the service is available")
		Routing(HEAD("/ping"), GET("/ping"))
		Response(OK)
	})
	Action("metrics", func() {
		Description("Return a snapshot of metrics")
		Routing(GET("/metrics"))
		Params(func() {
			Param("pretty", Boolean, "Indent resulting JSON", func() {
				Default(true)
			})
		})
		Response(OK, "application/json")
		Response(InternalServerError)
	})
	Action("swagger.json", func() {
		Description("Retrieve Swagger spec as JSON")
		Routing(GET("/swagger/swagger.json"))
		Response(OK, "application/json")
		Response(InternalServerError)
	})
	Action("swagger", func() {
		Description("Display Swagger using ReDoc")
		Routing(GET("/swagger"))
		Response(OK, "text/html")
	})
})
