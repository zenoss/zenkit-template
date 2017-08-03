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
			Param("name", String, "User name", func() {
				Metadata("swagger:extension:x-example", "Keanu")
			})
		})
		Response(OK, Greeting)
		Response(BadRequest)
	})
	Action("add", func() {
		Description("Add two numbers")
		Routing(GET("/add/:a/:b"))
		Params(func() {
			Param("a", Integer, "A number to add", func() {
				Metadata("swagger:extension:x-example", "32")
			})
			Param("b", Integer, "A number to add", func() {
				Metadata("swagger:extension:x-example", "24")
			})
		})
		Response(OK, Sum)
	})
	Action("words", func() {
		Scheme("ws")
		Description("Streams the word 'word' over a websocket")
		Routing(GET("/words"))
		Params(func() {
			Param("count", Integer, "Number of times to say 'word'")
			Param("delay", Integer, "Milliseconds between each word")
			Required("count", "delay")
		})
		Response(SwitchingProtocols)
	})
})

var Greeting = MediaType("application/x.{{Name}}.greeting+json", func() {
	Description("The result of saying hello")
	Attributes(func() {
		Attribute("greeting", String, "The greeting", func() {
			Example("Hello, Keanu!")
		})
		Required("greeting")
	})
	View("default", func() {
		Attribute("greeting")
	})
})

var Sum = MediaType("application/x.{{Name}}.sum+json", func() {
	Description("The sum of two numbers")
	Attributes(func() {
		Attribute("total", Integer, "The sum total", func() {
			Example(56)
		})
		Required("total")
	})
	View("default", func() {
		Attribute("total")
	})
})
