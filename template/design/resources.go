package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("metrics", func() {
	BasePath("/")
	Action("ping", func() {
		Description("Respond with a 200 if the service is available")
		Routing(GET("_ping"))
		Response(OK)
	})
	Action("metrics", func() {
		Description("Return a snapshot of metrics")
		Routing(GET("_metrics"))
		Params(func() {
			Param("pretty", Boolean, "Indent resulting JSON", func() {
				Default(true)
			})
		})
		Response(OK, "application/json")
	})
})

// This exists for example purposes only.
// Please see https://goa.design/learn/guide to get started.
// Full DSL docs can be found at https://goa.design/reference/goa/design/apidsl/.
var _ = Resource("hello", func() {
	BasePath("/hello")
	Action("sayhello", func() {
		Description("Say hello to somebody")
		Routing(GET("/:name"))
		Params(func() {
			Param("name", String, "User name")
		})
		Response(OK, Greeting)
		Response(BadRequest)
	})
})

var Greeting = MediaType("application/x.{{Name}}.greeting+json", func() {
	Description("The result of saying hello")
	Attributes(func() {
		Attribute("greeting", String, "The greeting")
		Required("greeting")
	})
	View("default", func() {
		Attribute("greeting")
	})
})
