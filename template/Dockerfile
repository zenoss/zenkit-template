{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
FROM golang:1.8.3-alpine

COPY . /go/src/{{$pkg}}
WORKDIR /go/src/{{$pkg}}/{{Name}}
RUN go build -o /bin/{{Name}}

FROM alpine
COPY --from=0 /bin/{{Name}} /bin/{{Name}}
COPY ./{{Name}}.yml /etc/{{Name}}.yml

EXPOSE {{Port}}

CMD ["/bin/{{Name}}", "server"]
