# {{Title}} ({{Name}})
{{Description}}

## Purpose
TODO: Describe what this service does in general terms.

## Healthchecks
TODO: Describe the healthchecks emitted by this service and what they indicate.
* `ping`: Pings the local admin port.

## Metrics
TODO: Describe the metrics exposed by this service, and what they indicate.

## API
TODO: Describe the API endpoints this service provides, and what they do. No
need to go into as much specificity as the Swagger spec will, but be
descriptive.

## Configuration
* `{{replace Name "-" "_" -1 | toUpper}}_LOG_LEVEL`: Log level. Defaults to "info".
* `{{replace Name "-" "_" -1 | toUpper}}_HTTP_PORT`: Port on which the HTTP API service listens. Defaults to {{Port}}.
* `{{replace Name "-" "_" -1 | toUpper}}_ADMIN_PORT`: Port on which the Admin API service listens. Defaults to {{AdminPort}}.
* `{{replace Name "-" "_" -1 | toUpper}}_AUTH_DISABLED`: Whether authentication is enforced. If true, middleware is used that injects an admin identity into unauthenticated requests.
* `{{replace Name "-" "_" -1 | toUpper}}_AUTH_KEY_FILE`: The file containing the secret key that is used to validate incoming authentication tokens.
* `{{replace Name "-" "_" -1 | toUpper}}_TRACING_ENABLED`: Whether request tracing is enabled.
