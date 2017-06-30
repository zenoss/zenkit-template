# Zenkit μsvc Template

Given Docker and a Go development environment, this repository can create you
a ready-to-go microservice scaffold.

## Required environment
You should have Go and Docker set up and usable by your current user. You
should also be in the directory under `$GOPATH` representing the owner of the
repository you're creating (e.g., `$GOPATH/src/github.com/zenoss`).

## Quickstart
Just run this to create a microservice named `examplesvc`:

    bash <(curl -sSL https://goo.gl/ckYau) examplesvc

(That shortlink is just [create.sh](https://raw.githubusercontent.com/zenoss/zenkit-template/master/create.sh) from this very repo.)

This will ask you a series of questions. You can always change the answers
later, except the first one, which is prefilled for you.

Once it's generated, go into your new directory and run `make` to pull in
dependencies and get everything set up:

    cd examplesvc
    make

Now you can start the thing, if you want. It doesn't do much, but it will
run:

    make run

