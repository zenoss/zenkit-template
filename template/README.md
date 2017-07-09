# {{Title}} ({{Name}})
{{Description}}

## Purpose
TODO: Describe what this service does in general terms.

## Metrics
TODO: Describe the metrics exposed by this service, and what they indicate.

## API
TODO: Describe the API endpoints this service provides, and what they do. No
need to go into as much specificity as the Swagger spec will, but be
descriptive.

## Configuration
* `{{Name | toUpper}}_LOG_LEVEL`: Log level. Defaults to "info".
* `{{Name | toUpper}}_HTTP_PORT`: Port on which the HTTP API service listens. Defaults to {{Port}}.
* `{{Name | toUpper}}_AUTH_ENABLED`: Whether authentication is enforced. If false, middleware is used that injects an admin identity into unauthenticated requests.
* `{{Name | toUpper}}_AUTH_KEY_FILE`: The file containing the secret key that is used to validate incoming authentication tokens.
* `{{Name | toUpper}}_TRACING_ENABLED`: Whether request tracing is enabled.
