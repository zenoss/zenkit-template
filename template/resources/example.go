{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package resources

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/goadesign/goa"
	"{{$pkg}}/resources/app"
	"github.com/zenoss/zenkit/logging"
	"github.com/zenoss/zenkit/metrics"
	"golang.org/x/net/websocket"
)

// ExampleController implements the example resource.
type ExampleController struct {
	*goa.Controller
}

// NewExampleController creates a example controller.
func NewExampleController(service *goa.Service) *ExampleController {
	return &ExampleController{Controller: service.NewController("ExampleController")}
}

// Add runs the add action.
func (c *ExampleController) Add(ctx *app.AddExampleContext) error {
	// ExampleController_Add: start_implement
	defer logging.LogEntryAndExit(ctx)()
	defer metrics.MeasureTime(ctx)()

	total := ctx.A + ctx.B

	// Keep a running total for no reason
	metrics.IncrementCounter(ctx, "{{camel Name "-" | title}}.Add.total", 1)

	return ctx.OK(&app.X{{camel Name "-" | title}}Sum{Total: total})

	// ExampleController_Add: end_implement
	res := &app.X{{camel Name "-" | title}}Sum{}
	return ctx.OK(res)
}

// Greet runs the greet action.
func (c *ExampleController) Greet(ctx *app.GreetExampleContext) error {
	// ExampleController_Greet: start_implement

	defer logging.LogEntryAndExit(ctx)()
	defer metrics.MeasureTime(ctx)()
	if strings.ToLower(ctx.Name) == "newman" {
		return ctx.BadRequest()
	}
	result := fmt.Sprintf("Hello, %s!", ctx.Name)
	return ctx.OK(&app.X{{camel Name "-" | title}}Greeting{Greeting: result})

	// ExampleController_Greet: end_implement
	res := &app.X{{camel Name "-" | title}}Greeting{}
	return ctx.OK(res)
}

// Words runs the words action.
func (c *ExampleController) Words(ctx *app.WordsExampleContext) error {
       c.WordsWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
       return nil
}

// WordsWSHandler establishes a websocket connection to run the words action.
func (c *ExampleController) WordsWSHandler(ctx *app.WordsExampleContext) websocket.Handler {
       return func(ws *websocket.Conn) {
               // ExampleController_Words: start_implement

               fmt.Fprintf(ws, "Here are %d words\n\n", ctx.Count)
               time.Sleep(time.Millisecond * time.Duration(ctx.Delay))

               for i := 0; i < ctx.Count; i++ {
                       fmt.Fprintln(ws, "word")
                       time.Sleep(time.Millisecond * time.Duration(ctx.Delay))
               }

               fmt.Fprintln(ws, "Done!")
               return

               // ExampleController_Words: end_implement
               ws.Write([]byte("words example"))
               // Dummy echo websocket server
               io.Copy(ws, ws)
       }
}
