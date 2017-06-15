package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This exists for example purposes only.
// Please see https://goa.design/learn/guide to get started.
// Full DSL docs can be found at https://goa.design/reference/goa/design/apidsl/.
var _ = Resource("example", func() {
	BasePath("/")
	Action("greet", func() {
		Description("Say hello to somebody")
		Routing(GET("/hello/:name"))
		Params(func() {
			Param("name", String, "User name")
		})
		Response(OK, Greeting)
		Response(BadRequest)
	})
	Action("add", func() {
		Description("Add two numbers")
		Routing(GET("/add/:a/:b"))
		Params(func() {
			Param("a", Integer, "A number to add")
			Param("b", Integer, "A number to add")
		})
		Response(OK, Integer)
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
