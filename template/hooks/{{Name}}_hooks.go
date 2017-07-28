package main

import (
	"strings"

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

	// add additional hooks here

	server.Serve()
	defer server.Listener.Close()
}
