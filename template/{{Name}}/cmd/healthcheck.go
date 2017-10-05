package cmd

import (
	"fmt"
	"net/http"
	"time"

	. "github.com/zenoss/zenkit/healthcheck"
	. "github.com/zenoss/zenkit/healthcheck/checks"
)

func registerPing(port int, timeout, period time.Duration) {
	r := fmt.Sprintf("http://localhost:%d/ping", port)
	c := PeriodicChecker(HTTPChecker(r, 200, timeout, http.Header{}), period)
	Register("ping", c)
}
