## Status
**READY** | **IN DEVELOPMENT** | **HOLD**

## Description
A few sentences describing the overall goals of the pull request's commits

## Code Review Checklist

### Design/Swagger

- [ ] All actions, params, and attributes have descriptions and examples.

### Resource

- [ ] Changes follow [contribution guidelines](https://sites.google.com/a/zenoss.com/engineering/home/faq/contributing-guidelines)
- [ ] Controller implementation is minimal, and non-trivial business logic is offloaded to helper files or packages.  Calling `make app` does not blow away changes.

### Helpers and Packages

- [ ] Types and function declarations are commented
- [ ] Errors are wrapped with user-friendly messages using `github.com/pkg/errors`

### Configuration

- [ ] Default config is set in the `{{Name}}.yml`
- [ ] Developer-based config is set in the `docker-compose.yml`
- [ ] External service dependencies is configured in the `docker-compose.yml`. Calling `make run` works without additional setup.

### Documentation

- [ ] New metrics, configs, and api calls are documented in the README and runbook
- [ ] Changes to service operations are documented in the runbook 