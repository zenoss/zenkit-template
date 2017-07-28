package main

import (
	"github.com/snikch/goodman/hooks"
	trans "github.com/snikch/goodman/transaction"
)

var AdminAPI = []string{
	"admin > /_admin/metrics > metrics admin > 200 > application/json",
	"admin > /_admin/ping > ping admin > 200",
	"admin > /_admin/swagger > swagger admin > 200",
	"admin > /_admin/swagger/swagger.json > swagger.json admin > 200 > application/json",
}

func main() {
	h := hooks.NewHooks()
	server := hooks.NewServer(hooks.NewHooksRunner(h))

	// skip admin calls
	for _, api := range AdminAPI {
		h.Before(api, func(t *trans.Transaction) {
			t.Skip = true
		})
	}

	// add additional hooks here

	server.Serve()
	defer server.Listener.Close()
}
