{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
ARG ZENKIT_BUILD_VERSION=1.5.1

FROM zenoss/zenkit-build:${ZENKIT_BUILD_VERSION}

COPY . /go/src/{{$pkg}}
WORKDIR /go/src/{{$pkg}}/{{Name}}
RUN go build -o /bin/{{Name}}

FROM alpine
RUN apk add --no-cache curl
COPY --from=0 /bin/{{Name}} /bin/{{Name}}
COPY ./{{Name}}.yml /etc/{{Name}}.yml

EXPOSE {{Port}}

CMD ["/bin/{{Name}}", "server"]
