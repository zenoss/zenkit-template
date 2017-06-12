# {{Title}} ({{Name}})
{{Description}}

## Prerequisites
To develop and run `{{Name}}`, you will need:
* make
* docker-ce >= 17.05 (Official installation instructions for [Ubuntu](https://docs.docker.com/engine/installation/linux/ubuntu/) | [CentOS](https://docs.docker.com/engine/installation/linux/centos/) | [macOS](https://docs.docker.com/docker-for-mac/install/))
* docker-compose (Install from here via `sudo make docker-compose`, or see the [official instructions](https://docs.docker.com/compose/install/))

Additional helpful utilities include:
* [httpie](https://httpie.org/) (`apt install httpie` on Ubuntu)
* [jq](https://stedolan.github.io/jq/) (`apt install jq` on Ubuntu)

## Development

## Building for production

## Testing


## Quickstart
<aside class="notice">
Hey, developer: change this once the Hello resource has been removed and you
have something representing this service's actual function.
</aside>
1. `make app run`
2. `http :{{Port}}/hello/world`
3. `http :{{Port}}/hello/newman`

## Environment Variables
* `{{Name | toUpper}}_PORT`: {{Port}}

change dc version to 3.1


Git Repo: zenoss/zing-{{Name}}
* develop is default branch
* Can't push to master or whatever
* Tests are required
* Permissions are not wrong
* git flow init -d
* Dockerfile
* docker-compose.yml (optional)
* Makefile
 * make test
 * make build
 * make run
 * make clean
* README with config documented
* Swagger spec
* Runbook template
Docker Repo: zenoss/zing-{{Name}}
Jenkins - test
Jenkins - build
