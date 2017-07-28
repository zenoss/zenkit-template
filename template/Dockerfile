{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
FROM zenoss/zenkit-build:1.3

COPY . /go/src/{{$pkg}}
WORKDIR /go/src/{{$pkg}}/{{Name}}
RUN go build -o /bin/{{Name}}

FROM alpine
RUN apk add --no-cache curl
COPY --from=0 /bin/{{Name}} /bin/{{Name}}
COPY ./{{Name}}.yml /etc/{{Name}}.yml

EXPOSE {{Port}}

CMD ["/bin/{{Name}}", "server"]
