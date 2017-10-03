{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
ARG ZENKIT_BUILD_VERSION=1.6.2

FROM zenoss/zenkit-build:${ZENKIT_BUILD_VERSION}

COPY . /go/src/{{$pkg}}
WORKDIR /go/src/{{$pkg}}/{{Name}}
RUN go build -o /bin/{{Name}}

FROM alpine
RUN apk add --no-cache curl
COPY --from=0 /bin/{{Name}} /bin/{{Name}}
COPY ./{{Name}}.yml /etc/{{Name}}.yml

EXPOSE {{Port}}
EXPOSE {{AdminPort}}

CMD ["/bin/{{Name}}", "server"]
