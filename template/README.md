# {{Title}} ({{Name}})
{{Description}}

## tl;dr
* _`make`_ to regenerate code after changes to the API
* _`make run`_ to run the service using docker-compose
* _`make test`_ to run tests
* _`make build`_ to build the image `zenoss/zing-{{Name}}:dev`
* _`glide up`_ to update dependencies

## What's Inside
This microservice comprises:
* [Cobra](https://github.com/spf13/cobra) for CLI. Cobra files live under the
  `cmd` directory and are created using the `cobra` command line application.
* [Viper](https://github.com/spf13/viper) for configuration. All configuration
  is able to be specified via environment variables and config file, and live
  reloading of configuration is supported.
* [Goa](https://goa.design/) for APIs, service boilerplate, security, and
  [Swagger](http://swagger.io/) generation. Much of the development process
  involves modifying the resources defined in `design/resources.go` and using
  the `goagen` tool (encapsulated fully by the Makefile) to regenerate
  scaffolding code and boilerplate, then adding business logic.
* [go-metrics](https://github.com/rcrowley/go-metrics) for metrics.
* [Logrus](https://github.com/sirupsen/logrus) for structured logging.
* [Ginkgo](https://onsi.github.io/ginkgo/) and
  [Gomega](https://onsi.github.io/gomega) for testing.

## Prerequisites
To develop and run `{{Name}}`, you will need:
* make
* docker-ce >= 17.05 (Official installation instructions for [Ubuntu](https://docs.docker.com/engine/installation/linux/ubuntu/) | [CentOS](https://docs.docker.com/engine/installation/linux/centos/) | [macOS](https://docs.docker.com/docker-for-mac/install/))
* docker-compose (Install from here via `sudo make docker-compose`, or see the [official instructions](https://docs.docker.com/compose/install/))
* A Go environment. [gvm](https://github.com/moovweb/gvm) is an easy way to get one.

Additional helpful utilities include:
* [ginkgo](https://onsi.github.io/ginkgo/) (`go get github.com/onsi/ginkgo/ginkgo`)
* [cobra](https://github.com/spf13/cobra) (`go get github.com/spf13/cobra/cobra`)
* [httpie](https://httpie.org/) (`apt install httpie` on Ubuntu)
* [jq](https://stedolan.github.io/jq/) (`apt install jq` on Ubuntu)

## Development
1. Add or modify resources and actions in `design/resources.go`, using [Goa's
   DSL](https://goa.design/reference/goa/design/apidsl/). The
   [goa-cellar](https://github.com/goadesign/goa-cellar) example implementation
   may also be a useful reference.

2. `make app`. This will generate scaffolding code in the `resources`
   directory, or modify existing scaffolding.
   <aside class="notice">Note: Goa owns all the generated code under `resources/app`. Don't modify it.</aside>

3. Implement the resource action you've just defined. You'll find commented
   body in the boilerplate methods:

        // ControllerName_Action: start_implement

        // Put your logic here

        // ControllerName_Action: end_implement

   Like it says, put your logic in between the two outer comments. This allows
   `goagen` to regenerate the scaffolding around your logic as needed.

4. Add tests for your new code. There may already be a `CONTROLLER_test.go`
   defined. If not, run `ginkgo generate CONTROLLER`, where `CONTROLLER` is, of
   course, the name of the Go file containing your controller implementations.
   Goa generates test helpers for all resources to validate the contract, so
   that the DSL matches the implementation matches the Swagger output. You can
   lean on these in your tests to write them much faster, simply passing in the
   arguments that you expect to trigger each response.

5. `make test` or `go test ./...` or `ginkgo -r`. You may also run tests
   automatically on save by running `ginkgo watch resources` or `ginkgo watch
   -r`.

6. `make run` to rebuild the image and redeploy the service locally. This will
   bring it up on port {{Port}}, allowing you to use `curl` or `httpie`.  You
   may also simply use `go build {{Name}}`, then run the resulting binary
   manually, although if supporting services are required, the `docker-compose`
   functionality the Makefile implements is very convenient.

7. Modify the runbook (runbook.md) to account for your new functionality.


## Building for production
The only artifact this produces is the Docker image `zenoss/zing-{{Name}}`. It will
be pushed to Amazon ECR automatically by our build system once changes have
been verified.

## Environment Variables
* `{{Name | toUpper}}_PORT`: {{Port}}

