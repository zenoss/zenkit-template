package design

import (
	_ "github.com/goadesign/goa/goagen" // This is unnecessary, but helps out glide

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("{{.Name}}", func() {
	Title("{{.Title}}")
	Description("{{.Description}}")
	Scheme("http")
	Host("localhost:{{.Port}}")
})
