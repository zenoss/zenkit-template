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

- `{{replace Name "-" "_" -1 | toUpper}}_LOG_LEVEL`: Log level. Defaults to "info".
- `{{replace Name "-" "_" -1 | toUpper}}_LOG_STACKDRIVER`: Whether to format logs for Stackdriver. Defaults to true.
- `{{replace Name "-" "_" -1 | toUpper}}_AUTH_DISABLED`: Whether authentication is enforced. If true, middleware is used that injects an admin identity into unauthenticated requests.
- `{{replace Name "-" "_" -1 | toUpper}}_AUTH_DEV_TENANT`: When auth is disabled, the tenant name to use as the identity. Defaults to "ACME".
- `{{replace Name "-" "_" -1 | toUpper}}_AUTH_DEV_USER`: When auth is disabled, the user id to use as the identity. Defaults to "zcuser@acme.example.com".
- `{{replace Name "-" "_" -1 | toUpper}}_TRACING_ENABLED`: Whether request tracing is enabled.
