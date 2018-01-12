package design

import (
	. "github.com/goadesign/goa/design/apidsl"
	"github.com/zenoss/zenkit/auth"
)

var _ = API("{{camel Name "-" | title}}", func() {
	Title("{{Title}}")
	Description("{{Description}}")
	Scheme("http")
	Host("localhost:{{Port}}")
	Consumes("application/json")
	Produces("application/json")

	Security(auth.JWT(), func() {
		Scope("api:access")
	})

})
