{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"{{$pkg}}/resources/app"
	"github.com/snikch/goodman/hooks"
	trans "github.com/snikch/goodman/transaction"
)

func main() {
	h := hooks.NewHooks()
	server := hooks.NewServer(hooks.NewHooksRunner(h))

	// skip admin calls
	h.BeforeEach(func(t *trans.Transaction) {
		t.Skip = strings.HasPrefix(t.Name, "admin >")
	})

	// skip the 400 result from the example
	h.Before("example > /hello/{name} > greet example > 400 > application/x.{{Name}}.greeting+json", func(t *trans.Transaction) {
		t.Skip = true
	})

	// skip the websocket example
	h.Before("example > /words > words example > 101 > application/json", func(t *trans.Transaction) {
		t.Skip = true
	})

	// a detailed example
	h.After("example > /add/{a}/{b} > add example > 200 > application/x.tester.sum+json", func(t *trans.Transaction) {
		parts := strings.Split(t.FullPath, "/")
		a, _ := strconv.Atoi(parts[2])
		b, _ := strconv.Atoi(parts[3])

		actual := &app.X{{camel Name "-" | title}}Sum{}
		_ = json.Unmarshal([]byte(t.Real.Body), actual)

		if total := a + b; total != actual.Total {
			t.Fail = fmt.Sprintf("expected (%d) and actual (%d) values do not match", total, actual.Total)
		}
	})

	// add additional hooks here

	server.Serve()
	defer server.Listener.Close()
}
